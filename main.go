package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

func main() {
	token, found := os.LookupEnv("GH_TOOLS_TOKEN")
	if !found {
		log.Fatalln("GH_TOOLS_TOKEN env var is not set")
	}

	ctx := context.Background()

	// Create the http client.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	rd, _, err := client.Repositories.GetReadme(ctx, "appscode", "voyager", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rd.GetContent())
}
