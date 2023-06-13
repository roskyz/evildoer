package service

import (
	"crypto/tls"
	"crypto/x509"
	"evildoer/protogen"
	"os"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// todo: mTLS serverName [config.conf]

func loadServerCert() credentials.TransportCredentials {
	caCertPem, err := os.ReadFile("conf/cert/ca-cert.pem")
	if err != nil {
		return nil
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCertPem)

	serverCert, err := tls.LoadX509KeyPair("conf/cert/server-cert.pem", "conf/cert/server-key.pem")
	if err != nil {
		return nil
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ServerName:   "server.web.local",
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS12,
	})
}

var defaultServerMaxMessageSize = 1024 * 1024 * 1

func NewService() (*grpc.Server, error) {
	authInterceptor := newServerAuthInterceptor()
	grpcSrv := grpc.NewServer(
		grpc.MaxRecvMsgSize(defaultServerMaxMessageSize),
		grpc.MaxSendMsgSize(defaultServerMaxMessageSize),
		grpc.NumStreamWorkers(uint32(runtime.NumCPU())),
		grpc.Creds(loadServerCert()),
		//grpc.InTapHandle(newServerAuthInterceptor().serverTapFunc),
		grpc.UnaryInterceptor(authInterceptor.Unary()),
		grpc.StreamInterceptor(authInterceptor.Stream()),
	)
	// todo: errgroup
	if err := regAuthService(grpcSrv); err != nil {
		return nil, err
	}
	if err := regGroupService(grpcSrv); err != nil {
		return nil, err
	}
	if err := regFormService(grpcSrv); err != nil {
		return nil, err
	}
	return grpcSrv, nil
}

func defaultResponse(err ...error) *protogen.DefaultResponse {
	if len(err) != 0 && err[0] != nil {
		return &protogen.DefaultResponse{
			Code:    -1,
			Message: err[0].Error(),
		}
	}
	return &protogen.DefaultResponse{
		Code: 0, Message: "rpc invoked with no errors",
	}
}
