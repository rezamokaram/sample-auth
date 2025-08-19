package port

import (
	"context"

	"github.com/rezamokaram/sample-auth/internal/common"
	"github.com/rezamokaram/sample-auth/internal/notification/domain"
	userDomain "github.com/rezamokaram/sample-auth/internal/user/domain"
)

type Service interface {
	Send(ctx context.Context, notif *domain.Notification) error
	CheckUserNotifValue(ctx context.Context, userID userDomain.UserID, val string) (bool, error)
	common.OutboxHandler[domain.NotificationOutbox]
}
