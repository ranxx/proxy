package utils

import "gorm.io/gorm"

// Pager 页面
func Pager(in *gorm.DB, offset, limit int64) *gorm.DB {
	ret := in
	if limit >= 0 {
		ret = ret.Limit(int(limit))
	}
	if offset > 0 && limit >= 0 {
		ret = ret.Offset(int((offset - 1) * limit))
	}
	return ret
}
