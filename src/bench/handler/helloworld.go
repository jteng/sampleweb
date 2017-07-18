package handler

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"

	"bench/engine"
)

type HelloWorld struct {
	Srvc *engine.ThirdpartyService
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	id := r.URL.Query().Get("id")
	dataChan := h.Srvc.GetData(id, ctx)
	var data string
	select {
	case data = <-dataChan:
		w.Write([]byte(data))
	case <-time.After(1 * time.Second):
		cancel()
		w.Write([]byte("timed out"))

	}
	log.Printf("number of goroutines %d", runtime.NumGoroutine())
}
