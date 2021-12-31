package bootstrap

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/model"
	tunnelsV1 "github.com/ranxx/proxy/proto/tunnels/v1"
	usersV1 "github.com/ranxx/proxy/proto/users/v1"
	"github.com/ranxx/proxy/tunnels"
	tunnelsStore "github.com/ranxx/proxy/tunnels/store"
	"github.com/ranxx/proxy/users"
	"github.com/ranxx/proxy/users/store"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Start ...
func Start() {
	// 初始化 db
	db, err := gorm.Open(mysql.Open(config.Cfg.Mysql.Dsn))
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Tunnel{})

	// 初始化module
	_users := users.NewUsers(store.NewUsers(db))
	_tunnes := tunnels.NewTunnels(tunnelsStore.NewTunnels(db))

	// 初始化 grpc
	gserver := grpc.NewServer()
	usersV1.RegisterUsersServer(gserver, _users)
	tunnelsV1.RegisterTunnelsServer(gserver, _tunnes)

	// 初始化 http
	root := mux.NewRouter()
	//router := root.PathPrefix("/api").Subrouter()
	// _users.ProvideHTTP(router)
	// _tunnes.ProvideHTTP(router)
	root.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		m, _ := route.GetMethods()
		log.Printf("walk: %s %#v\n", tpl, m)
		return nil
	})

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
