package function

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

type AnalyticsData struct {
	ProjectName string
	From        string
}

func Handle(req []byte) string {
	gitlab.NewClient(nil, "yourtokengoeshere")

	return fmt.Sprintf("The half of: %s", string(req))
}
