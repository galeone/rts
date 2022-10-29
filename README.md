# RTS: Request to Struct

[![GoDoc](https://godoc.org/github.com/galeone/rts?status.svg)](https://godoc.org/github.com/galeone/rts)
[![Build Status](https://travis-ci.org/galeone/rts.svg?branch=master)](https://travis-ci.org/galeone/rts)

Generate Go structs definitions from JSON server responses.

RTS defines type names using the specified lines in the route file and skipping numbers.
e.g: a request to a route like `/users/1/posts` generates `type UsersPosts`

It supports *parameters*: a line like `/users/:user/posts/:pid 1 200` generates `type UsersUserPostsPid` from the response to the request `GET /users/1/posts/200`.

RTS supports headers personalization as well, thus it can be used to generate types from responses protected by some authorization method

## Install

### CLI Application

`go get -u github.com/galeone/rts/cmd/rts`

#### CLI Usage

```
rts [options]
  -headers string
    	Headers to add in every request
  -help
    	prints this help
  -insecure
    	Disables TLS Certificate check for HTTPS, use in case HTTPS Server Certificate is signed by an unknown authority
  -out string
    	Output file. Stdout is used if not specified
  -pkg string
    	Package name (default "main")
  -routes string
    	Routes to request. One per line (default "routes.txt")
  -server string
    	sets the server address (default "http://localhost:9090")
  -substruct
    	Creates types for sub-structs
```

#### Examples

You can invoke `rts` piping from stdin a single JSON (anonymous) and get it converted to a go structure

```
echo '  {
    "Book Id": 30558257,
    "Title": "Unsouled (Cradle, #1)",
    "Author": "Will Wight",
    "Author l-f": "Wight, Will",
    "Additional Authors": "",
    "BCID": ""
  }' | ./rts
```

obtaining

```go
package main

type Foo1 struct {
        Additional_Authors string `json:"Additional Authors"`
        Author             string `json:"Author"`
        Author_l_f         string `json:"Author l-f"`
        Bcid               string `json:"BCID"`
        Book_Id            int64  `json:"Book Id"`
        Title              string `json:"Title"`
}
```

Or you can define a more complex scenario, definining the `routes.txt` file with a line for each (parametric) request and use it as shown below.

*routes.txt*:

```txt
/
/repos/:user/:repo galeone igor
```

Run:

```
rts -server https://api.github.com -pkg example
```

Returns:

```go
package example

type Foo1 struct {
	AuthorizationsURL                string `json:"authorizations_url"`
	CodeSearchURL                    string `json:"code_search_url"`
	CommitSearchURL                  string `json:"commit_search_url"`
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
	LabelSearchURL                   string `json:"label_search_url"`
	NotificationsURL                 string `json:"notifications_url"`
	OrganizationRepositoriesURL      string `json:"organization_repositories_url"`
	OrganizationTeamsURL             string `json:"organization_teams_url"`
	OrganizationURL                  string `json:"organization_url"`
	PublicGistsURL                   string `json:"public_gists_url"`
	RateLimitURL                     string `json:"rate_limit_url"`
	RepositorySearchURL              string `json:"repository_search_url"`
	RepositoryURL                    string `json:"repository_url"`
	StarredGistsURL                  string `json:"starred_gists_url"`
	StarredURL                       string `json:"starred_url"`
	TopicSearchURL                   string `json:"topic_search_url"`
	UserOrganizationsURL             string `json:"user_organizations_url"`
	UserRepositoriesURL              string `json:"user_repositories_url"`
	UserSearchURL                    string `json:"user_search_url"`
	UserURL                          string `json:"user_url"`
}

type ReposUserRepo struct {
	AllowForking             bool               `json:"allow_forking"`
	ArchiveURL               string             `json:"archive_url"`
	Archived                 bool               `json:"archived"`
	AssigneesURL             string             `json:"assignees_url"`
	BlobsURL                 string             `json:"blobs_url"`
	BranchesURL              string             `json:"branches_url"`
	CloneURL                 string             `json:"clone_url"`
	CollaboratorsURL         string             `json:"collaborators_url"`
	CommentsURL              string             `json:"comments_url"`
	CommitsURL               string             `json:"commits_url"`
	CompareURL               string             `json:"compare_url"`
	ContentsURL              string             `json:"contents_url"`
	ContributorsURL          string             `json:"contributors_url"`
	CreatedAt                string             `json:"created_at"`
	DefaultBranch            string             `json:"default_branch"`
	DeploymentsURL           string             `json:"deployments_url"`
	Description              string             `json:"description"`
	Disabled                 bool               `json:"disabled"`
	DownloadsURL             string             `json:"downloads_url"`
	EventsURL                string             `json:"events_url"`
	Fork                     bool               `json:"fork"`
	Forks                    int64              `json:"forks"`
	ForksCount               int64              `json:"forks_count"`
	ForksURL                 string             `json:"forks_url"`
	FullName                 string             `json:"full_name"`
	GitCommitsURL            string             `json:"git_commits_url"`
	GitRefsURL               string             `json:"git_refs_url"`
	GitTagsURL               string             `json:"git_tags_url"`
	GitURL                   string             `json:"git_url"`
	HasDownloads             bool               `json:"has_downloads"`
	HasIssues                bool               `json:"has_issues"`
	HasPages                 bool               `json:"has_pages"`
	HasProjects              bool               `json:"has_projects"`
	HasWiki                  bool               `json:"has_wiki"`
	Homepage                 string             `json:"homepage"`
	HooksURL                 string             `json:"hooks_url"`
	HTMLURL                  string             `json:"html_url"`
	ID                       int64              `json:"id"`
	IsTemplate               bool               `json:"is_template"`
	IssueCommentURL          string             `json:"issue_comment_url"`
	IssueEventsURL           string             `json:"issue_events_url"`
	IssuesURL                string             `json:"issues_url"`
	KeysURL                  string             `json:"keys_url"`
	LabelsURL                string             `json:"labels_url"`
	Language                 string             `json:"language"`
	LanguagesURL             string             `json:"languages_url"`
	License                  ReposUserRepo_sub1 `json:"license"`
	MergesURL                string             `json:"merges_url"`
	MilestonesURL            string             `json:"milestones_url"`
	MirrorURL                interface{}        `json:"mirror_url"`
	Name                     string             `json:"name"`
	NetworkCount             int64              `json:"network_count"`
	NodeID                   string             `json:"node_id"`
	NotificationsURL         string             `json:"notifications_url"`
	OpenIssues               int64              `json:"open_issues"`
	OpenIssuesCount          int64              `json:"open_issues_count"`
	Owner                    ReposUserRepo_sub2 `json:"owner"`
	Private                  bool               `json:"private"`
	PullsURL                 string             `json:"pulls_url"`
	PushedAt                 string             `json:"pushed_at"`
	ReleasesURL              string             `json:"releases_url"`
	Size                     int64              `json:"size"`
	SSHURL                   string             `json:"ssh_url"`
	StargazersCount          int64              `json:"stargazers_count"`
	StargazersURL            string             `json:"stargazers_url"`
	StatusesURL              string             `json:"statuses_url"`
	SubscribersCount         int64              `json:"subscribers_count"`
	SubscribersURL           string             `json:"subscribers_url"`
	SubscriptionURL          string             `json:"subscription_url"`
	SvnURL                   string             `json:"svn_url"`
	TagsURL                  string             `json:"tags_url"`
	TeamsURL                 string             `json:"teams_url"`
	TempCloneToken           interface{}        `json:"temp_clone_token"`
	Topics                   []string           `json:"topics"`
	TreesURL                 string             `json:"trees_url"`
	UpdatedAt                string             `json:"updated_at"`
	URL                      string             `json:"url"`
	Visibility               string             `json:"visibility"`
	Watchers                 int64              `json:"watchers"`
	WatchersCount            int64              `json:"watchers_count"`
	WebCommitSignoffRequired bool               `json:"web_commit_signoff_required"`
}

type ReposUserRepo_sub2 struct {
	AvatarURL         string `json:"avatar_url"`
	EventsURL         string `json:"events_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	GravatarID        string `json:"gravatar_id"`
	HTMLURL           string `json:"html_url"`
	ID                int64  `json:"id"`
	Login             string `json:"login"`
	NodeID            string `json:"node_id"`
	OrganizationsURL  string `json:"organizations_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	ReposURL          string `json:"repos_url"`
	SiteAdmin         bool   `json:"site_admin"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	Type              string `json:"type"`
	URL               string `json:"url"`
}

type ReposUserRepo_sub1 struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	NodeID string `json:"node_id"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
}
```

### Library

```go
import "github.com/galeone/rts"

byteFile, err := rts.Do(pkg, server, lines, headerMap)
```

# License

RTS: Request to Struct. Generates Go structs from a server response.
Copyright (C) 2016-2022 Paolo Galeone <nessuno@nerdz.eu>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
Exhibit B is not attached; this software is compatible with the
licenses expressed under Section 1.12 of the MPL v2.
