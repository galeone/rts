# RTS: Request to Struct

[![GoDoc](https://godoc.org/github.com/galeone/rts?status.svg)](https://godoc.org/github.com/galeone/rts)
[![Build Status](https://travis-ci.org/galeone/rts.svg?branch=master)](https://travis-ci.org/galeone/rts)

Generate Go structs definitions from JSON server responses.

RTS defines type names using the specified lines in the route file and skipping numbers.
e.g: a request to a route like `/users/1/posts` generates `type UsersPosts`

It supports *parameters*: a line like `/users/:user/posts/:pid 1 200` generates `type UsersUserPostsPid` from the response to the request `GET /users/1/posts/200`.

RTS supports headers personalization as well, thus it can be used to generate types from responses protected by some authorization method

Updated: 6/17/2016 by Krish Verma <https://github.com/kverma>

In case the JSON server is HTTPS with unknown certificate signing authority, pass the -insecure flag to disable TLS certificate check

# Install

## CLI Application

`go get -u github.com/galeone/rts/cmd/rts`

## Library

```go
import "github.com/galeone/rts"

byteFile, err := rts.Do(pkg, server, lines, headerMap)
```

# CLI Usage

```
rts [options]
  -headers string
    	Headers to add in every request
  -help
    	prints this help
  -out string
    	Output file. Stdout is used if not specified
  -pkg string
    	Package name (default "main")
  -routes string
    	Routes to request. One per line (default "routes.txt")
  -server string
    	sets the server address (default "http://localhost:9090")
  -insecure
        Disables TLS Certificate check for HTTPS, use in case HTTPS Server Certificate is signed by an unknown authority
```

## Example

*routes.txt*:
```
/
/repos/:user/:repo galeone igor
```

Run:
```
rts -server https://api.github.com -pkg example
```

Returns:

```
package example

type Foo1 struct {
	AuthorizationsURL                string `json:"authorizations_url"`
	CodeSearchURL                    string `json:"code_search_url"`
	CurrentUserAuthorizationsHTMLURL string `json:"current_user_authorizations_html_url"`
	CurrentUserRepositoriesURL       string `json:"current_user_repositories_url"`
	CurrentUserURL                   string `json:"current_user_url"`
	EmailsURL                        string `json:"emails_url"`
	EmojisURL                        string `json:"emojis_url"`
	EventsURL                        string `json:"events_url"`
	FeedsURL                         string `json:"feeds_url"`
	FollowersURL                     string `json:"followers_url"`
	FollowingURL                     string `json:"following_url"`
	GistsURL                         string `json:"gists_url"`
	HubURL                           string `json:"hub_url"`
	IssueSearchURL                   string `json:"issue_search_url"`
	IssuesURL                        string `json:"issues_url"`
	KeysURL                          string `json:"keys_url"`
	NotificationsURL                 string `json:"notifications_url"`
	OrganizationRepositoriesURL      string `json:"organization_repositories_url"`
	OrganizationURL                  string `json:"organization_url"`
	PublicGistsURL                   string `json:"public_gists_url"`
	RateLimitURL                     string `json:"rate_limit_url"`
	RepositorySearchURL              string `json:"repository_search_url"`
	RepositoryURL                    string `json:"repository_url"`
	StarredGistsURL                  string `json:"starred_gists_url"`
	StarredURL                       string `json:"starred_url"`
	TeamURL                          string `json:"team_url"`
	UserOrganizationsURL             string `json:"user_organizations_url"`
	UserRepositoriesURL              string `json:"user_repositories_url"`
	UserSearchURL                    string `json:"user_search_url"`
	UserURL                          string `json:"user_url"`
}

type ReposUserRepo struct {
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CloneURL         string      `json:"clone_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	CreatedAt        string      `json:"created_at"`
	DefaultBranch    string      `json:"default_branch"`
	DeploymentsURL   string      `json:"deployments_url"`
	Description      string      `json:"description"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	Fork             bool        `json:"fork"`
	Forks            int         `json:"forks"`
	ForksCount       int         `json:"forks_count"`
	ForksURL         string      `json:"forks_url"`
	FullName         string      `json:"full_name"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitURL           string      `json:"git_url"`
	HasDownloads     bool        `json:"has_downloads"`
	HasIssues        bool        `json:"has_issues"`
	HasPages         bool        `json:"has_pages"`
	HasWiki          bool        `json:"has_wiki"`
	Homepage         string      `json:"homepage"`
	HooksURL         string      `json:"hooks_url"`
	HTMLURL          string      `json:"html_url"`
	ID               int         `json:"id"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	Language         string      `json:"language"`
	LanguagesURL     string      `json:"languages_url"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	MirrorURL        interface{} `json:"mirror_url"`
	Name             string      `json:"name"`
	NetworkCount     int         `json:"network_count"`
	NotificationsURL string      `json:"notifications_url"`
	OpenIssues       int         `json:"open_issues"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	Owner            struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int    `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"owner"`
	Private          bool   `json:"private"`
	PullsURL         string `json:"pulls_url"`
	PushedAt         string `json:"pushed_at"`
	ReleasesURL      string `json:"releases_url"`
	Size             int    `json:"size"`
	SSHURL           string `json:"ssh_url"`
	StargazersCount  int    `json:"stargazers_count"`
	StargazersURL    string `json:"stargazers_url"`
	StatusesURL      string `json:"statuses_url"`
	SubscribersCount int    `json:"subscribers_count"`
	SubscribersURL   string `json:"subscribers_url"`
	SubscriptionURL  string `json:"subscription_url"`
	SvnURL           string `json:"svn_url"`
	TagsURL          string `json:"tags_url"`
	TeamsURL         string `json:"teams_url"`
	TreesURL         string `json:"trees_url"`
	UpdatedAt        string `json:"updated_at"`
	URL              string `json:"url"`
	Watchers         int    `json:"watchers"`
	WatchersCount    int    `json:"watchers_count"`
}

```

# License

RTS: Request to Struct. Generates Go structs from a server response.
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
Exhibit B is not attached; this software is compatible with the
licenses expressed under Section 1.12 of the MPL v2.
