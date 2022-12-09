package main

import (
	"regexp"
	"testing"
)

func TestResolveWorkflow(t *testing.T) {
	repo := Repo{Name: "name", Owner: "owner", Workflow: ""}
	repo2 := Repo{Name: "name", Owner: "owner", Workflow: "OverrideWorkflow"}
	appconf := AppConfig{Repos: []Repo{repo, repo2}, DefaultWorkflow: "DefaultWorkflow"}

	want := regexp.MustCompile(`\b` + appconf.DefaultWorkflow + `\b`)
	want2 := regexp.MustCompile(`\b` + repo2.Workflow + `\b`)

	result := resolveWorkflow(&repo, &appconf)
	result2 := resolveWorkflow(&repo2, &appconf)

	if !want.MatchString(result) {
		t.Fatalf(`Hello("Gladys") = %q, want match for %#q, nil`, result, want)
	}

	if !want2.MatchString(result2) {
		t.Fatalf(`Hello("Gladys") = %q, want match for %#q, nil`, result, want)
	}

}
