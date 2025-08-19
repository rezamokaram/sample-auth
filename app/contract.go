package app

import (
	"context"

	"github.com/rezamokaram/sample-auth/config"
	notifPort "github.com/rezamokaram/sample-auth/internal/notification/port"
	userPort "github.com/rezamokaram/sample-auth/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	UserService(ctx context.Context) userPort.Service
	NotificationService(ctx context.Context) notifPort.Service
	DB() *gorm.DB
	Config() config.Config
}
