package debugger

import "runtime"
import "fmt"

func Where() string {
	var _, file, line, _ = runtime.Caller(1)
	return fmt.Sprintf("%s:%d", file, line)
}
