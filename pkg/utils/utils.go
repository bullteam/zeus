package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"net"
	"reflect"
	"time"
)

func TestTCPConn(addr string, timeout, interval int) error {
	success := make(chan int)
	cancel := make(chan int)

	go func() {
		for {
			select {
			case <-cancel:
				break
			default:
				conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Second)
				if err != nil {
					beego.Error("failed to connect to tcp://%s, retry after %d seconds :%v",
						addr, interval, err)
					time.Sleep(time.Duration(interval) * time.Second)
					continue
				}
				conn.Close()
				success <- 1
				break
			}
		}
	}()

	select {
	case <-success:
		return nil
	case <-time.After(time.Duration(timeout) * time.Second):
		cancel <- 1
		return fmt.Errorf("failed to connect to tcp:%s after %d seconds", addr, timeout)
	}
}



func IsNilObject(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}