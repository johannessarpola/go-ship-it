package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

func useTokenAuth(token string) *http.Client {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return tc
}

func useBasicAuth(user string, pass string) *http.Client {
	tp := github.BasicAuthTransport{
		Username: user,
		Password: pass,
	}
	return tp.Client()
}

func createAuth() *http.Client {
	token := strings.TrimSpace(os.Getenv("GITHUB_ACCESS_TOKEN"))

	if len(token) > 0 {
		return useTokenAuth(token)
	} else {
		user := strings.TrimSpace(os.Getenv("GITHUB_USERNAME"))
		pass := strings.TrimSpace(os.Getenv("GITHUB_PASSWORD"))
		return useBasicAuth(user, pass)
	}
}

func createClient() (*github.Client, error) {

	ghes := strings.TrimSpace(os.Getenv("GITHUB_ENTEPRISE_HOST"))
	http := createAuth()

	if len(ghes) > 0 {
		return github.NewEnterpriseClient(ghes, ghes, http)
	} else {
		return github.NewClient(http), nil
	}

}

type Repo struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
}

type AppConfig struct {
	Repos []Repo `yaml:"repos"`
}

func readConfig() (AppConfig, error) {

	var appConfig AppConfig
	yamlPath := os.Getenv("APP_CONFIG_YAML_PATH")

	fmt.Printf("Loading YAML from %s\n", yamlPath)

	yamlFile, err := ioutil.ReadFile(yamlPath)
	err2 := yaml.Unmarshal(yamlFile, &appConfig)

	fmt.Printf("--- t:\n%v\n\n", appConfig)
	if err != nil || err2 != nil {
		cerr := errors.New("AppConfig initialization")
		if err != nil {
			cerr = errors.Wrap(err, "Could not open config file")
		}
		if err2 != nil {
			cerr = errors.Wrap(err2, "Could not unmarshal config")
		}
		return appConfig, cerr
	}

	return appConfig, nil
}

type Release struct {
	Title string
	Body  string
	Tag   string
}

type Module struct {
	Name     string
	Releases []Release
}

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load env variables")
	}

	client, err := createClient()
	if err != nil {
		log.Fatal("Could not create Github client", err)
	}

	appConfig, err := readConfig()

	if err != nil {
		log.Fatal(err)
	}

	user, _, err := client.Users.Get(ctx, "")
	fmt.Println(user.GetName())
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range appConfig.Repos {
		fmt.Println(repo)
		remote, _, err := client.Repositories.Get(ctx, repo.Owner, repo.Name)
		if err != nil {
			fmt.Printf("Could not find repository %s\n", repo)
		}

		releases, _, _ := client.Repositories.ListReleases(ctx, repo.Owner, repo.Name, nil)

		// module := Module {
		// 	Name: fmt.Sprintf("%s/%s", repo.Owner, repo.Name)§

		// }

		for _, release := range releases {
			fmt.Println(release.GetTagName())
			fmt.Println(release.GetBody())
		}

		fmt.Println(fmt.Printf("%s:%s", remote.GetName(), remote.GetURL()))
	}
	fmt.Println("Hello, World!")
}
