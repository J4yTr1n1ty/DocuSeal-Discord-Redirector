package docuseal

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"slices"

	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/config"
	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/discord"
	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/types"
	"github.com/gorilla/mux"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Verify request key
	if params["key"] == "" || !slices.Contains(config.Config.AuthorizedKeys, params["key"]) {
		log.Println("Unauthorized request from " + r.RemoteAddr)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body: " + err.Error())
		return
	}

	// Parse the body as a DocuSealEvent
	var dsEvent types.DocuSealEvent
	err = json.Unmarshal(body, &dsEvent)
	if err != nil {
		log.Println("Error unmarshalling request body: " + err.Error())
		return
	}

	discord.AssembleMessage(dsEvent)

	w.WriteHeader(http.StatusOK)
}
