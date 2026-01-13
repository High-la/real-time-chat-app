// Package mid provides app level middleware support.
package mid

import (
	"github.com/High-la/real-time-chat-app/chat/foundation/web"
)

// tests if the encoder has an error inside of it.
func checkIsError(e web.Encoder) error {
	err, hasError := e.(error)
	if hasError {
		return err
	}

	return nil
}
