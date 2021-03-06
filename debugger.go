package debugger

import "runtime"
import "fmt"
import "time"

func Where() string {
	return where(2)
}

func where(ignore int) string {
	var _, file, line, _ = runtime.Caller(ignore)
	return fmt.Sprintf("%s:%d", file, line)
}

func Time(begin time.Time) {
	fmt.Printf("Func Duration: %v\n", time.Now().Sub(begin))
}

func TimePrompt(begin time.Time, prompt string) {
	fmt.Printf("Func Duration: %v(begin at %s)\n", time.Now().Sub(begin), prompt)
}
