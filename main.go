package main

import (
	"fmt"
	"log"
	"os"

	"code.google.com/p/goauth2/oauth"

	"github.com/google/go-github/github"
)

func main() {
	options := &github.ActivityListStarredOptions{Sort: "created"}

	if len(os.Args) <= 1 {
		log.Fatalf("Usage: %s <GitHub username>", os.Args[0])
	}
	user := os.Args[1]

	client := newClient()

	for page := 1; ; page++ {
		options.Page = page

		starredRepos, res, err := client.Activity.ListStarred(user, options)
		if err != nil {
			log.Fatalf("ListStarred: %s", err)
		}

		log.Printf("page: %d/%d", page, res.LastPage)
		for _, starredRepo := range starredRepos {
			fmt.Println(*starredRepo.Repository.HTMLURL)
		}

		if page >= res.LastPage {
			break
		}
	}
}

func newClient() *github.Client {
	githubToken := os.Getenv("GITHUB_TOKEN")

	if githubToken != "" {
		oauthTransport := &oauth.Transport{
			Token: &oauth.Token{AccessToken: githubToken},
		}
		return github.NewClient(oauthTransport.Client())
	}

	return github.NewClient(nil)
}
