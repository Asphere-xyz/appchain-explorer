package gateway

import (
	"context"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/database"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strings"
)

type key int

const (
	requestIDKey key = 0
)

type Service struct {
	databaseService *database.Service
	// state
	grpcServer *grpc.Server
}

func NewService(databaseService *database.Service) *Service {
	return &Service{databaseService: databaseService}
}

func (s *Service) Start(cp shared.IConfigProvider) error {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	listener, err := net.Listen("tcp", config.GrpcAddress)
	if err != nil {
		return errors.Wrapf(err, "failed to listen for address (%s)", config.GrpcAddress)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	grpc_zap.ReplaceGrpcLoggerV2WithVerbosity(logger, int(zap.DebugLevel))
	s.grpcServer = grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
		),
	)
	grpc_prometheus.Register(s.grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()
	types.RegisterBlockscoutGatewayServer(s.grpcServer, s)
	go func() {
		log.Printf("gateway gRPC server is listening on address %s", config.GrpcAddress)
		if err = s.grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), config.GrpcAddress, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	if err := types.RegisterBlockscoutGatewayHandler(context.Background(), mux, conn); err != nil {
		return fmt.Errorf("failed to register gRPC-gateway handler: %+v", err)
	}
	var httpListener net.Listener
	if config.HttpAddress != config.GrpcAddress {
		httpListener, err = net.Listen("tcp", config.HttpAddress)
		if err != nil {
			return errors.Wrapf(err, "failed to listen for address (%s)", config.GrpcAddress)
		}
	} else {
		httpListener = listener
	}
	go func() {
		log.Infof("gateway HTTP server is listening on address %s", config.HttpAddress)
		err := http.Serve(httpListener, allowCORS(mux))
		if err != nil {
			log.Panicf("failed to start http server: %+v", err)
		}
	}()
	return nil
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	return
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func logging(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Debug("Request",
					zap.String("RequestID", requestID),
					zap.String("Method", r.Method),
					zap.String("Path", r.URL.Path),
					zap.String("Addr", r.RemoteAddr),
					zap.String("UserAgent", r.UserAgent()),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
