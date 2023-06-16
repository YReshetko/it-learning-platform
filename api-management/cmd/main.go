package main

import "github.com/YReshetko/it-learning-platform/api-management/internal/http"

func main() {
	server := http.NewServer()
	server.Start()
}
