package oauth

import (
	"sync"

	"github.com/Jack-Gledhill/robojack/utils"
)

// StateLength is the integer number of characters that every state string will container
const StateLength = 32

var (
	mutex  = &sync.Mutex{}
	states []string
)

// NewState generates a new random state and adds it to the list of valid states
// This uses a mutex to prevent two goroutines trying to add a new state at the same time
func NewState() string {
	mutex.Lock()
	defer mutex.Unlock()

	state := utils.RandString(StateLength)
	states = append(states, state)

	return state
}

// PopState checks to see if a given state is valid. If it is, it's removed from the list of valid states
// This uses a mutex lock to prevent concurrent goroutines from trying to pop a state at the same time
func PopState(state string) bool {
	found := false
	var newStates []string

	mutex.Lock()
	defer mutex.Unlock()

	for _, s := range states {
		newStates = append(newStates, s)
		if s == state {
			found = true
		}
	}

	return found
}
