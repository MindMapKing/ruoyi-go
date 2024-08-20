//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-30 21:59:34 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"main/internal/iot_dev/service"
)

// List 查询页
func (w IotPrdActionController) List(c *gin.Context) {
	util.BuildTpl(c, "iot_dev/action/list.html").WriteTpl()
}

// Add 新增页
func (w IotPrdActionController) Add(c *gin.Context) {
	util.BuildTpl(c, "iot_dev/action/add.html").WriteTpl()
}

// Edit 修改页
func (w IotPrdActionController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	actionService := service.IotPrdActionService{}
	entity, err := actionService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.BuildTpl(c, "iot_dev/action/edit.html").WriteTpl(gin.H{
		"action": entity,
	})
}
