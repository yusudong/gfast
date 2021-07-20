// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/model/internal"
)

// SysUserOnline is the golang structure for table sys_user_online.
type SysUserOnline internal.SysUserOnline

// Fill with you ideas below.

// SysUserOnlineSearchReq 列表搜索参数
type SysUserOnlineSearchReq struct {
	Username string `p:"userName"`
	Ip       string `p:"ipaddr"`
	comModel.PageReq
}