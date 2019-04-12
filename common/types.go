package common

import "time"

// IDHeader is a common struct for message id headers
type IDHeader struct {
	MessageID     string
	CorrelationID string
	CreatedTS     *time.Time
}
