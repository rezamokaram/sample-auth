package http

import (
	"context"

	"github.com/rezamokaram/sample-auth/api/service"
	"github.com/rezamokaram/sample-auth/app"
	"github.com/rezamokaram/sample-auth/config"
)

// user service transient instance handler
func userServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.UserService] {
	return func(ctx context.Context) *service.UserService {
		return service.NewUserService(appContainer.UserService(ctx),
			cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute, appContainer.NotificationService(ctx))
	}
}
