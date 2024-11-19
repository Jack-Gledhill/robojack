package presence

import (
	"google.golang.org/api/calendar/v3"
)

var (
	// FocusedBase is shown when I'm focusing on studying and don't want to be disturbed
	FocusedBase = Presence{
		State:      "Please do not disturb",
		Details:    "Focused on studying",
		LargeImage: "denied",
		LargeText:  "Focused Study",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
)

// FocusTime processes a focusTime calendar event and shows it as a focused study session
func FocusTime(e *calendar.Event) (*Presence, error) {
	base := FocusedBase
	err := base.TemplateData(TemplateData{Event: e})
	if err != nil {
		return nil, err
	}

	return &base, nil
}
