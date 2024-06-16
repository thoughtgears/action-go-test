package generator

import (
	"context"

	"github.com/sethvargo/go-githubactions"
)

func Run(ctx context.Context, cfg *Config) error {
	if cfg.Environment == "prod" {
		// Do something here
	}

	githubactions.SetOutput("gcp_project_id", cfg.ProjectID)

	return nil
}
