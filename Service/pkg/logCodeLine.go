package pkg

import (
	"log"
	"runtime"
)

// LogginHandleError logs the error and code line where the error occurred
func LogginHandleError(err error) (b bool) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)

		log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
		b = true
	}
	return
}
