package main

import (
	"context"
	"log"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	// Replace with your GitHub personal access token
	token := ""

	// Replace with your GitHub username and repository name
	owner := ""
	repo := ""

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// File content
	content := []byte("Hello, this is a sample file created by Go!")

	// Create or update the file in the repository
	_, _, err := client.Repositories.CreateFile(ctx, owner, repo, "sample.txt", &github.RepositoryContentFileOptions{
		Message: github.String("Add sample file"),
		Content: content,
	})

	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	log.Println("File created and pushed successfully!")
}
