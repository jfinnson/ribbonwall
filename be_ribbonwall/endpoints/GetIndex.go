package endpoints

import (
	"fmt"
	"net/http"

	log "github.com/ribbonwall/common/logging"
)

func (s *Endpoints) GetIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			log.Error("Error!!")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Hello, World!!")
	})
}
