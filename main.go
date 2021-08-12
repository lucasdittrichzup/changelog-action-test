package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	token := getConfig("TOKEN")
	owner := getConfig("OWNER")
	repo := getConfig("REPO")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Pega a última tag do repositório
	lastTag, _, err := client.Repositories.GetLatestRelease(
		context.Background(),
		owner,
		repo,
	)
	if err != nil {
		log.Fatalf("error getting the last tag: %v", err)
	}

	// Seleciona todas as issues fechadas depois da data de criação da tag
	issues, _, err := client.Issues.ListByRepo(
		context.Background(),
		owner,
		repo,
		&github.IssueListByRepoOptions{State: "closed", Since: lastTag.CreatedAt.Time},
	)
	if err != nil {
		log.Fatalf("error listing issues: %v", err)
	}

	// Coloca todos os títulos das issues elegíveis dentro do slice para uso posterior
	var allIssuesTitle []string
	for _, issue := range issues {
		if issue.ClosedAt.After(lastTag.CreatedAt.Time) {
			allIssuesTitle = append(allIssuesTitle, *issue.Title)
			fmt.Println(allIssuesTitle)
		}
	}
}

func getConfig(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error while reading config file: %v", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("invalid type assertion")
	}

	return value
}
