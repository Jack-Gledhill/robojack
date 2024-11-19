package env

import (
	"runtime"
	"runtime/debug"
	"time"
)

const (
	// GitPrefix is a prefix used for the below keys
	GitPrefix = "vcs."
	// RevisionKey maps to a string of the latest commit hash
	RevisionKey = GitPrefix + "revision"
	// ModifiedKey will be "true" if the repo was modified
	ModifiedKey = GitPrefix + "modified"
	// TimeKey is the time of the latest commit in the RFC3339 format
	TimeKey = GitPrefix + "time"
)

// Build holds information about the current build, including version control and runtime information
var Build = BuildInfo{
	Git: GetGitInfo(),
	OS:  runtime.GOOS,
}

// BuildInfo contains information about the build
// This is used for debugging purposes
type BuildInfo struct {
	Git GitInfo
	OS  string
}

// Mac is a helper function to check if the build is for macOS
func (b *BuildInfo) Mac() bool {
	return b.OS == "darwin"
}

// Windows will return true if the current OS is detected as windows
func (b *BuildInfo) Windows() bool {
	return b.OS == "windows"
}

// Linux returns true if running on a Linux distro
func (b *BuildInfo) Linux() bool {
	return b.OS == "linux"
}

// GoVersion is a helper function to get the current Go version
// All it does is call the builtin function
func (b *BuildInfo) GoVersion() string {
	return runtime.Version()
}

// GitInfo contains information about the version control state of the build
// This is useful for seeing the current state of the codebase
type GitInfo struct {
	Modified  bool
	Revision  string
	Timestamp time.Time
}

// GetGitInfo fetches version control information from the build
// This info is added by the Go compiler at build time
// This allows us to get the commit hash, whether the repo was modified, and the time of the commit
func GetGitInfo() GitInfo {
	commit := GitInfo{}

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

	return commit
}
