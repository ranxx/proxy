package constant

import (
	"github.com/ranxx/ztcp/pkg/message"
)

// message id
const (
	Normal     message.MsgID = iota
	HeartBeat                // 心跳
	ClientBind               // 客户端绑定
)

// message id for transfer [100-200)
const (
	Transfer message.MsgID = 100 + iota
	TCP
	HTTP
)
