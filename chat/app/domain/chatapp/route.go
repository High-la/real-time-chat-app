package chatapp

import (
	"net/http"

	"github.com/High-la/real-time-chat-app/chat/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App) {

	api := newApp()

	app.HandlerFunc(http.MethodGet, "", "/test", api.test)
}
