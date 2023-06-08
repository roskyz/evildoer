package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"evildoer/pkg/authmgr"
	"evildoer/protogen"
	"fmt"
	"os"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type caller struct {
	cc grpc.ClientConnInterface
}

func loadClientCert() credentials.TransportCredentials {
	caCertPem, err := os.ReadFile("conf/cert/ca-cert.pem")
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCertPem)

	clientCert, err := tls.LoadX509KeyPair("conf/cert/client-cert.pem", "conf/cert/client-key.pem")
	if err != nil {
		panic(err)
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		ServerName:   "client.web.local",
		RootCAs:      certPool,
	})
}

type Interceptor interface {
	Unary() grpc.UnaryClientInterceptor
	Stream() grpc.StreamClientInterceptor
}

type CallerOption func(ctx context.Context, targetURI string) ([]grpc.DialOption, error)

func WithAuthInterceptor(auth *authmgr.AuthOpt) CallerOption {
	return func(ctx context.Context, targetURI string) ([]grpc.DialOption, error) {
		interceptor, err := newClientAuthInterceptor(ctx, targetURI, auth)
		if err != nil {
			return nil, err
		}
		return WithInterceptor(interceptor)(ctx, targetURI)
	}
}

func WithInterceptor(interceptor Interceptor) CallerOption {
	return func(ctx context.Context, targetURI string) ([]grpc.DialOption, error) {
		return []grpc.DialOption{
			grpc.WithUnaryInterceptor(interceptor.Unary()),
			grpc.WithStreamInterceptor(interceptor.Stream()),
		}, nil
	}
}

func newCaller(ctx context.Context, targetURI string, options ...CallerOption) (grpc.ClientConnInterface, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(loadClientCert()),
	}
	for _, option := range options {
		extOpts, err := option(ctx, targetURI)
		if err != nil {
			return nil, err
		}
		dialOpts = append(dialOpts, extOpts...)
	}
	//	perCredentials, _ := oauth.NewJWTAccessFromKey(credentials)
	//	dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(perCredentials))
	conn, err := grpc.DialContext(ctx, targetURI, dialOpts...)
	if err != nil {
		return nil, err
	}
	return newCallerWithConn(conn), nil
}

func newCallerWithConn(cc grpc.ClientConnInterface) grpc.ClientConnInterface {
	return &caller{cc: cc}
}

type callerError struct {
	*protogen.DefaultResponse
}

func (c *callerError) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", c.Code, c.Message)
}

func (c *caller) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	err := c.cc.Invoke(ctx, method, args, reply, opts...)
	if err == nil {
		// handle app biz error
		defaultResp := handleResponseError(reply)
		if defaultResp != nil && defaultResp.Code != 0 {
			return &callerError{defaultResp}
		}
	}
	return err
}

func handleResponseError(reply interface{}) *protogen.DefaultResponse {
	defaultResp, ok := reply.(*protogen.DefaultResponse)
	if ok {
		return defaultResp
	}
	defaultRespField := reflect.Indirect(reflect.ValueOf(reply)).
		FieldByName(reflect.TypeOf(protogen.DefaultResponse{}).Name())
	if !defaultRespField.IsZero() {
		return defaultRespField.Interface().(*protogen.DefaultResponse)
	}
	return nil
}

func (c *caller) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return c.cc.NewStream(ctx, desc, method, opts...)
}

var _ grpc.ClientConnInterface = new(caller)
