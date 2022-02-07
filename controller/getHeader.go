package controller

import (
	"fmt"
	"httpserver/lib"
	"httpserver/model"
	"io"
	"net/http"
)

func GetHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		vStr := lib.Slice2Str(v)
		w.Header().Set(k, vStr)
	}
	io.WriteString(w, fmt.Sprintf("version = %s\n", model.GetVersion()))
	w.WriteHeader(http.StatusOK)
}
