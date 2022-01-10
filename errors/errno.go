package errors

// ErrNumber ...
type ErrNumber int

// const 错误码
const (
	Err = (ErrNumber)(10000)

	// 已存在
	ErrExists = (ErrNumber)(10400)

	// 密码错误
	ErrPasswd = (ErrNumber)(10401)

	// 没有权限
	ErrPermissionDenied = (ErrNumber)(10403)

	// 不存在
	ErrNotFound = (ErrNumber)(10404)
)
