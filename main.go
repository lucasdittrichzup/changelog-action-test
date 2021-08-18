package main

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type tagData struct {
	tagName     string
	link        string
	publishedAt *github.Timestamp
}

type issuesData struct {
	title       string
	issueNumber int
	link        string
}

type pullsData struct {
	title        string
	prNumber     int
	link         string
	assigneeUser string
	assigneeLink string
}

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

	previousRelease := getPreviousRelease()
	nextRelease := getNextRelease()

	// Seleciona todas as issues fechadas depois da data de criação da tag
	issues, _, err := client.Issues.ListByRepo(
		context.Background(),
		owner,
		repo,
		&github.IssueListByRepoOptions{State: "closed", Since: previousRelease.publishedAt.Time},
	)
	if err != nil {
		log.Fatalf("error listing issues: %v", err)
	}

	// Coloca todos os títulos das issues elegíveis dentro do slice para uso posterior
	var allIssuesTitle []issuesData
	for _, issue := range issues {
		if issue.ClosedAt.After(previousRelease.publishedAt.Time) && issue.PullRequestLinks == nil && issue.ClosedAt.Before(nextRelease.publishedAt.Time) {
			filterIssue := issuesData{
				title:       *issue.Title,
				issueNumber: *issue.Number,
				link:        *issue.HTMLURL,
			}
			allIssuesTitle = append(allIssuesTitle, filterIssue)
		}
	}

	// Seleciona todas as pr fechadas
	prs, _, err := client.PullRequests.List(
		ctx,
		owner,
		repo,
		&github.PullRequestListOptions{State: "closed"},
	)
	if err != nil {
		log.Fatalf("error listing prs: %v", err)
	}

	// Filtra as prs mergeadas após a data de criação da tag
	// TODO: abrir issue no repo go-github, pois retorna erro ao usar o campo name da struct de user
	var mergedPulls []pullsData
	for _, pr := range prs {
		if pr.MergedAt.After(previousRelease.publishedAt.Time) && pr.MergedAt.Before(nextRelease.publishedAt.Time) {
			filterPull := pullsData{
				title:        *pr.Title,
				prNumber:     *pr.Number,
				link:         *pr.HTMLURL,
				assigneeUser: *pr.User.Login,
				assigneeLink: *pr.User.HTMLURL,
			}
			mergedPulls = append(mergedPulls, filterPull)
		}
	}
}

func getNextRelease() tagData {
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

	nrData := tagData{
		tagName:     *lastTag.TagName,
		link:        *lastTag.HTMLURL,
		publishedAt: lastTag.PublishedAt,
	}

	return nrData
}

func getPreviousRelease() tagData {
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
	tags, _, err := client.Repositories.ListReleases(
		context.Background(),
		owner,
		repo,
		&github.ListOptions{},
	)
	if err != nil {
		log.Fatalf("error getting the last tag: %v", err)
	}

	previousRelease := tags[1]

	prData := tagData{
		tagName:     *previousRelease.TagName,
		link:        *previousRelease.HTMLURL,
		publishedAt: previousRelease.PublishedAt,
	}

	return prData
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
