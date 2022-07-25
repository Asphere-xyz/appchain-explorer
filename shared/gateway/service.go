package gateway

import (
	"context"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/database"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strings"
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
	s.grpcServer = grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_prometheus.StreamServerInterceptor,
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
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
	fs := http.FileServer(http.Dir(config.StaticFolder))
	go func() {
		log.Infof("gateway HTTP server is listening on address %s and static folder serving at %s", config.HttpAddress, config.StaticFolder)
		err := http.Serve(httpListener, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			// check is it cors request
			if ok := handleCorsRequest(res, req); ok {
				return
			}
			// check is it api request
			if strings.HasPrefix(req.URL.Path, "/v1alpha") {
				mux.ServeHTTP(res, req)
				return
			}
			// serve static files from dir
			if !strings.Contains(req.URL.Path, "/static") {
				req, _ = http.NewRequest(req.Method, "/", req.Body)
			}
			fs.ServeHTTP(res, req)
		}))
		if err != nil {
			log.Panicf("failed to start http server: %+v", err)
		}
	}()
	return nil
}

func handleCorsRequest(w http.ResponseWriter, r *http.Request) bool {
	var origin string
	if origin = r.Header.Get("Origin"); origin == "" {
		return false
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method != "OPTIONS" || r.Header.Get("Access-Control-Request-Method") == "" {
		return false
	}
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	return true
}
