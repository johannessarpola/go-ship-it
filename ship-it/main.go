package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
	Name     string `yaml:"name"`
	Owner    string `yaml:"owner"`
	Workflow string `yaml:"workflow"`
}

type AppConfig struct {
	DefaultWorkflow string `yaml:"workflow"`
	Port            string `yaml:"port"`
	Host            string `yaml:"host"`
	Repos           []Repo `yaml:"repos"`
}

func resolveWorkflow(repo *Repo, app *AppConfig) string {
	if len(repo.Workflow) == 0 {
		return app.DefaultWorkflow
	} else {
		return repo.Workflow
	}
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
	ID   string `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Repository struct {
	Name     string    `json:"name"`
	Releases []Release `json:"releases"`
}

func convertRelease(release *github.RepositoryRelease) Release {
	return Release{
		ID:   strconv.FormatInt(release.GetID(), 10),
		Name: release.GetName(),
		Tag:  release.GetTagName(),
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/repositories", getRepositoryReleases)

	return router
}

func getRepositoryReleases(c *gin.Context) {
	ctx := context.Background()

	var repositories []Repository

	for _, repo := range appConfig.Repos {
		releases, _, _ := githubClient.Repositories.ListReleases(ctx, repo.Owner, repo.Name, nil) // TODO Handle errors
		var convertedReleases []Release                                                           // TODO Paginate so it is not all the releases
		for _, release := range releases {
			convertedReleases = append(convertedReleases, convertRelease(release))
		}

		r := Repository{
			Name:     repo.Name,
			Releases: convertedReleases,
		}
		repositories = append(repositories, r)
	}
	c.IndentedJSON(http.StatusOK, repositories)
}

var githubClient *github.Client
var appConfig AppConfig

func main() {

	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load env variables")
	}

	githubClient, err = createClient()
	if err != nil {
		log.Fatal("Could not create Github client", err)
	}

	appConfig, err = readConfig()

	if err != nil {
		log.Fatal(err)
	}

	router := setupRouter()
	router.Run(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port))

	user, _, err := githubClient.Users.Get(ctx, "")
	fmt.Println(user.GetName())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, World!")
}
