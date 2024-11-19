package web

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Jack-Gledhill/robojack/log"
	"github.com/Jack-Gledhill/robojack/presence"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func init() {
	// Register the routes under a group
	p := Server.Group("/presence")
	p.GET("/", Presence)
	p.GET("/auth", PresenceAuth)
}

// Presence calculates and returns the current presence based on the ongoing events
func Presence(c *gin.Context) {
	// Send an empty response if the service isn't authenticated yet
	if presence.Client == nil {
		c.Status(401)
		return
	}

	// Fetch ongoing events
	events, err := presence.Now()
	if err != nil {
		log.Error("Couldn't get current events: %s", err)
		c.Status(500)
		return
	}

	// If no events are ongoing, send an empty response
	if len(events) == 0 {
		c.Status(204)
	}

	// Return first event we got
	// TODO: possibly prioritise events based on start time, event type, etc.
	e := events[0]
	p, err := presence.GetPresence(e)
	if err != nil {
		log.Error("Couldn't get presence for event '%s': %s", e.Summary, err)
	}

	c.JSON(200, gin.H{
		"presence": p,
	})
}

// PresenceAuth is a callback for Google OAuth and will exchange an auth code for a token that can be used to access the API
func PresenceAuth(c *gin.Context) {
	// Validate the state token to prevent CSRF attacks
	state := c.Query("state")
	if state != presence.OAuthStateToken {
		log.Error("Invalid state token")
		c.Status(400)
		return
	}

	// Exchange the code for an OAuth token
	code := c.Query("code")
	token, err := presence.Config.Exchange(context.TODO(), code)
	if err != nil {
		log.Error("Couldn't exchange token: %s", err)
		c.Status(500)
		return
	}

	// Save the token to a file so we don't have to authenticate on each restart
	err = SaveToken("token.json", token)
	if err != nil {
		log.Error("Couldn't save token: %s", err)
	}

	// Set up the calendar service with the new token
	presence.Token = token
	presence.SetupService()
	c.JSON(200, gin.H{
		"message": "Authentication successful",
	})
}

// SaveToken saves the Google OAuth token to a local file so it can be used in the future
func SaveToken(filename string, token *oauth2.Token) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(token)
}
