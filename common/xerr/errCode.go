package xerr

const OK uint32 = 200

// 全局错误码

const SERVER_COMMON_ERROR uint32 = 100001
const REQUES_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005

// 业务模块

const (
	UserNotExists = 200404 // 用户不存在
)