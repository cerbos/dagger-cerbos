// Module for working with open source Cerbos PDP
// Visit https://cerbos.dev and https://github.com/cerbos/cerbos for more information.
package main

import (
	"context"
)

type Cerbos struct{}

// Compile Cerbos policies and run any discovered tests. See https://docs.cerbos.dev/cerbos/latest/policies/compile.
func (c *Cerbos) Compile(
	ctx context.Context,
	// Directory containing the Cerbos policies.
	policyDir *Directory,
	// Compile results output format. Valid values are "tree", "list" or "json".
	// +optional
	outputFormat string,
	// Ignore schemas during compilation.
	// +optional
	ignoreSchemas bool,
	// Only compile without running tests.
	// +optional
	skipTests bool,
	// Regular expression matching the tests to run.
	// +optional
	run string,
	// Test output format. Valid values are "tree", "list", "json" or "junit".
	// +optional
	testOutputFormat string,
	// Produce execution traces on test failure.
	// +optional
	verboseFailures bool,
	// Cerbos version to use.
	// +optional
	// +default="latest"
	cerbosVersion string,
) (string, error) {
	args := []string{"compile"}

	if outputFormat != "" {
		args = append(args, "--output="+outputFormat)
	}

	if ignoreSchemas {
		args = append(args, "--ignore-schemas")
	}

	if skipTests {
		args = append(args, "--skip-tests")
	}

	if run != "" {
		args = append(args, "--run="+run)
	}

	if testOutputFormat != "" {
		args = append(args, "--test-output="+testOutputFormat)
	}

	if verboseFailures {
		args = append(args, "--verbose")
	}

	args = append(args, "/policies")

	return dag.Container().
		From("ghcr.io/cerbos/cerbos:"+cerbosVersion).
		WithMountedDirectory("/policies", policyDir).
		WithExec(args).
		Stdout(ctx)
}

func (c *Cerbos) Server(
	// Cerbos version to use.
	// +optional
	// +default="latest"
	cerbosVersion string,
	// List of configuration values to pass to Cerbos in the format config.key=value format.
	// +optional
	config []string,
	// Directory containing the Cerbos configuration file.
	// +optional
	configFile *File,
	// Directory containing the Cerbos policies.
	// +optional
	policyDir *Directory,
	// Cerbos log level.
	// +optional
	logLevel string,
) *Service {
	container := dag.Container().
		From("ghcr.io/cerbos/cerbos:" + cerbosVersion).
		WithExposedPort(3592).
		WithExposedPort(3593)

	if configFile != nil {
		container = container.WithMountedFile("/conf/.cerbos.yaml", configFile).WithEnvVariable("CERBOS_CONFIG", "/conf/.cerbos.yaml")
	}

	if policyDir != nil {
		container = container.WithMountedDirectory("/policies", policyDir)
	} else {
		container = container.WithMountedTemp("/policies")
	}

	if len(config) > 0 {
		args := make([]string, len(config)+1)
		args[0] = "server"
		for i, c := range config {
			args[i+1] = "--set=" + c
		}
		container = container.WithExec(args)
	}

	if logLevel != "" {
		container = container.WithEnvVariable("CERBOS_LOG_LEVEL", logLevel)
	}

	return container.AsService()
}
