package model

import v1 "github.com/ranxx/proxy/proto/clients/v1"

// ClientStatus client 状态
type ClientStatus v1.ClientStatus

// Client ...
type Client struct {
	Base

	UserID  int64        `gorm:"colume:user_id;type:bigint(20);comment:用户id" json:"user_id"`
	Machine string       `gorm:"colume:machine;size:64;comment:机器" json:"machine"`
	Status  ClientStatus `gorm:"colume:status;size:8;default:1;comment:状态" json:"status"`
}

// TableName ...
func (*Client) TableName() string {
	return "client"
}
