package orders

import (
	"encoding/json"
	"github.com/nickname038/architecture-3/lib"
	"log"
	"net/http"
)

func HttpHandler(facade *OrderFacade) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handleCreateOrder(r, rw, facade)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleCreateOrder(r *http.Request, rw http.ResponseWriter, facade *OrderFacade) {
	var request CreateOrderRequest
	json.NewDecoder(r.Body).Decode(&request)
	res, err := facade.CreateOrder(request)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		lib.WriteJsonInternalError(rw)
		return
	}
	lib.WriteJsonOk(rw, res)
}
