package main

import (
	"context"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"github.com/shurcooL/githubv4"
)

// const PRS_LIMIT = 10

type User struct {
	Login      githubv4.String
	CreatedAt  githubv4.DateTime
	WebsiteURL githubv4.URI
}

func FetchViewer(client *githubv4.Client, ctx context.Context) (User, error) {
	var q struct {
		Viewer User
	}
	err := client.Query(ctx, &q, nil)
	if err != nil {
		return q.Viewer, err
	}

	return q.Viewer, nil

}

type PullRequest struct {
	Author struct {
		Login string
	}
	Number int
	Url    string
	// BodyText  string
	Closed    bool
	Merged    bool
	Title     string
	CreatedAt githubv4.DateTime
}

// Color returnsaurora's color
// based on the PR's state
func (pr PullRequest) Color() (clr Color) {
	clr = GreenFg
	if pr.Closed {
		clr = RedFg
	}
	if pr.Merged {
		clr = MagentaFg
	}
	return
}

func (pr PullRequest) String() string {
	prClr := pr.Color()
	num := fmt.Sprintf("%s%03d", Bold(Colorize("#", prClr)), Bold(Colorize(pr.Number, prClr)))
	return fmt.Sprintf("%s: %s by %s %-10s", num, Bold(pr.Title), Cyan(pr.Author.Login), pr.Url)
}

func FetchPullReqs(client *githubv4.Client, ctx context.Context) ([]PullRequest, error) {
	var q struct {
		Viewer struct {
			PullRequests struct {
				Nodes []PullRequest
			} `graphql:"pullRequests(last: 10, orderBy: { field: CREATED_AT, direction: ASC })"`
		}
	}

	err := client.Query(ctx, &q, nil)
	return q.Viewer.PullRequests.Nodes, err
}
