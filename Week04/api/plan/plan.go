package plan

import (
	"Week04/models"
	"Week04/pkg/app"
	"Week04/pkg/utils"

	"github.com/gin-gonic/gin"
)

/**
 * File :   plan.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/12/16 19:47
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

// @Summary 获取套餐数据
// @Description 获取JSON
// @Tags 套餐
// @Param id path int true "id"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/plans/{id} [get]
// @Security
func GetPlan(c *gin.Context) {
	var Plan models.Plan
	Plan.ID, _ = utils.StringToInt64(c.Param("id"))
	result, err := Plan.Get()
	utils.HasError(err, "未找到相关数据", -1)
	app.OK(c, result, "")
}
