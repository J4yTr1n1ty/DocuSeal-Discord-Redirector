package api

import (
	"net/http"

	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/docuseal"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.HandleFunc("/incoming/{key}", docuseal.Handle)

	return router
}
