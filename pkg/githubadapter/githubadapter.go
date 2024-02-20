package githubadapter

import (
	"context"
	"strings"

	"github.com/google/go-github/v59/github"
	"github.com/patrik-koska-clone/nixconfig-go/pkg/config"
	"golang.org/x/oauth2"
)

var (
	ctx = context.Background()
)

type GithubAdapter struct {
	Client *github.Client
	CR     DownloadOpts
}

type DownloadOpts struct {
	FilePath string
	Repo     string
	Owner    string
}

func New(c config.Config) GithubAdapter {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: c.Token,
		},
	)

	httpClient := oauth2.NewClient(ctx, ts)

	github := GithubAdapter{
		Client: github.NewClient(httpClient),
	}

	return github

}

func (g *GithubAdapter) SearchAndDownload(
	page, perPage int,
	fileName string) ([]string, []string, error) {

	var (
		base64Content  string
		downloadURL    string
		downloadURLs   []string
		base64Contents []string
		err            error
		results        *github.CodeSearchResult
		opts           *github.SearchOptions
	)

	opts = &github.SearchOptions{
		ListOptions: github.ListOptions{Page: page, PerPage: perPage},
	}

	results, _, err = g.Client.Search.Code(ctx, "filename:"+fileName, opts)
	if err != nil {
		return base64Contents, downloadURLs, err
	}

	for _, result := range results.CodeResults {
		g.assignDownloadOpts(result)

		base64Content, downloadURL, err = g.getContent()
		if err != nil {
			return base64Contents, downloadURLs, err
		}

		base64Contents = append(base64Contents, base64Content)
		downloadURLs = append(downloadURLs, downloadURL)

	}

	return base64Contents, downloadURLs, nil
}

func (g GithubAdapter) getContent() (string, string, error) {

	fileContent, _, _, err := g.Client.Repositories.GetContents(
		ctx,
		g.CR.Owner,
		g.CR.Repo,
		g.CR.FilePath,
		&github.RepositoryContentGetOptions{})

	if err != nil {
		return "", "", err
	}

	if *fileContent.Content != "" && *fileContent.DownloadURL != "" {
		return *fileContent.Content, *fileContent.DownloadURL, nil
	} else {
		return "", "", err
	}
}

func (g *GithubAdapter) assignDownloadOpts(result *github.CodeResult) {
	if result != nil &&
		result.Repository != nil &&
		result.Repository.FullName != nil &&
		*result.Repository.FullName != "" &&
		result.Path != nil &&
		*result.Path != "" {

		fullName := *result.Repository.FullName
		parts := strings.SplitN(fullName, "/", 2)
		if len(parts) == 2 {
			g.CR.Owner = parts[0]
			g.CR.Repo = parts[1]
		}

		g.CR.FilePath = *result.Path
	}
}
