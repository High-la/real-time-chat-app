package mux

import (
	"context"
	"net/http"

	"github.com/High-la/real-time-chat-app/chat/app/domain/chatapp"
	"github.com/High-la/real-time-chat-app/chat/app/sdk/mid"
	"github.com/High-la/real-time-chat-app/chat/foundation/logger"
	"github.com/High-la/real-time-chat-app/chat/foundation/web"
)

type Config struct {
	Log *logger.Logger
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config) http.Handler {

	logger := func(ctx context.Context, msg string, args ...any) {
		cfg.Log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		logger,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	chatapp.Routes(app)

	return app
}
