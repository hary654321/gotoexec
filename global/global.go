package global

type LoginLog struct {
	Time int64
	Ip   string
}

var LoginQue []LoginLog
