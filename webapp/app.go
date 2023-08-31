package webapp

import (
	"context"
	"log"
	"net/http"
	"time"
)

type App struct {
	mux *http.ServeMux
}

func New() *App {

	return &App{
		mux: http.NewServeMux(),
	}
}

func (a *App) AddHandle(route string, handler http.Handler) {
	a.mux.Handle(route, handler)
}

func (a *App) Serve(ctx context.Context) (err error) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: a.mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen error:%+s\n", err)
		}
	}()
	log.Println("Serving static content on localhost:8080...")

	<-ctx.Done()

	log.Println("Server stopped")

	ctxShut, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err = srv.Shutdown(ctxShut); err != nil {
		log.Fatal("Server shutdown failed: %s", err)
	}

	log.Println("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
