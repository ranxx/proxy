package bootstrap

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/ranxx/proxy/apiserver"
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/model"
	tunnelsV1 "github.com/ranxx/proxy/proto/tunnels/v1"
	usersV1 "github.com/ranxx/proxy/proto/users/v1"
	"github.com/ranxx/proxy/tunnels"
	tunnelsStore "github.com/ranxx/proxy/tunnels/store"
	"github.com/ranxx/proxy/users"
	"github.com/ranxx/proxy/users/store"
	"google.golang.org/grpc"
)

// Start ...
func Start() {
	db := config.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Tunnel{})

	// 初始化module
	_users := users.NewUsers(store.NewUsers())
	_tunnes := tunnels.NewTunnels(tunnelsStore.NewTunnels())

	// 初始化 grpc
	gserver := grpc.NewServer()
	usersV1.RegisterUsersServer(gserver, _users)
	tunnelsV1.RegisterTunnelsServer(gserver, _tunnes)

	root := apiserver.Router()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.GRPC.Port))
	if err != nil {
		panic(err)
	}

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println(gserver.Serve(listener))
	}()

	go func() {
		defer wait.Done()
		fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", config.Cfg.HTTP.Port), root))
	}()

	wait.Wait()
}
