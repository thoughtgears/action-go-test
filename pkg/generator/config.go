package generator

import (
	"fmt"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

var allowedEnvironments = []string{"dev", "staging", "prod"}

type Config struct {
	ProjectID   string
	Environment string
}

func NewFromInputs(action *githubactions.Action) (*Config, error) {
	env := action.GetInput("environment")

	if !strings.Contains(strings.Join(allowedEnvironments, ","), env) {
		return nil, fmt.Errorf("invalid environment: %s. Allowed values: %s", env, strings.Join(allowedEnvironments, ", "))
	}

	c := Config{
		ProjectID:   action.GetInput("gcp_project_id"),
		Environment: env,
	}

	return &c, nil
}
