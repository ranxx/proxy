package model

// Client ...
type Client struct {
	Base

	Acccount string `json:"account"`
	ClientID int64  `json:"client_id"`
	Address  string `json:"address"`
}
