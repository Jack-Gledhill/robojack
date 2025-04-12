package debug

import "runtime"

// System is populated at runtime with system information
var System = &SystemInfo{
	Arch:        runtime.GOARCH,
	LogicalCPUs: runtime.NumCPU(),
	OS:          runtime.GOOS,
}

// SystemInfo contains information about the system the program is currently running on
type SystemInfo struct {
	Arch        string `json:"arch"`
	LogicalCPUs int    `json:"logical_cpus"`
	OS          string `json:"os"`
}

// IsMac is a helper function to check if the build is for macOS
func (s *SystemInfo) IsMac() bool {
	return s.OS == "darwin"
}

// IsWindows will return true if the current OS is detected as windows
func (s *SystemInfo) IsWindows() bool {
	return s.OS == "windows"
}

// IsLinux returns true if running on a Linux distro
func (s *SystemInfo) IsLinux() bool {
	return s.OS == "linux"
}
