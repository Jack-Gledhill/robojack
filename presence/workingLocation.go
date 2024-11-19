package presence

import "google.golang.org/api/calendar/v3"

var (
	// WorkingOnCampusBase is shown when I'm on campus but not doing anything specific
	WorkingOnCampusBase = Presence{
		State:      "Feel free to come say hi!",
		Details:    "Working in {{.Event.Summary}}",
		LargeImage: "computer",
		LargeText:  "Working",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// WorkingFromHomeBase is shown when I'm at home during working hours but not doing anything in particular
	WorkingFromHomeBase = Presence{
		State:      "Feel free to send me a message!",
		Details:    "Working from home",
		LargeImage: "computer",
		LargeText:  "Working",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
)

// WorkingLocation processes a workingLocation calendar event and shows it as a location
func WorkingLocation(e *calendar.Event) (*Presence, error) {
	var base Presence
	if e.Summary == "Home" {
		base = WorkingFromHomeBase
	} else {
		base = WorkingOnCampusBase
	}

	err := base.TemplateData(TemplateData{Event: e})
	if err != nil {
		return nil, err
	}

	return &base, nil
}
