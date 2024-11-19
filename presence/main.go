package presence

import "google.golang.org/api/calendar/v3"

var (
	// GenericBusyBase will be used when an event is scheduled but is not matched to any particular base
	GenericBusyBase = Presence{
		State:      "I'm not available right now",
		Details:    "Busy",
		LargeImage: "stop",
		LargeText:  "Busy",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// GenericFreeBase is shown when there's nothing scheduled
	GenericFreeBase = Presence{
		State:      "I'm free right now, come say hi!",
		Details:    "Free",
		LargeImage: "chat",
		LargeText:  "Free",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
)

// GetPresence returns a presence object based on the event type and any additional metadata
func GetPresence(e *calendar.Event) (*Presence, error) {
	var base *Presence
	var err error

	switch e.EventType {
	case "default":
		base, err = Default(e)
	case "focusTime":
		base, err = FocusTime(e)
	case "outOfOffice":
		base, err = OutOfOffice(e)
	case "workingLocation":
		base, err = WorkingLocation(e)
	default:
		base = &GenericBusyBase
		err = base.TemplateData(TemplateData{Event: e})
	}
	if err != nil {
		return nil, err
	}

	err = base.AddTimestamps(e)
	if err != nil {
		return nil, err
	}

	return base, nil
}
