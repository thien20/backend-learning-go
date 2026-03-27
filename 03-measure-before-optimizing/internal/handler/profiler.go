package handler

import "net/http"

func registerProfiler(mux *http.ServeMux) {
	// TODO: Wire `net/http/pprof` here and confirm which profiles are useful for
	// this app. Start with CPU and heap profiles before adding more endpoints.
	mux.HandleFunc("GET /debug/pprof/", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "pprof wiring is a student task", http.StatusNotImplemented)
	})
}
