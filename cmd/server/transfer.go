package server

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"strconv"

// 	"github.com/ranxx/grequests"
// 	"github.com/ranxx/proxy/api"
// 	"github.com/ranxx/proxy/config"
// 	"github.com/ranxx/proxy/pkg/utils"
// 	proto "github.com/ranxx/proxy/proto/msg"
// 	"github.com/spf13/cobra"
// )

// type addCommand struct {
// 	laddrs  *[]string
// 	raddr   *string
// 	machine *string
// }

// func (a *addCommand) Cmd() *cobra.Command {
// 	add := &cobra.Command{
// 		Use:   "add",
// 		Short: "添加穿透的ip:port 列表",
// 		Long: `内网穿透服务端，需要指定被穿透的机器。
// eg:
// 1) goproxy server nat add --laddr ip:port --raddr ip:port`,
// 		Run: a.Run,
// 	}

// 	a.laddrs = add.Flags().StringSliceP("laddr", "l", nil, "localhost addr")
// 	a.raddr = add.Flags().StringP("raddr", "r", "", "remote addr")
// 	a.machine = add.Flags().StringP("machine", "m", "", "client ip:port；可通过server list命令获取")
// 	add.MarkFlagRequired("laddr")
// 	add.MarkFlagRequired("raddr")
// 	return add
// }

// func (a *addCommand) Run(cmd *cobra.Command, args []string) {
// 	l := make([]proto.Addr, 0, len(*a.laddrs))
// 	for _, laddr := range *a.laddrs {
// 		laddrIP, laddrPort, err := utils.ParseAddrString(laddr)
// 		if err != nil {
// 			panic(err)
// 		}
// 		l = append(l, proto.Addr{Ip: laddrIP, Port: int32(laddrPort)})
// 	}
// 	raddrIP, raddrPort, err := utils.ParseAddrString(*a.raddr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	r := proto.Addr{Ip: raddrIP, Port: int32(raddrPort)}

// 	url := fmt.Sprintf("%s:%d/transfer/tcp", utils.RetRealHTTP(config.Cfg.Server.IP), config.Cfg.Server.APIPort)

// 	req := struct {
// 		Laddr   []proto.Addr `json:"laddr"`
// 		Raddr   proto.Addr   `json:"raddr"`
// 		Machine string       `json:"machine"`
// 	}{
// 		Laddr: l, Raddr: r, Machine: *a.machine,
// 	}

// 	resp := api.Message{}
// 	err = grequests.Post(context.TODO(), url, req, &resp)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp.Code, resp.Msg, resp.Data)
// }

// type removeCommand struct{}

// func (r *removeCommand) Cmd() *cobra.Command {
// 	remove := &cobra.Command{
// 		Use:   "rm",
// 		Short: "移除监听的端口",
// 		Long: `内网穿透服务端，按端口移除transfer
// eg:
// 1) goproxy server rm port [port...]`,
// 		Run: r.Run,
// 	}
// 	return remove
// }

// func (r *removeCommand) Run(cmd *cobra.Command, args []string) {
// 	// 端口
// 	// 处理args
// 	ports := make([]int, 0, len(args))
// 	for _, v := range args {
// 		port, err := strconv.Atoi(v)
// 		if err != nil {
// 			panic(err)
// 		}
// 		ports = append(ports, port)
// 	}

// 	url := fmt.Sprintf("%s:%d/transfer/port", utils.RetRealHTTP(config.Cfg.Server.IP), config.Cfg.Server.APIPort)

// 	req := struct {
// 		Ports []int `json:"ports"`
// 	}{
// 		Ports: ports,
// 	}

// 	resp := api.Message{}
// 	err := grequests.Delete(context.TODO(), url, req, &resp)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp.Code, resp.Msg, resp.Data)
// }

// type listCommand struct {
// 	transfer *bool
// }

// func (l *listCommand) Cmd() *cobra.Command {
// 	list := &cobra.Command{
// 		Use:   "list",
// 		Short: "列表transfer",
// 		Long: `内网穿透服务端，列表transfer
// eg:
// 1) goproxy server list`,
// 		Run: l.Run,
// 	}
// 	l.transfer = list.Flags().BoolP("transfer", "t", false, "查看tansfer列表")
// 	return list
// }

// func (l *listCommand) Run(cmd *cobra.Command, args []string) {
// 	if l.transfer == nil || *l.transfer == false {
// 		l.ListClient()
// 		return
// 	}
// 	l.ListTransfer()
// 	return
// }

// func (l *listCommand) ListTransfer() {
// 	url := fmt.Sprintf("%s:%d/transfer", utils.RetRealHTTP(config.Cfg.Server.IP), config.Cfg.Server.APIPort)

// 	resp := api.Message{}
// 	err := grequests.Get(context.TODO(), url, nil, &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if resp.Data == nil {
// 		fmt.Println(resp.Code, resp.Msg, resp.Data)
// 		return
// 	}

// 	data, err := json.MarshalIndent(resp.Data, "", "    ")
// 	fmt.Println(string(data))
// }

// func (l *listCommand) ListClient() {
// 	url := fmt.Sprintf("%s:%d/client", utils.RetRealHTTP(config.Cfg.Server.IP), config.Cfg.Server.APIPort)

// 	resp := api.Message{}
// 	err := grequests.Get(context.TODO(), url, nil, &resp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if resp.Data == nil {
// 		fmt.Println(resp.Code, resp.Msg, resp.Data)
// 		return
// 	}

// 	data, err := json.MarshalIndent(resp.Data, "", "    ")
// 	fmt.Println(string(data))
// }
