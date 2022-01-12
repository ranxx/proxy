package constant

// event
const (
	ExitEvent = "exit_event"
	// 创建 tcp tunnel 的 client事件
	CreateTCPTunnelClientEvent = "create_tcp_tunnel_client_event"
	// 新增 tcp tunnel 事件
	NewTCPTunnelEvent = "new_tcp_tunnel_event"
	// 创建 tcp tunnel 事件
	CreateTCPTunnelEvent = "create_tcp_tunnel_event"
	// 停止 tcp tunnel 事件
	StopTCPTunnelEvent = "stop_tcp_tunnel_event"
	// 新 client 事件
	NewClientEvent = "new_client_event"
	// 停止 client 事件
	StopClientEvent = "stop_client_event"
	// 批量新增 client 事件
	BatchNewClientEvent = "batch_new_client_event"
	// 删除 client 事件
	DelClientEvent = "del_client_event"
	// 新增 tcp 连接事件
	NewTCPConnectionEvent = "new_tcp_connection_event"
)
