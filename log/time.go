package log

import "time"

var (
	// TimestampFormat is the format that timestamps will be outputted in, including in JSON logs
	TimestampFormat = "02-01-06 15:04:05"
	// TimestampLocation is the location in which timestamp timezones are derived
	TimestampLocation = "Europe/London"
)

// GetTimezone converts TimestampLocation into a time.Location so that timestamps are properly outputted in the correct timezone
func GetTimezone() *time.Location {
	tz, err := time.LoadLocation(TimestampLocation)
	if err != nil {
		return time.UTC
	}

	return tz
}

// GetTimestamp returns the current time in the format of TimestampFormat, in the timezone specified by TimestampLocation
func GetTimestamp() string {
	return time.Now().In(GetTimezone()).Format(TimestampFormat)
}
