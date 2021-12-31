package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/ranxx/proxy/internal/model"
)

// Clients users
type Clients interface {
	// 获取 user
	Get(ctx context.Context, email string) (*model.User, bool, error)

	// 创建
	Create(ctx context.Context, item *model.User) error
}

// NewClients ...
func NewClients(db *gorm.DB) Clients {
	return &localClients{db: db}
}

type localClients struct {
	db *gorm.DB
}

func (l *localClients) Get(ctx context.Context, email string) (*model.User, bool, error) {
	var item model.User

	ret := l.db.Model(&model.User{}).Where("email = ?", email).First(&item)
	if ret.RowsAffected <= 0 {
		return nil, false, nil
	}
	if ret.Error != nil {
		return nil, false, ret.Error
	}

	return &item, true, nil
}

func (l *localClients) Create(ctx context.Context, item *model.User) error {

	return l.db.Model(&model.User{}).Create(item).Error
}
