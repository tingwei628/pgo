package repository

import (
	"context"
	"github.com/tingwei628/pgo/webapi/internal/entity"
)

type Todo interface {
	InsertItem(ctx context.Context, item entity.Item) error
	GetAllItems(ctx context.Context) ([]entity.Item, error)
	GetItemsByKeyword(ctx context.Context, keyword string) ([]entity.Item, error)
}
