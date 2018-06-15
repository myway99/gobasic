package main

import (
	"net/http"
	//"log"
	"os"
	_ "net/http/pprof"

	"mytest04/basic/errhandling/filelistingserver/filelisting"

	"github.com/gpmgo/gopm/modules/log"
	"fmt"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				//log.Printf("Panic: %v", r)
				fmt.Printf("Panic: %v", r)

				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			// 来自gopm包的log
			log.Warn("Error occurred"+
				"handling request: %s",
				err.Error())
			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden

			default:
				code = http.StatusInternalServerError
				//http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}
			http.Error(writer,
				http.StatusText(code), code)

		}
	}

}

type userError interface {
	error
	Message() string
}

//通过网页显示文件
func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
