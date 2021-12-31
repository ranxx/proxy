package model

import (
	"database/sql/driver"

	proto "github.com/ranxx/proxy/proto/msg/v1"
	v1 "github.com/ranxx/proxy/proto/tunnels/v1"
)

// TunnelStatus 状态
type TunnelStatus int8

// tunnel 状态
const (
	TunnelStatusStop = (TunnelStatus)(0) + iota
	TunnelStatusRun
	TunnelStatusDel
)

// Tunnel ...
type Tunnel struct {
	Base

	Account string            `gorm:"colume:account;size:64;comment:用户" json:"account"`
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

// // TunnelMatch 匹配
// type TunnelMatch struct {
// 	MachinePrefix string   `json:"machine_prefix" yaml:"machine_prefix"`
// 	Machines      []string `json:"machines" yaml:"machines"`
// }

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

// TransferMatch matching and 关系
type TransferMatch struct {
	Acccount      string   `json:"account" yaml:"account"`
	Machines      []string `json:"machines" yaml:"machines"`
	MachinePrefix string   `json:"machine_prefix" yaml:"machine_prefix"`
}

// TransferConfig config
type TransferConfig struct {
	Laddr   proto.Addr        `json:"laddr" yaml:"laddr"`
	Raddr   proto.Addr        `json:"raddr" yaml:"raddr"`
	Match   TransferMatch     `json:"match" yaml:"match"`
	Network proto.NetworkType `json:"network" yaml:"network"`
}

// TransferInfo info
type TransferInfo struct {
	Index int64 `json:"index"`
	TransferConfig
}
