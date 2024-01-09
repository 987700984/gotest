package main

import (
	"flag"
	"fmt"

	"github.com/987700984/gotest/video/api/rpc/internal/config"
	useractionServer "github.com/987700984/gotest/video/api/rpc/internal/server/useraction"
	userinfoServer "github.com/987700984/gotest/video/api/rpc/internal/server/userinfo"
	"github.com/987700984/gotest/video/api/rpc/internal/svc"
	"github.com/987700984/gotest/video/api/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserInfoServer(grpcServer, userinfoServer.NewUserInfoServer(ctx))
		user.RegisterUserActionServer(grpcServer, useractionServer.NewUserActionServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
