package port

import (
	"context"
)

type LoginService interface {
	UserLogin(ctx context.Context, email string, password string) (string, error)
}
