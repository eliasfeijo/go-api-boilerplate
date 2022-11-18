package v1

// Error is a generic error response
type Error struct {
	Error interface{} `json:"error"`
}
