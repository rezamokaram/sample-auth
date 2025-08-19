package notification

import (
	"context"
	"fmt"
	"time"

	"github.com/rezamokaram/sample-auth/internal/common"
	"github.com/rezamokaram/sample-auth/internal/notification/domain"
	"github.com/rezamokaram/sample-auth/internal/notification/port"
	userDomain "github.com/rezamokaram/sample-auth/internal/user/domain"
	userPort "github.com/rezamokaram/sample-auth/internal/user/port"
	"github.com/rezamokaram/sample-auth/pkg/fp"
)

type service struct {
	repo       port.Repo
	outboxRepo common.OutboxRepo
	userPort   userPort.Service
}

func NewService(repo port.Repo, userPort userPort.Service, outboxRepo common.OutboxRepo) port.Service {
	return &service{
		repo:       repo,
		userPort:   userPort,
		outboxRepo: outboxRepo,
	}
}

func (s *service) Send(ctx context.Context, notif *domain.Notification) error {
	user, err := s.userPort.GetUserByFilter(ctx, &userDomain.UserFilter{
		ID: notif.UserID,
	})

	if err != nil {
		return err
	}

	notifID, err := s.repo.Create(ctx, notif)
	if err != nil {
		return err
	}

	return s.repo.CreateOutbox(ctx, &domain.NotificationOutbox{
		NotifID: notifID,
		Data: domain.OutboxData{
			Dest: func() string {
				switch notif.Type {
				case domain.NotifTypeSMS:
					return string(user.Phone)
				default:
					return ""
				}
			}(),
			Content: notif.Content,
			Type:    notif.Type,
		},
		Status: common.OutboxStatusCreated,
		Type:   common.OutboxTypeNotif,
	})
}

func (s *service) Handle(ctx context.Context, outboxes []domain.NotificationOutbox) error {
	outBoxIDs := fp.Map(outboxes, func(o domain.NotificationOutbox) common.OutboxID {
		return o.OutboxID
	})

	if err := s.outboxRepo.UpdateBulkStatuses(ctx, common.OutboxStatusPicked, outBoxIDs...); err != nil {
		return fmt.Errorf("failed to update notif outbox statuses to picked %w", err)
	}

	for _, outbox := range outboxes {
		fmt.Printf("dest : %s, content : %s\n", outbox.Data.Dest, outbox.Data.Content)
	}

	if err := s.outboxRepo.UpdateBulkStatuses(ctx, common.OutboxStatusDone, outBoxIDs...); err != nil {
		return fmt.Errorf("failed to update notif outbox statuses to done %w", err)
	}

	return nil
}

func (s *service) Interval() time.Duration {
	return time.Second * 10
}

func (s *service) Query(ctx context.Context) ([]domain.NotificationOutbox, error) {
	return s.repo.QueryOutboxes(ctx, 100, common.OutboxStatusCreated)
}

func (s *service) CheckUserNotifValue(ctx context.Context, userID userDomain.UserID, val string) (bool, error) {
	expected, err := s.repo.GetUserNotifValue(ctx, userID)
	if err != nil {
		return false, err
	}

	return expected == val, nil
}
