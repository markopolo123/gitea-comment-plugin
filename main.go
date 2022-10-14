package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Payload struct {
	Body string `json:"body"`
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

func main() {
	var giteaToken string
	var giteaAddress string
	var comment string
	var repoOwner string
	var repoName string
	var prIndex int

	flag.StringVar(&giteaToken, "gitea-token", LookupEnvOrString("PLUGIN_GITEA_TOKEN", giteaToken), "API token for Gitea")
	flag.StringVar(&giteaAddress, "gitea-address", LookupEnvOrString("PLUGIN_GITEA_ADDRESS", giteaAddress), "Gitea URL")
	flag.StringVar(&comment, "comment", LookupEnvOrString("PLUGIN_COMMENT", comment), "Comment for Gitea")
	flag.StringVar(&repoOwner, "repo-owner", LookupEnvOrString("CI_REPO_OWNER", repoOwner), "Owner of the repository")
	flag.StringVar(&repoName, "repo-name", LookupEnvOrString("CI_REPO_NAME", repoName), "Name of the repository")
	flag.IntVar(&prIndex, "pr-index", LookupEnvOrInt("CI_COMMIT_PULL_REQUEST", prIndex), "Index of the PR")

	flag.Parse()

	if comment == "" {
		panic("You must provide a comment")
	}
	if giteaToken == "" {
		panic("You must provide a Gitea API Token")
	}
	if giteaAddress == "" {
		panic("You must provide a Gitea URL")
	}
	if repoOwner == "" {
		panic("You must provide an repo owner")
	}
	if repoName == "" {
		panic("You must provide a repo name")
	}
	if prIndex == 0 {
		panic("You must provide an index for PR")
	}

	data := Payload{
		Body: comment,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(payloadBytes)

	url := fmt.Sprintf("%s/api/v1/repos/%s/%s/issues/%d/comments?access_token=%s", giteaAddress, repoOwner, repoName, prIndex, giteaToken)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
