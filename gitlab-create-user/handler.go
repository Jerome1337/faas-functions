package function

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/xanzy/go-gitlab"
)

type User struct {
	Username string
	Name     string
	Email    string
}

type Response struct {
	Code    int
	Message string
}

func Handle(req []byte) string {
	git := gitlab.NewClient(nil, os.Getenv("GITLAB_API_TOKEN"))

	if err := git.SetBaseURL(os.Getenv("GITLAB_API_URL")); err != nil {
		log.Fatal("Fail connecting to Gitlab API", err)
	}

	var newUser User

	if err := json.Unmarshal(req, &newUser); err != nil {
		log.Fatal("Failed unmarshal JSON", err)
	}

	u := &gitlab.CreateUserOptions{
		Email:         gitlab.String(newUser.Email),
		Name:          gitlab.String(newUser.Name),
		Username:      gitlab.String(newUser.Username),
		ResetPassword: gitlab.Bool(true),
	}

	user, _, err := git.Users.CreateUser(u)

	if err != nil {
		log.Fatal("Fail adding user", err)
	}

	return createResponse(http.StatusOK, "User "+user.Name+" is now created")
}

func createResponse(code int, message string) string {
	r := Response{Code: code, Message: message}

	rString, _ := json.Marshal(r)

	return string(rString)
}
