package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/model"
	"github.com/ranxx/proxy/pkg/utils"
	v1 "github.com/ranxx/proxy/proto/clients/v1"
)

// Clients users
type Clients interface {
	// 列表
	List(ctx context.Context, userID int64, machines []string, offset, limit int64) ([]*model.Client, int64, error)

	// 获取
	Get(ctx context.Context, id int64) (*model.Client, bool, error)

	// 创建
	Create(ctx context.Context, items ...*model.Client) error

	// 更新
	Update(ctx context.Context, item *model.Client) error

	// 删除
	Delete(ctx context.Context, id int64) error

	// 暂停所有
	StopAll(ctx context.Context) error
}

// NewClients ...
func NewClients() Clients {
	return &localClients{db: config.GetDB()}
}

type localClients struct {
	db *gorm.DB
}

func (l *localClients) List(ctx context.Context, userID int64, machines []string, offset, limit int64) ([]*model.Client, int64, error) {
	var items []*model.Client
	var total int64

	db := utils.Pager(l.db.Model(&model.Client{}), offset, limit)

	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}

	if len(machines) > 0 {
		db = db.Where("machine in (?)", machines)
	}

	if err := db.Count(&total).Error; err != nil || total <= 0 {
		return nil, 0, err
	}

	if err := db.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (l *localClients) Get(ctx context.Context, id int64) (*model.Client, bool, error) {
	item := new(model.Client)

	ret := l.db.Model(&model.Client{}).Where("id = ?", id).First(item)
	if ret.RowsAffected == 0 {
		return nil, false, nil
	}

	if ret.Error != nil {
		return nil, false, ret.Error
	}

	return item, true, nil
}

func (l *localClients) Create(ctx context.Context, items ...*model.Client) error {
	return l.db.Model(&model.Client{}).Create(items).Error
}

func (l *localClients) Update(ctx context.Context, item *model.Client) error {
	return l.db.Model(&model.Client{}).Where("id = ?", item.ID).Updates(map[string]interface{}{
		"user_id": item.UserID,
		"status":  item.Status,
	}).Error
}

func (l *localClients) Delete(ctx context.Context, id int64) error {
	return l.db.Model(&model.Client{}).Where("id = ?", id).Delete(nil).Error
}

func (l *localClients) StopAll(ctx context.Context) error {
	return l.db.Model(&model.Client{}).Where("1=1").Updates(map[string]interface{}{"status": model.ClientStatus(v1.ClientStatus_Stop)}).Error
}
