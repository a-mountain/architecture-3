package menu

import (
	"github.com/nickname038/architecture-3/lib"
	"log"
	"net/http"
)

func HttpHandler(facade *MenuFacade) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListChannels(facade, rw)
		} else if r.Method == "POST" {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListChannels(facade *MenuFacade, rw http.ResponseWriter) {
	res, err := facade.GetAllMenuItems()
	if err != nil {
		log.Printf("Error %s", err)
		lib.WriteJsonInternalError(rw)
		return
	}
	lib.WriteJsonOk(rw, res)
}
