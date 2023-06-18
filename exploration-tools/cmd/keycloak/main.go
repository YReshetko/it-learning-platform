package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"log"
)

func main() {
	var client = gocloak.NewClient("http://localhost:8082")
	clientId := "academy"
	realm := "it-academy"
	clientSecret := "Ch9EWb4fMp3ksYqAVLRTf39vsFuI6s8K"

	userAccessToken := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJvRWJSNWNkdnRVUl8teVJxRzlqVTE3SkVBaWI2UWZwc01mTk9DVkR2bEZBIn0.eyJleHAiOjE2ODcwNTgyMTMsImlhdCI6MTY4NzA1NzMxMywiYXV0aF90aW1lIjoxNjg3MDU3MzEzLCJqdGkiOiI1NTFjNWU0ZS1jZTE1LTQ5MzQtYjQ1ZC02YTNmMjQ5OWE2N2EiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODIvcmVhbG1zL2l0LWFjYWRlbXkiLCJzdWIiOiJmZDM1OTY2Yy1mZGY1LTRhOTUtODJkMi1kZTUzMDJlNWUwNmYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhY2FkZW15Iiwic2Vzc2lvbl9zdGF0ZSI6ImJmZDBlOTgyLWM2MDAtNDBiNi1hOTM5LWVkZGM2ZjViZDUzZiIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiKiJdLCJzY29wZSI6Im9wZW5pZCBlbWFpbCBwcm9maWxlIiwic2lkIjoiYmZkMGU5ODItYzYwMC00MGI2LWE5MzktZWRkYzZmNWJkNTNmIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiJ0ZXN0IHRlc3QiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJ0ZXN0IiwiZ2l2ZW5fbmFtZSI6InRlc3QiLCJmYW1pbHlfbmFtZSI6InRlc3QiLCJlbWFpbCI6InRlc3RAZy5jb20iLCJhY2FkZW15Ijp7InJvbGVzIjpbIlNUVURFTlQiXX19.J_PfyFHMAr4tLm2cc_fFZQWMIdUGERW_czZbAT2tRvlTs5J65_3-35o0QBhGc0vw1UJacnRLyu7dHgPY15FPLTcofHV0CRunvfnRYmO_lE-EVIaXgrMxh5nqgio0Y7gA00ltjjckMth1obSUBuuKClIlFI5GQlhvnWYHR3r5wYt5H3BvMdc4_dXUbP6O_dVJRGsXrTycvTjGB6OsBnBzc82yxzV-Zd7zy_Pan1I7oUg2eZF79K1Dhk-FmOMtmRywjSZAAwrt38kmWeeedcYO2YPcY3Y2N99tFJINli2yYviqmVHq5VEMcUysv5VgE7Gt1hSQa7d0JiSO6GQDvPcJuA"

	ctx := context.WithValue(context.Background(), "Host", "localhost:8081")
	retrospectResult, err := client.RetrospectToken(ctx, userAccessToken, clientId, clientSecret, realm)

	if err != nil {
		log.Fatalln("Retro", err)
	}

	printObject(retrospectResult)

	info, err := client.GetUserInfo(context.Background(), userAccessToken, realm)
	if err != nil {
		log.Fatalln("User info", err)
	}
	printObject(info)

	roles, err := client.GetRealmRolesByUserID(context.Background(), newClientToken(client), realm, *info.Sub)
	if err != nil {
		log.Fatalln(err)
	}
	printObject(roles)

}

func printObject(o any) {
	b, _ := json.Marshal(o)
	fmt.Println(string(b))
}

func newClientToken(client *gocloak.GoCloak) string {
	clientId := "academy"
	realm := "it-academy"
	clientSecret := "Ch9EWb4fMp3ksYqAVLRTf39vsFuI6s8K"
	grantType := "client_credentials"

	//createdAt := time.Now()
	t, err := client.GetToken(context.Background(), realm, gocloak.TokenOptions{
		ClientID:     &clientId,
		ClientSecret: &clientSecret,
		GrantType:    &grantType,
		//Scope:         toPtr("profile"),
		//ResponseTypes: &[]string{"token", "id_token"},
	})
	if err != nil {
		log.Fatalln(err)
	}
	return t.AccessToken
}
