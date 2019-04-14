package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	. "github.com/logrusorgru/aurora"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func main() {
	// load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to github

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("LASTWEEK_GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	// query user

	ctx := context.Background()

	viewer, err := FetchViewer(client, ctx)
	if err != nil {
		log.Fatalf("Error querying for viewer: %v", err)
	}
	fmt.Printf("Recent PRs for %s\n\n", Bold(Cyan(viewer.Login)))

	// query prs

	prs, err := FetchPullReqs(client, ctx)
	if err != nil {
		log.Fatalf("Error querying for prs: %v", err)
	}
	for _, pr := range prs {
		fmt.Println(pr)
	}
}
