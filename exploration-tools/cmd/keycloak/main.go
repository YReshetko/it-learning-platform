package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"log"
)

func main() {
	var client = gocloak.NewClient("http://localhost:8081")
	clientId := "academy"
	realm := "it-academy"
	clientSecret := "Ch9EWb4fMp3ksYqAVLRTf39vsFuI6s8K"

	userAccessToken := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJvRWJSNWNkdnRVUl8teVJxRzlqVTE3SkVBaWI2UWZwc01mTk9DVkR2bEZBIn0.eyJleHAiOjE2ODY5ODY3MDcsImlhdCI6MTY4Njk4NTgwNywiYXV0aF90aW1lIjoxNjg2OTg1ODA3LCJqdGkiOiIwNmQwZGJiOC02Nzc1LTRkMDYtOWEzMy01NDY0NmQ1NjJkYTAiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODEvcmVhbG1zL2l0LWFjYWRlbXkiLCJzdWIiOiJmZDM1OTY2Yy1mZGY1LTRhOTUtODJkMi1kZTUzMDJlNWUwNmYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhY2FkZW15Iiwic2Vzc2lvbl9zdGF0ZSI6Ijc1OWUwMzA4LWVhMWYtNDhlZS05ZDQwLWZhZWRmODgwN2VkYSIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cDovL2xvY2FsaG9zdDo4MDgwIl0sInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUiLCJzaWQiOiI3NTllMDMwOC1lYTFmLTQ4ZWUtOWQ0MC1mYWVkZjg4MDdlZGEiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InRlc3QgdGVzdCIsInByZWZlcnJlZF91c2VybmFtZSI6InRlc3QiLCJnaXZlbl9uYW1lIjoidGVzdCIsImZhbWlseV9uYW1lIjoidGVzdCIsImVtYWlsIjoidGVzdEBnLmNvbSIsImFjYWRlbXkiOnsicm9sZXMiOlsiU1RVREVOVCJdfX0.IvHSHKL9izMEcxOXqUVUUUEgSUB_R8WKcigx1xWXouZWUIy3fkG7aRLNLfWGYgCYi4wS3tIntThnM3TmIBgeYBv_qf22X0JsbD5Z_tMoTHL_a5SviNds5t_5sFfU2Jj3LUmRTwRR-IBqUSmfGU5bMe8vg9sNTKYH9NncpnoRdK8DMS805r1Nfvuj9JU1tCL4CjJLGyKekecEYhm766b_Q2fPtgMhu4Chfi9oibl9FzumArtlGYF67T6uEzberNc4nRlruJsOHHnDIuo299X8calO7Sh1FA9iA7DIXS9fFsni8W0skUY_wn-Uxd-1wWIUq8kQKj9QYdlZE6e-Pg2Wag"

	retrospectResult, err := client.RetrospectToken(context.Background(), userAccessToken, clientId, clientSecret, realm)

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
