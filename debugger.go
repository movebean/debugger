package debugger

import "runtime"
import "fmt"
import "time"

func Where() string {
	var _, file, line, _ = runtime.Caller(1)
	return fmt.Sprintf("%s:%d", file, line)
}

func Time(begin time.Time) time.Duration {
	return time.Now().Sub(begin)
}
