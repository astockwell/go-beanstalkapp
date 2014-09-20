// Copyright 2014 Alex Stockwell. All rights reserved.
//
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package beanstalkapp

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

const (
	libraryVersion     = "0.0"
	defaultAPIProtocol = "http://"
	defaultAPIURL      = "beanstalkapp.com/api/"
	userAgent          = "go-beanstalkapp/" + libraryVersion

	// Only use JSON (TODO: add XML support)
	defaultContentType = "application/json"
)

// A Client manages communication with the Beanstalkapp API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.  Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise.  BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Username is the HTTP basic auth username for API calls made by this Client.
	Username string

	// Password is the HTTP basic auth password for API calls made by this Client.
	Password string

	// User agent used when communicating with the GitHub API.
	UserAgent string

	// Rate specifies the current rate limit for the client as determined by the
	// most recent API call.  If the client is used in a multi-user application,
	// this rate may not always be up-to-date.  Call RateLimit() to check the
	// current rate.
	// Rate Rate

	// Services used for talking to different parts of the GitHub API.
	// Activity      *ActivityService
	// Gists         *GistsService
	// Git           *GitService
	// Gitignores    *GitignoresService
	// Issues        *IssuesService
	// Organizations *OrganizationsService
	// PullRequests  *PullRequestsService
	// Repositories  *RepositoriesService
	// Search        *SearchService
	// Users         *UsersService
}

// NewClient returns a new Beanstalkapp API client.  If a nil httpClient is
// provided, http.DefaultClient will be used.  To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the goauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	uploadURL, _ := url.Parse(uploadBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, UploadURL: uploadURL}
	c.Activity = &ActivityService{client: c}
	c.Gists = &GistsService{client: c}
	c.Git = &GitService{client: c}
	c.Gitignores = &GitignoresService{client: c}
	c.Issues = &IssuesService{client: c}
	c.Organizations = &OrganizationsService{client: c}
	c.PullRequests = &PullRequestsService{client: c}
	c.Repositories = &RepositoriesService{client: c}
	c.Search = &SearchService{client: c}
	c.Users = &UsersService{client: c}
	return c
}
