package generator

import (
	"log"
	"os"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
)

func TestNewFromInputs(t *testing.T) {
	envMap := map[string]string{
		"INPUT_GCP_PROJECT_ID": "my-project-1234",
		"INPUT_ENVIRONMENT":    "dev",
	}

	for key, value := range envMap {
		if err := os.Setenv(key, value); err != nil {
			log.Fatal(err)
		}
	}

	action := githubactions.New()

	cfg, err := NewFromInputs(action)

	assert.NoError(t, err)
	assert.Equal(t, envMap["INPUT_GCP_PROJECT_ID"], cfg.ProjectID)
	assert.Equal(t, envMap["INPUT_ENVIRONMENT"], cfg.Environment)

	for key, _ := range envMap {
		if err := os.Unsetenv(key); err != nil {
			log.Fatal(err)
		}
	}
}

func TestNewFromInputsWrongEnvironment(t *testing.T) {
	envMap := map[string]string{
		"INPUT_GCP_PROJECT_ID": "my-project-1234",
		"INPUT_ENVIRONMENT":    "super",
	}

	for key, value := range envMap {
		if err := os.Setenv(key, value); err != nil {
			log.Fatal(err)
		}
	}

	action := githubactions.New()

	_, err := NewFromInputs(action)

	assert.Error(t, err)
	assert.Equal(t, "invalid environment: super. Allowed values: dev, staging, prod", err.Error())

	for key, _ := range envMap {
		if err := os.Unsetenv(key); err != nil {
			log.Fatal(err)
		}
	}
}
