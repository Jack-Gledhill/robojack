package presence

import "google.golang.org/api/calendar/v3"

// OutOfOfficeBase is used during working hours where I'm explicitly not working for a particular reason
var OutOfOfficeBase = Presence{
	State:      "Currently out of office, I'll be back soon",
	Details:    "{{.Event.Summary}}",
	LargeImage: "door",
	LargeText:  "Out of office",
	SmallImage: "robojack",
	SmallText:  "Powered by RoboJack",
}

// OutOfOffice processes the outOfOffice type of event
func OutOfOffice(e *calendar.Event) (*Presence, error) {
	base := OutOfOfficeBase
	err := base.TemplateData(TemplateData{Event: e})
	if err != nil {
		return nil, err
	}

	return &base, nil
}
