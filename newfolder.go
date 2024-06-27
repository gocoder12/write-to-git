package main

import (
	"context"
	"encoding/base64"
	"fmt"
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

	// Name of the folder and file
	folderName := "sample-folder"
	fileName := "sample.txt"
	filePath := fmt.Sprintf("%s/%s", folderName, fileName)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// File content
	content := []byte("Hello, this is a sample file created or updated by Go in a folder!")

	// Encode content to base64
	encodedContent := base64.StdEncoding.EncodeToString(content)

	fmt.Println("encodedContent===", encodedContent)
	// Check if the file already exists
	fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filePath, &github.RepositoryContentGetOptions{})

	var sha *string
	if err == nil && fileContent != nil {
		// File exists, get its SHA
		sha = fileContent.SHA
		log.Println("File already exists. Updating content...")
	} else {
		log.Println("File does not exist. Creating new file...")
	}

	// Create or update the file in the repository
	_, _, err = client.Repositories.CreateFile(
		ctx,
		owner,
		repo,
		filePath,
		&github.RepositoryContentFileOptions{
			Message: github.String(fmt.Sprintf("Add or update sample file in %s folder", folderName)),
			Content: []byte(encodedContent),
			SHA:     sha, // This is nil for new files, and set for updates
		},
	)

	if err != nil {
		log.Fatalf("Error creating or updating file: %v", err)
	}

	log.Printf("File created or updated successfully in %s folder!", folderName)
}
