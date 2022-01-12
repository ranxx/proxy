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
	List(ctx context.Context, userID int64, status []model.TunnelStatus, offset, limit int64) ([]*model.Tunnel, int64, error)

	// 创建
	Create(ctx context.Context, items ...*model.Tunnel) error

	// 获取
	Get(ctx context.Context, id int64) (*model.Tunnel, bool, error)

	// 删除
	Delete(ctx context.Context, id int64) error

	// 修改状态
	Status(ctx context.Context, id int64, status model.TunnelStatus) error
}

// NewTunnels ...
func NewTunnels() Tunnels {
	return &localTunnels{db: config.GetDB()}
}

type localTunnels struct {
	db *gorm.DB
}

func (l *localTunnels) List(ctx context.Context, userID int64, status []model.TunnelStatus, offset, limit int64) ([]*model.Tunnel, int64, error) {
	items := make([]*model.Tunnel, 0, 10)
	total := int64(0)

	db := utils.Pager(l.db.Model(&model.Tunnel{}), offset, limit)

	if userID >= 0 {
		db = db.Where("user_id = ?", userID)
	}

	if len(status) > 0 {
		db = db.Where("status in (?)", status)
	}

	if err := db.Count(&total).Error; err != nil || total <= 0 {
		return nil, 0, err
	}

	if err := db.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (l *localTunnels) Create(ctx context.Context, items ...*model.Tunnel) error {
	return l.db.Model(&model.Tunnel{}).Create(items).Error
}

// 获取
func (l *localTunnels) Get(ctx context.Context, id int64) (*model.Tunnel, bool, error) {
	item := new(model.Tunnel)

	ret := l.db.Model(&model.Tunnel{}).Where("id = ?", id).First(item)
	if ret.RowsAffected == 0 {
		return nil, false, nil
	}

	if ret.Error != nil {
		return nil, false, ret.Error
	}

	return item, true, nil
}

func (l *localTunnels) Delete(ctx context.Context, id int64) error {
	return l.db.Model(&model.Tunnel{}).Where("id = ?", id).Delete(nil).Error
}

func (l *localTunnels) Status(ctx context.Context, id int64, status model.TunnelStatus) error {
	return l.db.Model(&model.Tunnel{}).Where("id = ?", id).Update("status", status).Error
}
