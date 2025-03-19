package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("用户不存在")
	ErrUserExists        = errors.New("用户已存在")
	ErrInvalidPassword   = errors.New("密码错误")
	ErrInvalidToken      = errors.New("无效的token")
	ErrEventNotFound     = errors.New("事件不存在")
	ErrEventStatus       = errors.New("事件状态错误")
	ErrStaffNotFound     = errors.New("安保人员不存在")
	ErrStaffExists       = errors.New("已经是安保人员")
	ErrStaffOffline      = errors.New("安保人员离线")
	ErrStaffBusy         = errors.New("安保人员忙碌中")
	ErrInvalidLocation   = errors.New("无效的位置信息")
	ErrInvalidEventType  = errors.New("无效的事件类型")
	ErrInvalidAction     = errors.New("无效的操作")
	ErrInvalidParameter  = errors.New("无效的参数")
	ErrDatabaseOperation = errors.New("数据库操作失败")
	ErrConfigNotFound    = errors.New("配置不存在")
	ErrConfigExists      = errors.New("配置已存在")
	ErrInvalidConfig     = errors.New("无效的配置")
)
