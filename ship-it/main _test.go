package main

import (
	"regexp"
	"testing"
)

func TestResolveWorkflow(t *testing.T) {
	repo := RepoConfig{Name: "name", Owner: "owner", Workflow: ""}
	repo2 := RepoConfig{Name: "name", Owner: "owner", Workflow: "OverrideWorkflow"}
	appconf := AppConfig{Repos: []RepoConfig{repo, repo2}, DefaultWorkflow: "DefaultWorkflow"}

	want := regexp.MustCompile(`\b` + appconf.DefaultWorkflow + `\b`)
	want2 := regexp.MustCompile(`\b` + repo2.Workflow + `\b`)

	result := resolveWorkflow(&repo, appconf.DefaultWorkflow)
	result2 := resolveWorkflow(&repo2, appconf.DefaultWorkflow)

	if !want.MatchString(result) {
		t.Fatalf(`ResolveWorkflow = %q, want match for %#q, nil`, result, want)
	}

	if !want2.MatchString(result2) {
		t.Fatalf(`ResolveWorkflow = %q, want match for %#q, nil`, result, want)
	}

}
