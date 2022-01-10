package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/ranxx/proxy/constant"
	msg "github.com/ranxx/proxy/proto/msg/v1"
	"github.com/ranxx/ztcp/conn"
	"github.com/ranxx/ztcp/conner"
)

// Client ...
type Client struct {
	logPrefix             string
	body                  *msg.TCPBody
	once                  *sync.Once
	localConn, remoteConn conner.Conner
}

// NewClient ...
func NewClient(logPrefix string, body *msg.TCPBody) *Client {
	return &Client{
		logPrefix: logPrefix,
		body:      body,
		once:      new(sync.Once),
	}
}

// Start 开始
func (c *Client) Start() {
	// 连接本地
	localConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.body.Laddr.Ip, c.body.Laddr.Port))
	if err != nil {
		log.Println(c.logPrefix, "连接local失败", err)
		return
	}
	c.localConn = conn.NewConn(0, localConn)

	// 连接远程
	remoteConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.body.Raddr.Ip, c.body.Raddr.Port))
	if err != nil {
		log.Println(c.logPrefix, "连接remote失败", err)
		return
	}
	c.remoteConn = conn.NewConn(1, remoteConn)

	log.Println(c.logPrefix, fmt.Sprintf("%s -> %s", remoteConn.RemoteAddr(), localConn.RemoteAddr()))

	bind := &msg.TransferBind{Id: c.body.Rid}
	if _, err := c.remoteConn.Writer().WriteValueWithID(constant.Transfer, bind); err != nil {
		log.Println(c.logPrefix, "写入bind消息失败", err)
		return
	}

	// 开启读写
	go func() {
		defer c.Close()
		if rn, err := io.Copy(localConn, remoteConn); err != nil {
			log.Println(c.logPrefix, rn, err)
		}
	}()

	go func() {
		defer c.Close()
		if rn, err := io.Copy(remoteConn, localConn); err != nil {
			log.Println(c.logPrefix, rn, err)
		}
	}()
}

// Close 关闭
func (c *Client) Close() {
	c.once.Do(func() {
		if c.localConn != nil {
			c.localConn.Close()
		}
		if c.remoteConn != nil {
			c.remoteConn.Close()
		}
	})
}

// // Receive 接受数据
// func (c *Client) Receive(msg *[]byte) {

// }
