package pixaven

import (
    "errors"
    "fmt"
)

// PixavenError is an error returned by the Pixaven API.
type PixavenError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// Error implements the error interface.
func (c *PixavenError) Error() string {
    return fmt.Sprintf("pixaven: [%d] %s", c.Code, c.Message)
}

// In-code validation and other frequent error cases.
var (
    ErrInvalidSourceType = errors.New("pixaven: Invalid request source type")
    ErrBinaryWebhook     = errors.New("pixaven: Webhooks are not supported when using binary responses")
    ErrBinaryStorage     = errors.New("pixaven: External storage is not supported when using binary responses")
    ErrNoSuccess         = errors.New("pixaven: Success is missing in the response")
)
