package port

import (
	"context"

	"github.com/rezamokaram/sample-auth/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error)
}
