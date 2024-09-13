package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"

    "github.com/google/go-github/v39/github"
    "golang.org/x/oauth2"
)

func main() {
    // Replace with your GitHub token
    token := "your-github-token"
    
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)

    // You might need to adjust this URL
    baseURL, _ := url.Parse("https://github.aexp.com/api/v3/")
    client, err := github.NewEnterpriseClient(baseURL.String(), baseURL.String(), tc)
    if err != nil {
        log.Fatalf("Error creating GitHub client: %v", err)
    }

    org := "amex-eng"
    repo := "self-svc-intermediate"
    path := "README.md" // Adjust this path as needed

    // First, let's check if we can access the repository
    repoInfo, resp, err := client.Repositories.Get(ctx, org, repo)
    if err != nil {
        if resp != nil {
            body, _ := io.ReadAll(resp.Body)
            log.Printf("Response body: %s", body)
            log.Printf("Status code: %d", resp.StatusCode)
        }
        log.Fatalf("Error accessing repository: %v", err)
    }

    log.Printf("Successfully accessed repository: %s", repoInfo.GetFullName())

    fileContent, resp, _, err := client.Repositories.GetContents(ctx, org, repo, path, nil)
    if err != nil {
        if resp != nil {
            body, _ := io.ReadAll(resp.Body)
            log.Printf("Response body: %s", body)
            log.Printf("Status code: %d", resp.StatusCode)
        }
        log.Fatalf("Error getting file content: %v", err)
    }

    content, err := fileContent.GetContent()
    if err != nil {
        log.Fatalf("Error decoding content: %v", err)
    }

    fmt.Printf("File content:\n%s\n", content)
}
