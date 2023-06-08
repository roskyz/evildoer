package client

import (
	"context"
	"evildoer/pkg/authmgr"
	"evildoer/pkg/jwtx"
	"evildoer/protogen"
	"evildoer/utils"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type authClient struct {
	clientSrv          protogen.AuthSrvClient
	username, password string
}

func (client *authClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := client.clientSrv.Auth(ctx, &protogen.AuthReq{
		Username: client.username,
		Password: client.password,
	})
	if err != nil {
		return "", err
	}
	return response.GetAuthJwt(), nil
}

type authInterceptor struct {
	authClient      *authClient
	tknStr          string
	refreshDuration time.Duration
}

func newClientAuthInterceptor(ctx context.Context, targetURI string, auth *authmgr.AuthOpt) (*authInterceptor, error) {
	conn, err := newCaller(ctx, targetURI)
	if err != nil {
		return nil, err
	}
	interceptor := &authInterceptor{
		authClient: &authClient{
			clientSrv: protogen.NewAuthSrvClient(conn),
			username:  auth.Username, password: auth.Password,
		},
	}
	if err = interceptor.scheduleRefreshToken(); err != nil {
		return nil, err
	}
	return interceptor, nil
}

func (interceptor *authInterceptor) scheduleRefreshToken() error {
	if err := interceptor.refreshToken(); err != nil {
		return err
	}

	go func() {
		wait := interceptor.refreshDuration - time.Second*10
		timer := time.NewTimer(wait)
		defer timer.Stop()

		for range timer.C {
			utils.AssertErr(interceptor.refreshToken())
			timer.Reset(interceptor.refreshDuration - time.Second*10)
		}
	}()

	return nil
}

func (interceptor *authInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.tknStr)
}

func (interceptor *authInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
	}
}

func (interceptor *authInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(interceptor.attachToken(ctx), desc, cc, method, opts...)
	}
}

func (interceptor *authInterceptor) refreshToken() error {
	tknStr, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}
	refreshDuration := jwtx.GetExpireTime(tknStr).Sub(time.Now())
	interceptor.refreshDuration = refreshDuration
	interceptor.tknStr = tknStr
	return nil
}
