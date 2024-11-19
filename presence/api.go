package presence

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	// CredentialsFile is needed by the Google API to authenticate the application
	CredentialsFile = "credentials.json"
	// TimeFormat is the format that event timestamps are read as, this is set by Google
	TimeFormat = time.RFC3339
)

var (
	// Client is the API client that interacts with the Google Calendar API
	Client *calendar.Service
	// Config is used for the OAuth flow to authenticate with Google
	Config *oauth2.Config
	// Token is the OAuth token from when the user authenticated
	Token *oauth2.Token
	// OAuthStateToken is a random string used to prevent CSRF attacks
	OAuthStateToken = utils.RandString(16)
)

func init() {
	b, err := os.ReadFile(CredentialsFile)
	if err != nil {
		log.Fatal("Unable to read credentials.json: %s", err)
	}

	Config, err = google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatal("Unable to authenticate with Google: %s", err)
	}

	Token, err = TokenFromFile("token.json")
	if err != nil {
		authURL := Config.AuthCodeURL(OAuthStateToken, oauth2.AccessTypeOffline)
		log.Info("To authenticate Google Calendar, visit the following URL: %s", authURL)
		return
	}

	SetupService()
}

// SetupService is called whenever an OAuth token is available to setup the API client
func SetupService() {
	client := Config.Client(context.Background(), Token)
	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Unable to create Calendar service: %s", err)
	}

	Client = srv
}

// TokenFromFile pulls the OAuth token from a file, if it exists
func TokenFromFile(filename string) (*oauth2.Token, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Between fetches all events between a start and end time
func Between(start time.Time, end time.Time) (*calendar.Events, error) {
	return Client.Events.List("primary").
		SingleEvents(true).
		TimeMin(start.Format(TimeFormat)).
		TimeMax(end.Format(TimeFormat)).
		OrderBy("startTime").
		EventTypes("default", "focusTime", "fromGmail", "outOfOffice", "workingLocation").
		Do()
}

// Today fetches all events that start or end today
func Today() ([]*calendar.Event, error) {
	start := utils.DayStart(time.Now())
	end := utils.DayEnd(time.Now())

	events, err := Between(start, end)
	if err != nil {
		return nil, err
	}

	return RemoveAllDayEvents(events), nil
}

// Now fetches all events that are currently ongoing
func Now() ([]*calendar.Event, error) {
	start := utils.HourStart(time.Now())
	end := utils.HourEnd(time.Now())

	events, err := Between(start, end)
	if err != nil {
		return nil, err
	}

	return RemoveAllDayEvents(events), nil
}

// RemoveAllDayEvents purges all-day events from a list of events
func RemoveAllDayEvents(events *calendar.Events) []*calendar.Event {
	var timedEvents []*calendar.Event
	for _, e := range events.Items {
		if e.Start.DateTime != "" {
			timedEvents = append(timedEvents, e)
		}
	}

	return timedEvents
}
