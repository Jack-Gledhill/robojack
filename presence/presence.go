package presence

import (
	"time"

	"github.com/Jack-Gledhill/robojack/utils"

	"google.golang.org/api/calendar/v3"
)

// TemplateData is passed to text/template when rendering the presence.
// It contains both the standard event data and TTRPG metadata.
type TemplateData struct {
	Event     *calendar.Event
	TTRPGMeta TTRPGMeta
}

// Presence is the data structure used to update the Discord presence
// This is the exact same structure as expected by rich-go on the client side
type Presence struct {
	State      string     `json:"state"`
	Details    string     `json:"details"`
	LargeImage string     `json:"largeimage"`
	LargeText  string     `json:"largetext"`
	SmallImage string     `json:"smallimage"`
	SmallText  string     `json:"smalltext"`
	Timestamps Timestamps `json:"timestamps"`
	Party      Party      `json:"party"`
	Buttons    []Button   `json:"buttons"`
}

// Timestamps holds information about the start and end times of an event
type Timestamps struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// Party holds information about player counts in a game
type Party struct {
	ID         string `json:"id"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
}

// Button is a clickable button that can be added to the presence to do things, such as go to a URL
type Button struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// AddAttendees adds the number of attendees to the presence, accounting for any edge cases
func (p *Presence) AddAttendees(e *calendar.Event) {
	// If there's no attendee list (or only one attendee), assume it's a 1:1 (i.e. 2 attendees)
	if len(e.Attendees) <= 1 {
		p.Party.Players = 2
		p.Party.MaxPlayers = 2
		return
	}

	currentAttendees := 0
	organiserIsAttendee := false

	// Count attendees, but exclude those who've declined
	for _, a := range e.Attendees {
		// Ignore those who've already declined
		if a.ResponseStatus != "declined" {
			currentAttendees++
		}

		// If the organiser is on the attendee list, we don't need to count them separately
		if a.Organizer {
			organiserIsAttendee = true
		}
	}

	// Organiser (me) isn't on attendee list, so we need to count them
	if !organiserIsAttendee {
		currentAttendees++
	}

	p.Party.Players = currentAttendees
	p.Party.MaxPlayers = len(e.Attendees)
}

// AddTimestamps fills out the event's start & end times into the presence response
func (p *Presence) AddTimestamps(e *calendar.Event) error {
	start, err := time.Parse(TimeFormat, e.Start.DateTime)
	if err != nil {
		return err
	}

	end, err := time.Parse(TimeFormat, e.End.DateTime)
	if err != nil {
		return err
	}

	p.Timestamps = Timestamps{
		Start: start,
		End:   end,
	}
	return nil
}

// TemplateData fills out the presence's details and state with the event's data, according to the base used
func (p *Presence) TemplateData(data TemplateData) error {
	var err error
	p.Details, err = utils.TemplateString(p.Details, data)
	if err != nil {
		return err
	}

	p.State, err = utils.TemplateString(p.State, data)
	if err != nil {
		return err
	}

	return nil
}
