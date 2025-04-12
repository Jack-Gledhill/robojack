package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R is a payload for a typical response
type R struct {
	DataPayload any    `json:"d,omitempty"`
	StatusCode  int    `json:"s"`
	Message     string `json:"m,omitempty"`
}

// Status will set the status code of this response
// You don't need to call this function if your code is http.StatusOK, that will be set by default
func (r *R) Status(s int) *R {
	r.StatusCode = s
	return r
}

// Data adds any type of data to the response, and should mostly be used for returning API data
func (r *R) Data(d any) *R {
	r.DataPayload = d
	return r
}

// Msg adds a string message to the response, this should only be used for errors or debugging purposes
func (r *R) Msg(s string) *R {
	r.Message = s
	return r
}

// Send will send the completed payload to the client. You should have already called Status, Msg and Data by this point
func (r *R) Send(c *gin.Context) {
	c.JSON(r.StatusCode, r)
}

// New creates a new R instance
func New() *R {
	return &R{
		StatusCode: http.StatusOK,
	}
}
