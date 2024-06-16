package generator

import (
	"context"

	"github.com/sethvargo/go-githubactions"
)

func Run(ctx context.Context, cfg *Config) error {
	githubactions.SetOutput("project_id", cfg.ProjectID)
	githubactions.SetOutput("service_name", cfg.ServiceName)
	return nil
}
