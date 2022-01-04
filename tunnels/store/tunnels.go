package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/model"
	"github.com/ranxx/proxy/pkg/utils"
)

// Tunnels users
type Tunnels interface {
	// 列表
	List(ctx context.Context, account string, offset, limit int64) ([]*model.Tunnel, int64, error)

	// 创建
	Create(ctx context.Context, items ...*model.Tunnel) error
}

// NewTunnels ...
func NewTunnels() Tunnels {
	return &localTunnels{db: config.GetDB()}
}

type localTunnels struct {
	db *gorm.DB
}

func (l *localTunnels) List(ctx context.Context, account string, offset, limit int64) ([]*model.Tunnel, int64, error) {
	items := make([]*model.Tunnel, 0, 10)
	total := int64(0)

	db := utils.Pager(l.db.Model(&model.Tunnel{}), offset, limit).Where("account = ?", account)

	if err := db.Count(&total).Error; err != nil || total <= 0 {
		return nil, 0, err
	}

	if err := db.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (l *localTunnels) Create(ctx context.Context, items ...*model.Tunnel) error {
	// 加锁
	l.db.CreateBatchSize = 1024
	return l.db.Model(&model.Tunnel{}).Create(items).Error
}
