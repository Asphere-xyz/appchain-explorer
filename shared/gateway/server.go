package gateway

import (
	"context"
	"fmt"
	"github.com/Ankr-network/appchain-explorer/shared"
	"github.com/Ankr-network/appchain-explorer/shared/database"
	"github.com/Ankr-network/appchain-explorer/shared/staking"
	"github.com/Ankr-network/appchain-explorer/shared/types"
	"github.com/Ankr-network/appchain-explorer/shared/websocket"
	mux2 "github.com/gorilla/mux"
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

type Server struct {
	databaseService *database.Service
	stateDbService  *database.StateDb
	websocketServer *websocket.Server
	stakingService  *staking.Service
	// state
	grpcServer *grpc.Server
}

func NewServer(
	databaseService *database.Service,
	stateDbService *database.StateDb,
	websocketServer *websocket.Server,
	stakingService *staking.Service,
) *Server {
	return &Server{
		databaseService: databaseService,
		stateDbService:  stateDbService,
		websocketServer: websocketServer,
		stakingService:  stakingService,
	}
}

func (s *Server) Start(cp shared.IConfigProvider) error {
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
	types.RegisterStakingGatewayServer(s.grpcServer, s)
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
	} else if err := types.RegisterStakingGatewayHandler(context.Background(), mux, conn); err != nil {
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
		router := mux2.NewRouter()
		router.PathPrefix("/v1alpha").Handler(mux)
		router.PathPrefix("/static").Handler(fs)
		router.PathPrefix("/ws").Handler(s.websocketServer)
		router.PathPrefix("/").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			// special case for serving files with extensions
			path := strings.Split(req.URL.Path, "/")
			if len(path) > 0 {
				fileNameWithExt := strings.Split(path[len(path)-1], ".")
				if len(fileNameWithExt) > 1 {
					fs.ServeHTTP(res, req)
					return
				}
			}
			req, _ = http.NewRequest(req.Method, "/", req.Body)
			// fallback to the index.html by default
			fs.ServeHTTP(res, req)
		})
		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if !handleCorsRequest(w, req) {
					next.ServeHTTP(w, req)
				}
			})
		})
		if err := http.Serve(httpListener, router); err != nil {
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
