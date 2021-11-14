/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/24
   Description :
-------------------------------------------------
*/

package consts

// 足迹审核状态
const (
	FootprintStateNewWait               = 10 // 新增等待审核
	FootprintStateChangeWait            = 20 // 修改后等待审核
	FootprintStatePassAudit             = 30 // 通过审核
	FootprintStateNoPublishNoPassAudit  = 40 // 用户不发布，且该故事没有经过审核
	FootprintStateNoPublishPassAudit    = 50 // 用户不发布，但故事已经审核通过
	FootprintStateNoPassAudit           = 60 // 审核不通过
	FootprintStateThirdpartyNoPassAudit = 70 // 三方审核不通过
	FootprintStateDeleted               = 90 // 删除状态
)

// 足迹发布状态
const (
	FootprintStateNoPublish = 10 // 不发布
	FootprintStateIsPublish = 20 // 发布
)
