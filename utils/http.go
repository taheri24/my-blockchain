package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FileLocation = string
type ContextKey string

const (
	ErrorCaptureKey ContextKey = "errorCapture"
)

func ServeWithJsonFile(w http.ResponseWriter, filePath FileLocation) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w.Header().Add(`content-type`, "application/json")
	io.Copy(w, f)

}

func ServeError(ctx context.Context, w http.ResponseWriter, err error) {
	setError, okSetError := ctx.Value(ErrorCaptureKey).(func(error))
	if okSetError && setError != nil {
		setError(err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "{%q:%q}", "error", err.Error())

}

func ErrorHandlerByRequest(r *http.Request, w http.ResponseWriter) func(err error) {

	return func(err error) {
		ServeError(r.Context(), w, err)
	}
}
