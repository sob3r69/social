package main

import "net/http"

type applictaion struct {
	config config
}

type config struct {
	addr string
}

func (app *applictaion) run() error {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	return srv.ListenAndServe()
}
