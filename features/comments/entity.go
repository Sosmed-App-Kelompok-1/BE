package comments

import (
	"context"
	"sosmed/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Comment struct {
	Id   uint
	Text string

	PostId uint

	User users.User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	Create() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	Create(ctx context.Context, data Comment) error
	Delete(ctx context.Context, commentId uint) error
	DeleteByPostId(ctx context.Context, postId uint) error
	DeleteByUserId(ctx context.Context, userId uint) error
}

type Repository interface {
	Create(ctx context.Context, data Comment) error
	Delete(ctx context.Context, commentId uint) error
	DeleteByPostId(ctx context.Context, postId uint) error
	DeleteByUserId(ctx context.Context, userId uint) error
}
