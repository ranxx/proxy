package event

import (
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/constant"
	"github.com/ranxx/proxy/internal/model"
	msg "github.com/ranxx/proxy/proto/msg/v1"
)

// SubscribeNewClientEvent 订阅新增 client 事件
func SubscribeNewClientEvent(fc func(*model.Client)) func() {
	return config.Obs.SubscribeByTopicFunc(constant.NewClientEvent, fc)
}

// PublishNewClientEvent 发布新增 client 事件
func PublishNewClientEvent(body *model.Client) {
	config.Obs.Publish(constant.NewClientEvent, body)
}

// SubscribeBatchNewClientEvent 订阅批量新增 client 事件
func SubscribeBatchNewClientEvent(fc func([]*model.Client)) func() {
	return config.Obs.SubscribeByTopicFunc(constant.BatchNewClientEvent, fc)
}

// PublishBatchNewClientEvent 发布批量新增 client 事件
func PublishBatchNewClientEvent(body []*model.Client) {
	config.Obs.Publish(constant.BatchNewClientEvent, body)
}

// SubscribeDelClientEvent 订阅删除 client 事件
func SubscribeDelClientEvent(fc func(*model.Client)) func() {
	return config.Obs.SubscribeByTopicFunc(constant.DelClientEvent, fc)
}

// PublishDelClientEvent 发布删除 client 事件
func PublishDelClientEvent(body *model.Client) {
	config.Obs.Publish(constant.DelClientEvent, body)
}

// SubscribeCreateTCPTunnelClientEvent 订阅创建 tcp tunnel client 事件
func SubscribeCreateTCPTunnelClientEvent(fc func(body *msg.TCPBody)) {
	config.Obs.SubscribeByTopicFunc(constant.CreateTCPTunnelClientEvent, fc)
}

// PublishCreateTCPTunnelClientEvent 发布创建 tcp tunnel client 事件
func PublishCreateTCPTunnelClientEvent(body *msg.TCPBody) {
	config.Obs.Publish(constant.CreateTCPTunnelClientEvent, body)
}
