package main

import "github.com/YReshetko/it-academy-cources/api-management/internal/http"

func main() {
	server := http.NewServer()
	server.Start()
}
