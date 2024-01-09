package api

// CommandHeader is a header for a command.
// Contains general fields shared by many commands.
type CommandHeader struct {
	// AuthToken is an authentication token which grants access to a specific owner.
	AuthToken string
}
