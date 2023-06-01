package evildoer

import (
	"evildoer/biz/service"
	"evildoer/conf"
	"evildoer/utils"
	"net"
)

//go:generate protoc --proto_path=./idl --go_out=./protogen --go_opt=paths=source_relative --go-grpc_out=./protogen --go-grpc_opt=paths=source_relative idl/*.proto

func main() {
	srv, err := service.NewService()
	// todo: logs
	utils.AssertErr(err)
	listener, err := net.Listen("tcp", conf.GetConf().BindAddr)
	utils.AssertErr(err)
	utils.AssertErr(srv.Serve(listener))
}
