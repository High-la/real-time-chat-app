package chatapp

import (
	"context"
	"net/http"

	"github.com/High-la/real-time-chat-app/chat/foundation/web"
)

type app struct {
}

func newApp() *app {
	return &app{}
}

func (a *app) test(ctx context.Context, r *http.Request) web.Encoder {
	return status{
		Status: "ok",
	}
}
