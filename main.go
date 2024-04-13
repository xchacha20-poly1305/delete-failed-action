package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/google/go-github/v61/github"
)

var (
	token        string
	user, repo   string
	workflowName string
)

const timeout = 10 * time.Second

func init() {
	flag.StringVar(&token, "t", "", "Your token")
	flag.StringVar(&user, "u", "", "User name")
	flag.StringVar(&repo, "r", "", "Target repo")
	flag.StringVar(&workflowName, "w", "", "Target workflow file name")

	flag.Parse()
}

func main() {
	client := github.NewClientWithEnvProxy().WithAuthToken(token)

	actions := client.Actions
	var deleted int64

	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		runs, resp, err := actions.ListWorkflowRunsByFileName(ctx, user, repo, workflowName,
			&github.ListWorkflowRunsOptions{
				Status:              "failure",
				ExcludePullRequests: false,
				ListOptions:         github.ListOptions{},
			})
		cancel()
		if err != nil {
			log.Fatalln(err)
		}

		for _, run := range runs.WorkflowRuns {
			ctx, cancel = context.WithTimeout(context.Background(), timeout)
			_, err = actions.DeleteWorkflowRun(ctx, user, repo, *run.ID)
			cancel()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("Delete: ", *run.ID)
			deleted++
		}

		if resp.NextPage == 0 {
			break
		}
	}

	log.Printf("Delete %d failed workflow.\n", deleted)
}
