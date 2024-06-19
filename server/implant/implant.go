/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-06-19 11:23:41
 * @LastEditors: lhl
 * @LastEditTime: 2024-06-19 11:25:54
 */
package implant

import (
	"context"
	"errors"
	"fmt"
	"gotoexec/grpcapi"
)

var sleepTime int32 = 3

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
	var cmd = new(grpcapi.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {
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
	time.Time = sleepTime
	return time, nil
}
