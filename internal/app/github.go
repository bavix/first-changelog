package app

import (
	"bytes"
	"context"
	"embed"
	"net/http"
	"strings"
	"text/template"

	"github.com/google/go-github/v58/github"
)

const name = "templates/changelog.tmpl"

//go:embed templates/*.tmpl
var changelogTmpl embed.FS

func GenChangelog(ctx context.Context, path string) string {
	splits := strings.SplitN(path, "/", 2)
	client := github.NewClient(http.DefaultClient)

	if len(splits) != 2 {
		return ""
	}

	tags := make([]tag, 0, 1024)

	i := 1
	for {
		releases, resp, err := client.
			Repositories.
			ListReleases(ctx, splits[0], splits[1], &github.ListOptions{
				PerPage: 500,
				Page:    i,
			})
		if err != nil {
			return ""
		}

		resp.Body.Close()

		if len(releases) == 0 {
			break
		}

		for _, release := range releases {
			if release.GetPrerelease() {
				continue
			}

			tags = append(tags, tag{
				tagName:   *release.TagName,
				body:      *release.Body,
				createdAt: release.CreatedAt.Time,
			})
		}

		i++
	}

	parse, err := template.ParseFS(changelogTmpl, name)
	if err != nil {
		return ""
	}

	var buffer bytes.Buffer

	if err = parse.Execute(&buffer, tags); err != nil {
		return ""
	}

	return buffer.String()
}
