/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:19:16
 * @LastEditors: lhl
 * @LastEditTime: 2024-08-05 15:12:10
 */
package control

import (
	"context"
	"fmt"
	"gotoexec/config"
	"gotoexec/grpcapi"
	"gotoexec/server/implant"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type control struct {
	work, output chan *grpcapi.Command
}

var ControlInstance *control

var (
	implantListener net.Listener
	err             error
	opts            []grpc.ServerOption
	work, output    chan *grpcapi.Command
)

func init() {

	//加载配置
	config.Init("conf.toml")

	work, output = make(chan *grpcapi.Command), make(chan *grpcapi.Command)
	//植入程序服务端和管理程序服务端使用相同的通道
	implantServer := implant.NewImplantServer(work, output)

	ControlInstance = Newcontrols(work, output)

	//服务端建立监听，植入服务端与管理服务端监听的端口分别是4001和4002
	if implantListener, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.CoreConf.ListenPort)); err != nil {
		log.Fatalln("implantserver" + err.Error())
	}

	opts = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 12),
		grpc.MaxSendMsgSize(1024 * 1024 * 12),
	}
	grpcImplantServer := grpc.NewServer(opts...)

	grpcapi.RegisterImplantServer(grpcImplantServer, implantServer)

	//使用goroutine启动植入程序服务端，防止代码阻塞，毕竟后面还要开启管理程序服务端
	go func() {
		grpcImplantServer.Serve(implantListener)
	}()
}

func Newcontrols(work, output chan *grpcapi.Command) *control {
	s := new(control)
	s.work = work
	s.output = output
	return s
}

func (s *control) RunCommand(cmd *grpcapi.Command) (*grpcapi.Command, error) {
	var res *grpcapi.Command
	go func() {
		s.work <- cmd
	}()

	res = <-s.output

	return res, nil
}

func (s *control) RunCommandCtx(ctx context.Context, cmd *grpcapi.Command) (*grpcapi.Command, error) {
	var res *grpcapi.Command

	// 在goroutine中发送命令到工作队列
	go func() {
		s.work <- cmd
	}()

	// 设置超时时间
	select {
	case res = <-s.output:
		// 成功从输出通道接收到结果
	case <-ctx.Done():
		// 上下文超时或取消
		return nil, ctx.Err()
	case <-time.After(8 * time.Second):
		// 指定的超时时间，例如5秒
		return nil, fmt.Errorf("command execution timed out")
	}

	return res, nil
}

func (s *control) SetEmpty(cmd *grpcapi.Command) {

	s.output <- cmd

}

func (s *control) SetSleepTime(time *grpcapi.SleepTime) (*grpcapi.Empty, error) {
	return &grpcapi.Empty{}, nil
}
