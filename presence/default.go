package presence

import "google.golang.org/api/calendar/v3"

const (
	// TTRPGColour is the colour ID for TTRPG game events
	TTRPGColour = "10"
	// MeetingColour is the colour ID for meeting events
	MeetingColour = "4"
)

var (
	// EventBases maps an event's colour to a particular base
	EventBases = map[string]Presence{
		"5":           PresentingBase,
		"3":           LectureBase,
		"1":           LabBase,
		"9":           TutorialBase,
		"11":          ExamBase,
		"2":           SocialBase,
		TTRPGColour:   TTRPGBase,
		MeetingColour: MeetingBase,
	}

	// PresentingBase for when I'm giving a presentation or helping out with a workshop
	PresentingBase = Presence{
		State:      "{{.Event.Summary}}",
		Details:    "Giving a presentation",
		LargeImage: "microphone",
		LargeText:  "Presenting",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// LectureBase is for when I'm sat in a lecture
	LectureBase = Presence{
		State:      "{{.Event.Summary}}",
		Details:    "Attending a lecture",
		LargeImage: "speaker",
		LargeText:  "Lecture",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// LabBase is used when I'm scheduled to be in a lab session
	LabBase = Presence{
		State:      "{{.Event.Summary}}",
		Details:    "Working in a lab",
		LargeImage: "toolbox",
		LargeText:  "Lab",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// TutorialBase is for when I'm in a tutorial session
	TutorialBase = Presence{
		State:      "{{.Event.Summary}}",
		Details:    "Working in a tutorial",
		LargeImage: "document",
		LargeText:  "Tutorial",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// ExamBase is for when I'm sitting an exam
	ExamBase = Presence{
		State:      "{{.Event.Summary}}",
		Details:    "Sitting an exam",
		LargeImage: "clock",
		LargeText:  "Exam",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// SocialBase is for when I'm attending a social event, such as with a society
	SocialBase = Presence{
		State:      "{{.Event.Summary}}",
		Details:    "At a social",
		LargeImage: "gamecontroller",
		LargeText:  "Social",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// MeetingBase is shown when I'm attending a meeting
	MeetingBase = Presence{
		State:      "I'll get back to you when I can!",
		Details:    "In a meeting",
		LargeImage: "briefcase",
		LargeText:  "In a meeting",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
	// TTRPGBase will be shown when I'm currently playing a TTRPG game
	TTRPGBase = Presence{
		State:      "Playing as {{.TTRPGMeta.Character}} in {{.Event.Summary}}",
		Details:    "Playing {{.TTRPGMeta.System}}",
		LargeImage: "die",
		LargeText:  "TTRPG",
		SmallImage: "robojack",
		SmallText:  "Powered by RoboJack",
	}
)

// Default returns the default presence for an event, this is what most events are
func Default(e *calendar.Event) (*Presence, error) {
	base := EventBases[e.ColorId]
	templateData := TemplateData{Event: e}

	if e.ColorId == MeetingColour {
		base.AddAttendees(e)
	} else if e.ColorId == TTRPGColour {
		templateData.TTRPGMeta = GetTTRPGMeta(e)
	}

	err := base.TemplateData(templateData)
	if err != nil {
		return nil, err
	}

	return &base, nil
}
