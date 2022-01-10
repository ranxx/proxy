package model

// Client ...
type Client struct {
	Base

	UserID int64 `json:"user_id"`
	// Account  string `json:"account"`
	ClientID int64  `json:"client_id"`
	Address  string `json:"address"`
}
