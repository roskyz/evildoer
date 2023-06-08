package service

import (
	context "context"
	"evildoer/pkg/authmgr"
	"evildoer/protogen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/tap"
)

type authService struct {
	*protogen.UnimplementedAuthSrvServer
	jwtMgr *authmgr.JWTManager
}

func (a *authService) Auth(ctx context.Context, req *protogen.AuthReq) (*protogen.AuthResponse, error) {
	tknStr, err := a.jwtMgr.IsAllowedUser(req.Username, req.Password)
	return &protogen.AuthResponse{
		DefaultResponse: defaultResponse(err),
		AuthJwt:         tknStr,
	}, nil
}

func regAuthService(grpcSrv *grpc.Server) error {
	protogen.RegisterAuthSrvServer(grpcSrv, &authService{
		jwtMgr: authmgr.NewJWTManger(),
	})
	return nil
}

type authInterceptor struct {
	jwtMgr *authmgr.JWTManager
}

func newServerAuthInterceptor() *authInterceptor {
	return &authInterceptor{jwtMgr: authmgr.NewJWTManger()}
}

// todo exp method to handle all server request
// like: grpc.InTapHandle(tap.ServerInHandle)
func (interceptor *authInterceptor) serverTapFunc(ctx context.Context, info *tap.Info) (context.Context, error) {
	return interceptor.authorizeAndSetContext(ctx, info.FullMethodName)
}

func (interceptor *authInterceptor) authorizeAndSetContext(ctx context.Context, method string) (context.Context, error) {
	if method == protogen.AuthSrv_Auth_FullMethodName {
		return ctx, nil
	}
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := meta["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	credentials, err := interceptor.jwtMgr.GetCredentials(values[0])
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if credentials.Username != "" {
		return interceptor.jwtMgr.SetCtxInfo(ctx, credentials), nil
	}

	return ctx, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}

func (interceptor *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = interceptor.authorizeAndSetContext(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

type serverStreamWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStreamWrapper) Context() context.Context { return w.ctx }

func (interceptor *authInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		wrapped := &serverStreamWrapper{
			ss, ss.Context(),
		}
		wrapped.ctx, err = interceptor.authorizeAndSetContext(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, wrapped)
	}
}
