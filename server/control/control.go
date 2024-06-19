/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:19:16
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 11:54:23
 */
package control

import (
	"fmt"
	"gotoexec/config"
	"gotoexec/grpcapi"
	"gotoexec/server/implant"
	"log"
	"net"

	"google.golang.org/grpc"
)

type control struct {
	work, output chan *grpcapi.Command
}

var NewImplantcontrols *control

func init() {
	var (
		implantListener net.Listener
		err             error
		opts            []grpc.ServerOption
		work, output    chan *grpcapi.Command
	)

	//加载配置
	config.Init("conf.toml")

	work, output = make(chan *grpcapi.Command), make(chan *grpcapi.Command)
	//植入程序服务端和管理程序服务端使用相同的通道
	implantServer := implant.NewImplantServer(work, output)

	NewImplantcontrols = NewImplantcontrol(work, output)

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

func NewImplantcontrol(work, output chan *grpcapi.Command) *control {
	s := new(control)
	s.work = work
	s.output = output
	return s
}

func (s *control) RunCommand(cmd *grpcapi.Command) (*grpcapi.Command, error) {
	fmt.Println(cmd.In)
	var res *grpcapi.Command
	go func() {
		s.work <- cmd
	}()

	res = <-s.output

	return res, nil
}
func (s *control) SetSleepTime(time *grpcapi.SleepTime) (*grpcapi.Empty, error) {
	return &grpcapi.Empty{}, nil
}
