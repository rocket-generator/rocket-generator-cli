package objects

// RequestGroup ...
type RequestGroup struct {
	PathPrefix string
	Requests   []*Request
}
