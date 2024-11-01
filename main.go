package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = w.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	var h Handler
	s := &http.Server{
		Addr:           ":8080",
		Handler:        &h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
