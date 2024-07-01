package main

import (
    "context"
    "fmt"
    "log"
    "encoding/base64"

    "github.com/google/go-github/v39/github"
    "golang.org/x/oauth2"
)

func main() {
    // Replace with your GitHub personal access token
    token := "YOUR_GITHUB_TOKEN"
    
    // Replace with your GitHub username and repository name
    owner := "YOUR_USERNAME"
    repo := "YOUR_REPO"
    
    // Nested folder structure
    folderPath := "abc/xyz/pqr"
    fileName := "sample.txt"
    filePath := fmt.Sprintf("%s/%s", folderPath, fileName)
    
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)
    client := github.NewClient(tc)

    // File content
    content := []byte("Hello, this is a sample file created or updated by Go in nested folders!")
    
    // Encode content to base64
    encodedContent := base64.StdEncoding.EncodeToString(content)
    
    // Check if the file already exists
    fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filePath, &github.RepositoryContentGetOptions{})
    
    var sha *string
    if err == nil && fileContent != nil {
        // File exists, get its SHA
        sha = fileContent.SHA
        log.Println("File already exists. Updating content...")
    } else {
        log.Println("File does not exist. Creating new file in nested folders...")
    }

    // Create or update the file in the repository
    _, _, err = client.Repositories.CreateFile(
        ctx, 
        owner, 
        repo, 
        filePath,
        &github.RepositoryContentFileOptions{
            Message: github.String(fmt.Sprintf("Add or update sample file in %s folder", folderPath)),
            Content: []byte(encodedContent),
            SHA:     sha, // This is nil for new files, and set for updates
        },
    )

    if err != nil {
        log.Fatalf("Error creating or updating file: %v", err)
    }

    log.Printf("File created or updated successfully in %s folder!", folderPath)
}