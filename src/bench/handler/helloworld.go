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
	log.Printf("request %s", r.RequestURI)
	ctx, cancel := context.WithTimeout(context.Background(),1*time.Second)

	defer cancel()

	id := r.URL.Query().Get("id")
	dataChan := h.Srvc.GetData(id, ctx)
	var data string

	select {
	case data = <-dataChan:
		w.Write([]byte(data))
	case <-ctx.Done():
		w.Write([]byte("timed out"))
	}

	log.Printf("number of goroutines %d. request %s", runtime.NumGoroutine(),r.RequestURI)
}
