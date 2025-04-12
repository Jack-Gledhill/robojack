package debug

import (
	"os"
	"runtime/debug"
	"time"
)

// Repository is a URL for the bot's source repository
const Repository = "https://github.com/Jack-Gledhill/robojack"

const (
	defaultHash = "unknown"
	defaultRef  = "dev"
	revisionKey = "vcs.revision"
	modifiedKey = "vcs.modified"
	timeKey     = "vcs.time"
)

// Git is populated at runtime with source control information
var Git = &GitInfo{
	Commit: CommitInfo{
		Hash: defaultHash,
	},
	Ref:        getGitRef(),
	Repository: Repository,
}

// GitInfo contains information about the version control state of the build
// This is useful for seeing the current state of the codebase
type GitInfo struct {
	Commit     CommitInfo `json:"commit"`
	Ref        string     `json:"ref"`
	Repository string     `json:"repository"`
}

// CommitInfo contains info about the current commit that the codebase is checked out as
type CommitInfo struct {
	Hash      string    `json:"hash"`
	Modified  bool      `json:"modified"`
	Timestamp time.Time `json:"timestamp"`
}

func init() {
	if i, ok := debug.ReadBuildInfo(); ok {
		for _, s := range i.Settings {
			switch s.Key {
			case revisionKey:
				Git.Commit.Hash = s.Value
			case modifiedKey:
				Git.Commit.Modified = s.Value == "true"
			case timeKey:
				t, err := time.Parse(time.RFC3339, s.Value)
				if err == nil {
					Git.Commit.Timestamp = t
				}
			}
		}
	}
}

func getGitRef() string {
	ref := os.Getenv("GIT_REF")
	if ref == "" {
		return defaultRef
	}

	return ref
}
