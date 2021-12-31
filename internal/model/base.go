package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	proto "github.com/ranxx/proxy/proto/msg/v1"
)

// Base ...
type Base struct {
	ID        int64     `sql:"type:bigint(20);primary_key;default:autoIncrement;comment:唯一id" json:"id"`
	UpdatedAt time.Time `sql:"type:timestamp;not null;default:current_timestamp on update current_timestamp" json:"updated_at"`
	CreatedAt time.Time `sql:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
}

func jsonScan(value, dst interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(b, dst)
}

func jsonValue(src interface{}) ([]byte, error) {
	b, e := json.Marshal(src)
	if e != nil {
		return nil, e
	}
	return b, nil
}

// Addr addr
type Addr proto.Addr

// Value 存储
func (a *Addr) Value() (driver.Value, error) {
	return jsonValue(a)
}

// Scan 读取
func (a *Addr) Scan(value interface{}) error {
	return jsonScan(value, a)
}
