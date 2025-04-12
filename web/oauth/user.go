package oauth

import (
	"context"
	"encoding/json"
	"io"

	"golang.org/x/oauth2"
)

// UserEndpoint is the endpoint to request when fetching Discord user data
const UserEndpoint = "https://discord.com/api/oauth2/@me"

// UserEndpointPayload wraps the JSON response from Discord when requesting UserEndpoint
type UserEndpointPayload struct {
	User User `json:"user"`
}

// User holds information about a Discord user
type User struct {
	AvatarHash  string `json:"avatar"`
	DisplayName string `json:"global_name"`
	ID          string `json:"id"`
	Username    string `json:"username"`
}

// GetUser takes an OAuth bearer token and fetches data for the user it belongs to
func GetUser(token *oauth2.Token) (*User, error) {
	res, err := Config.Client(context.Background(), token).Get(UserEndpoint)
	if err != nil {
		return nil, err
	}

	// Read response from Discord API
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal response body into struct
	var data UserEndpointPayload
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data.User, nil
}
