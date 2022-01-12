package bootstrap

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/ranxx/proxy/apiserver"
	"github.com/ranxx/proxy/config"
	"github.com/ranxx/proxy/internal/model"
)

// Start ...
func Start() {
	db := config.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Tunnel{})
	db.AutoMigrate(&model.Client{})

	root := apiserver.Router()

	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		defer wait.Done()
		fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", config.Cfg.HTTP.Port), root))
	}()

	wait.Wait()
}
