package config

import "runtime"

// Build holds information about the current build, including version control and runtime information
var Build = &BuildConfig{
	OS: runtime.GOOS,
}

// BuildConfig contains information about the build
// This is used for debugging purposes
type BuildConfig struct {
	OS string
}

// IsMac is a helper function to check if the build is for macOS
func (b *BuildConfig) IsMac() bool {
	return b.OS == "darwin"
}

// IsWindows will return true if the current OS is detected as windows
func (b *BuildConfig) IsWindows() bool {
	return b.OS == "windows"
}

// IsLinux returns true if running on a Linux distro
func (b *BuildConfig) IsLinux() bool {
	return b.OS == "linux"
}

// GoVersion is a helper function to get the current Go version
// All it does is call the builtin function
func (b *BuildConfig) GoVersion() string {
	return runtime.Version()
}
