package grpc

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	grpc_middeware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"time"
)

/*
	创建登录grpc
*/
func CreateServiceLoginConn(c *gin.Context) *grpc.ClientConn {
	return createGrpcConn("127.0.0.1:9230", c)
}

/*
	连接grpc
*/
func createGrpcConn(address string, c *gin.Context) *grpc.ClientConn {
	var conn *grpc.ClientConn
	var err error

	//超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	conn, err = grpc.DialContext(
		ctx,
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(
			grpc_middeware.ChainUnaryClient(),
		),
	)
	if err != nil {
		fmt.Println(address, "grpc conn err:", err)
	}
	return conn
}

/*
	// 调用 gRPC 服务
	conn := grpc_client.CreateServiceListenConn(c)
	grpcListenClient := listen.NewListenClient(conn)
	resListen, _ := grpcListenClient.ListenData(context.Background(), &listen.Request{Name: "listen"})
 */
