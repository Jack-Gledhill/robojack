package config

import (
	"os"
	"runtime/debug"
	"time"
)

const (
	Repository = "https://github.com/Jack-Gledhill/robojack"
	// RevisionKey maps to a string of the latest commit hash
	RevisionKey = "vcs.revision"
	// ModifiedKey will be "true" if the repo was modified
	ModifiedKey = "vcs.modified"
	// TimeKey is the time of the latest commit in the RFC3339 format
	TimeKey = "vcs.time"
)

var Git = GetGitConfig()

// GitConfig contains information about the version control state of the build
// This is useful for seeing the current state of the codebase
type GitConfig struct {
	Modified   bool
	Ref        string
	Repository string
	Revision   string
	Timestamp  time.Time
}

// GetGitConfig fetches version control information from the build
// This info is added by the Go compiler at build time
// This allows us to get the commit hash, whether the repo was modified, and the time of the commit
func GetGitConfig() *GitConfig {
	ref := os.Getenv("GIT_REF")
	if ref == "" {
		ref = "dev"
	}

	commit := GitConfig{
		Ref:        ref,
		Repository: Repository,
		Revision:   "unknown",
	}

	if i, ok := debug.ReadBuildInfo(); ok {
		for _, s := range i.Settings {
			switch s.Key {
			case RevisionKey:
				commit.Revision = s.Value
			case ModifiedKey:
				commit.Modified = s.Value == "true"
			case TimeKey:
				t, err := time.Parse(time.RFC3339, s.Value)
				if err == nil {
					commit.Timestamp = t
				}
			}
		}
	}

	return &commit
}
