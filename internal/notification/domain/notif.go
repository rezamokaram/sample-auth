package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/rezamokaram/sample-auth/internal/common"
	userDomain "github.com/rezamokaram/sample-auth/internal/user/domain"
)

type (
	NotifID     uint
	NotifType   uint8
	NotifStatus uint8
)

const (
	NotifTypeSMS NotifType = iota + 1
)

const (
	NotifStatusCreated NotifStatus = iota + 1
	NotifStatusSent
)

type Notification struct {
	ID            NotifID
	CreatedAt     time.Time
	UserID        userDomain.UserID
	Type          NotifType
	Content       string
	ForValidation bool
	TTL           time.Duration
}

func (n *Notification) Normalize() {
	n.Content = strings.TrimSpace(n.Content)
	n.ID = 0
}

func (n *Notification) Validate() error {
	if n.UserID == 0 {
		return errors.New("empty user id")
	}

	return nil
}

func NewNotification(userID userDomain.UserID, content string, notifType NotifType, forValidation bool, ttl time.Duration) *Notification {
	return &Notification{
		UserID:        userID,
		Type:          notifType,
		Content:       content,
		CreatedAt:     time.Now(),
		ForValidation: forValidation,
		TTL:           ttl,
	}
}

type OutboxData struct {
	Dest    string
	Content string
	Type    NotifType
}

type NotificationOutbox struct {
	OutboxID common.OutboxID
	NotifID  NotifID
	Data     OutboxData
	Status   common.OutboxStatus
	Type     common.OutboxType
}
