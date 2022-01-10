package event

import (
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/internal/model"
	msg "github.com/ranxx/proxy/proto/msg/v1"
)

// SubscribeNewTCPTunnelEvent 订阅新增 tcp tunnel 事件
func SubscribeNewTCPTunnelEvent(fc func(...*model.Tunnel)) {
	config.Obs.SubscribeByTopicFunc(constant.NewTCPTunnelEvent, fc)
}

// PublishNewTCPTunnelEvent 发布新增 tcp tunnel 事件
func PublishNewTCPTunnelEvent(body ...*model.Tunnel) {
	i := make([]interface{}, 0, len(body))
	for _, v := range body {
		i = append(i, v)
	}
	config.Obs.Publish(constant.NewTCPTunnelEvent, i...)
}

// SubscribeNewTCPConnectionEvent 订阅新增 tcp 连接事件
func SubscribeNewTCPConnectionEvent(fc func(int64, *msg.TCPBody)) {
	config.Obs.SubscribeByTopicFunc(constant.NewTCPConnectionEvent, fc)
}

// PublishNewTCPConnectionEvent 发布新增 tcp 连接事件
func PublishNewTCPConnectionEvent(clientID int64, body *msg.TCPBody) {
	config.Obs.Publish(constant.NewTCPConnectionEvent, clientID, body)
}

// SubscribeCreateTCPTunnelEvent 订阅创建 tcp tunnel 事件
func SubscribeCreateTCPTunnelEvent(fc func(...*model.Tunnel) error) {
	config.Obs.SubscribeByTopicFunc(constant.CreateTCPTunnelEvent, fc)
}

// PublishCreateTCPTunnelEvent 发布创建 tcp tunnel 事件
func PublishCreateTCPTunnelEvent(body ...*model.Tunnel) {
	i := make([]interface{}, 0, len(body))
	for _, v := range body {
		i = append(i, v)
	}
	config.Obs.Publish(constant.CreateTCPTunnelEvent, i...)
}

// SubscribeStopTCPTunnelEvent 订阅停止 tcp tunnel 事件
func SubscribeStopTCPTunnelEvent(fc func(id int64)) {
	config.Obs.SubscribeByTopicFunc(constant.StopTCPTunnelEvent, fc)
}

// PublishStopTCPTunnelEvent 发布停止 tcp tunnel 事件
func PublishStopTCPTunnelEvent(id int64) {
	config.Obs.Publish(constant.StopTCPTunnelEvent, id)
}
