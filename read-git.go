package main

import (
    "context"
    "fmt"
    "log"

    "github.com/google/go-github/v39/github"
    "golang.org/x/oauth2"
)

func main() {
    // Replace with your GitHub personal access token
    token := "your-personal-access-token"
    
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)

    client := github.NewClient(tc)

    // Replace with the owner, repo, and file path you want to read
    owner := "octocat"
    repo := "Hello-World"
    path := "README.md"

    fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)
    if err != nil {
        log.Fatalf("Error getting file content: %v", err)
    }

    content, err := fileContent.GetContent()
    if err != nil {
        log.Fatalf("Error decoding content: %v", err)
    }

    fmt.Printf("File content:\n%s\n", content)
}
