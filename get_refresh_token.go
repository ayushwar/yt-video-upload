package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("YOUTUBE_CLIENT_ID"),
		ClientSecret: os.Getenv("YOUTUBE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/oauth/callback",  
		Scopes:       []string{youtube.YoutubeUploadScope},
		Endpoint:     google.Endpoint,
	}

	url := conf.AuthCodeURL(
	"state-token",
	oauth2.AccessTypeOffline,
	oauth2.ApprovalForce,
)
fmt.Println("Open this URL in browser:")
fmt.Println(url)

	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("REFRESH TOKEN:", token.RefreshToken)
		w.Write([]byte("Refresh token printed in terminal"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}