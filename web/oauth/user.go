package oauth

import (
	"context"
	"encoding/json"
	"fmt"
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

func (u *User) AvatarURL(size uint16) string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png?size=%d", u.ID, u.AvatarHash, size)
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
