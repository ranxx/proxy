package model

import (
	"database/sql/driver"

	proto "github.com/ranxx/proxy/proto/msg/v1"
	v1 "github.com/ranxx/proxy/proto/tunnels/v1"
)

// TunnelStatus 状态
type TunnelStatus v1.Status

// // tunnel 状态
// const (
// 	TunnelStatusStop = (TunnelStatus)(0) + iota
// 	TunnelStatusRun
// 	TunnelStatusDel
// )

// Tunnel ...
type Tunnel struct {
	Base

	UserID  int64             `gorm:"colume:user_id;type:bigint(20);comment:用户id" json:"user_id"`
	Laddr   *Addr             `gorm:"colume:laddr;type:text;comment:本地地址" json:"laddr"`
	Raddr   *Addr             `gorm:"colume:raddr;type:text;comment:远端地址" json:"raddr"`
	Network proto.NetworkType `gorm:"colume:network;type:int;comment:网络类型" json:"network"`
	Match   *TunnelMatch      `gorm:"colume:match;type:text;comment:匹配策略" json:"match"`
	Status  TunnelStatus      `gorm:"colume:status;size:8;comment:状态" json:"status"`
}

// TableName ...
func (*Tunnel) TableName() string {
	return "tunnel"
}

// TunnelMatch 匹配
type TunnelMatch v1.Match

// Value 存储
func (t *TunnelMatch) Value() (driver.Value, error) {
	return jsonValue(t)
}

// Scan 读取
func (t *TunnelMatch) Scan(value interface{}) error {
	return jsonScan(value, t)
}

// // TunnelConfig config
// type TunnelConfig struct {
// 	Laddr   proto.Addr        `json:"laddr" yaml:"laddr"`
// 	Raddr   proto.Addr        `json:"raddr" yaml:"raddr"`
// 	Match   TunnelMatch       `json:"match" yaml:"match"`
// 	Network proto.NetworkType `json:"network" yaml:"network"`

// 	Account string            `json:"account" yaml:"account"`
// }

// // TransferInfo info
// type TransferInfo struct {
// 	Index int64 `json:"index"`
// 	TunnelConfig
// }
