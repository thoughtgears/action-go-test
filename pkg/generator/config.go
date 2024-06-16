package generator

import (
	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	ProjectID   string
	ServiceName string
}

func NewFromInputs(action *githubactions.Action) (*Config, error) {
	c := Config{
		ProjectID:   action.GetInput("gcp_project_id"),
		ServiceName: action.GetInput("service_name"),
	}

	return &c, nil
}
