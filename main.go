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
	flag.StringVar(&token, "t", "", "Your token.")
	flag.StringVar(&user, "u", "", "User name.")
	flag.StringVar(&repo, "r", "", "Target repo.")
	flag.StringVar(&workflowName, "w", "build.yml", "Target workflow file name.")

	flag.Parse()
}

func main() {
	client := github.NewClientWithEnvProxy().WithAuthToken(token)

	actions := client.Actions
	var deleted int64
	listOpt := github.ListOptions{}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		runs, resp, err := actions.ListWorkflowRunsByFileName(ctx, user, repo, workflowName,
			&github.ListWorkflowRunsOptions{
				Status:              "failure",
				ExcludePullRequests: false,
				ListOptions:         listOpt,
			})
		cancel()
		if err != nil {
			log.Fatalln(err)
		}

		if len(runs.WorkflowRuns) == 0 {
			if resp.NextPage == 0 {
				break
			}

			listOpt.Page++
			continue
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

	}

	log.Printf("Delete %d failed workflow.\n", deleted)
}
