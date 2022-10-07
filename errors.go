package brandfetch

type warn string

// Error implements the error interface.
func (e warn) Error() string {
	return "brandfetch: " + string(e)
}

// List of supported errors.
const (
	// ErrHTTPClient is the error returning in case of missing or invalid data.
	ErrHTTPClient = warn("invalid http client")
	// ErrNoResults is returned when the request doesn't return a result.
	ErrNoResults = warn("no results for this query")
	// ErrRequest is returned when the request failed to be built or parsed.
	ErrRequest = warn("bad request body")
	// ErrResponse is the dedicated error to any response in error.
	ErrResponse = warn("unsupported response")
)
