package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func registerRoutes(router *mux.Router) {

	router.HandleFunc("/models/v1/users", protect(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(
			`
			{
				"users": [
					{
						"id": 1,
						"name": Jack,
						"role": User
					}
				}
			}
			`,
		))
	})).Methods("GET")
}
