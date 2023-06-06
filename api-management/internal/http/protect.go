package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
)

func protect(handler http.HandlerFunc) http.HandlerFunc {
	var client = gocloak.NewClient("http://localhost:8247")
	clientId := "academy"
	realm := "it-academy"
	clientSecret := "9v1wwWEJsyYMZF7AGXYZlWP0cNnIGAh0"

	return func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		if len(authHeader) < 1 {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(`{"error": "unauthorized"}`)
			return
		}
		accessToken := strings.Split(authHeader, " ")[1]
		//fmt.Println(accessToken)
		rptResult, err := client.RetrospectToken(request.Context(), accessToken, clientId, clientSecret, realm)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(`{"error": "bad request"}`)
			return
		}
		isTokenValid := *rptResult.Active
		js, _ := json.Marshal(rptResult)
		fmt.Println(string(js))
		if !isTokenValid {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(`{"error": "invalid token"}`)
			return
		}

		info, err := client.GetUserInfo(request.Context(), accessToken, realm)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(`{"error": "can not get user info"}`)
			return
		}
		js, _ = json.Marshal(info)
		fmt.Println(string(js))

		handler(writer, request)
	}
}
