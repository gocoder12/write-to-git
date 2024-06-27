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
    
    // Name of the new branch
    newBranchName := "new-sample-ccbranch"
    
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)
    client := github.NewClient(tc)

    // Get the default branch
    repoInfo, _, err := client.Repositories.Get(ctx, owner, repo)
    if err != nil {
        log.Fatalf("Error getting repository info: %v", err)
    }
    defaultBranch := repoInfo.GetDefaultBranch()

    // Get the reference for the default branch
    ref, _, err := client.Git.GetRef(ctx, owner, repo, "refs/heads/"+defaultBranch)
    if err != nil {
        log.Fatalf("Error getting reference for default branch: %v", err)
    }

    // Get the commit that the default branch points to
    commit, _, err := client.Git.GetCommit(ctx, owner, repo, ref.Object.GetSHA())
    if err != nil {
        log.Fatalf("Error getting commit: %v", err)
    }

    // Create a tree with the new file
    content := "Hello, this is a sample file created by Go on a new branch!"
    tree, _, err := client.Git.CreateTree(ctx, owner, repo, commit.Tree.GetSHA(), []*github.TreeEntry{
        {
            Path:    github.String("sample.txt"),
            Mode:    github.String("100644"),
            Type:    github.String("blob"),
            Content: github.String(content),
        },
    })
    if err != nil {
        log.Fatalf("Error creating tree: %v", err)
    }

    // Create a new commit
    newCommit, _, err := client.Git.CreateCommit(ctx, owner, repo, &github.Commit{
        Message: github.String("Add sample file on new branch"),
        Tree:    tree,
        Parents: []*github.Commit{{SHA: commit.SHA}},
    })
    if err != nil {
        log.Fatalf("Error creating commit: %v", err)
    }

    // Create the new branch
    _, _, err = client.Git.CreateRef(ctx, owner, repo, &github.Reference{
        Ref:    github.String("refs/heads/" + newBranchName),
        Object: &github.GitObject{SHA: newCommit.SHA},
    })
    if err != nil {
        log.Fatalf("Error creating new branch: %v", err)
    }

    log.Printf("Created new branch '%s' with sample file", newBranchName)
}
