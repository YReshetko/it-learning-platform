package main

import (
	"github.com/YReshetko/it-academy-cources/api-app/internal/clients"
	"github.com/YReshetko/it-academy-cources/api-app/internal/config"
	"github.com/YReshetko/it-academy-cources/api-app/internal/http"
	"github.com/YReshetko/it-academy-cources/api-app/internal/http/handlers"
	"github.com/YReshetko/it-academy-cources/api-app/internal/http/middlewares/authorization"
	"github.com/YReshetko/it-academy-cources/api-app/internal/http/routes"
	"log"
)

func main() {
	cfg := handleError(config.LoadConfig())
	server := http.NewServer(cfg.HTTP)

	authClient := clients.NewAuthClient(cfg.AuthClient)
	auth := handlers.NewAuth(authClient)
	r := routes.NewRouter(
		routes.WithAuthHandler(auth),
		routes.WithAuthService(authorization.Service{}),
	)
	r.Init(server.Engine)

	server.Start()

}
func handleError[T any](val T, err error) T {
	if err != nil {
		log.Fatalf("crashing app due to: %s", err)
	}
	return val
}
