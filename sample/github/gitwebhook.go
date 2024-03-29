/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
	"golang.org/x/oauth2"
	webhooks "gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"

	ghclient "github.com/google/go-github/github"
)

const (
	// Environment variable containing json credentials
	envSecret = "GITHUB_SECRET"
	// this is what we tack onto each PR title if not there already
	titleSuffix = "I buy it"
)

// GithubHandler holds necessary objects for communicating with the Github.
type GithubHandler struct {
	client *ghclient.Client
	ctx    context.Context
}

type GithubSecrets struct {
	AccessToken string `json:"accessToken"`
	SecretToken string `json:"secretToken"`
}

// HandlePullRequest is invoked whenever a PullRequest is modified (created, updated, etc.)
func (handler *GithubHandler) HandlePullRequest(payload interface{}, header webhooks.Header) {
	glog.Info("Handling Pull Request")

	pl := payload.(github.PullRequestPayload)

	// Do whatever you want from here...
	title := pl.PullRequest.Title
	glog.Infof("GOT PR with Title: %q", title)

	// Check the title and if it contains 'looks pretty legit' leave it alone
	if strings.Contains(title, titleSuffix) {
		// already modified, leave it alone.
		return
	}

	newTitle := fmt.Sprintf("%s (%s)", title, titleSuffix)
	updatedPR := ghclient.PullRequest{
		Title: &newTitle,
	}
	newPR, response, err := handler.client.PullRequests.Edit(handler.ctx, pl.Repository.Owner.Login, pl.Repository.Name, int(pl.Number), &updatedPR)
	if err != nil {
		glog.Warningf("Failed to update PR: %s\n%s", err, response)
		return
	}
	if newPR.Title != nil {
		glog.Infof("New PR Title: %q", *newPR.Title)
	} else {
		glog.Infof("New PR title is nil")
	}
}

func main() {
	flag.Parse()
	// set the logs to stderr so kube will see them.
	flag.Lookup("logtostderr").Value.Set("true")
	githubSecrets := os.Getenv(envSecret)

	var credentials GithubSecrets
	err := json.Unmarshal([]byte(githubSecrets), &credentials)
	if err != nil {
		glog.Fatalf("Failed to unmarshal credentials: %s", err)
		return
	}

	// Set up the auth for being able to talk to Github. It's
	// odd that you have to also pass context around for the
	// calls even after giving it to client. But, whatever.
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: credentials.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := ghclient.NewClient(tc)

	h := &GithubHandler{
		client: client,
		ctx:    ctx,
	}

	hook := github.New(&github.Config{Secret: credentials.SecretToken})
	hook.RegisterEvents(h.HandlePullRequest, github.PullRequestEvent)

	err = webhooks.Run(hook, ":8080", "/")
	if err != nil {
		glog.Fatalf("Failed to run the webhook")
	}
}
