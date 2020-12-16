package router

/**
 * File :   router.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/12/16 19:46
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

import (
	"Week04/api/plan"
	"Week04/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.GinEngine == nil {
		r = gin.New()
	} else {
		r = global.GinEngine
	}

	InitSysRouter(r)
	return r
}

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/api/v1")
	// 注册业务（组）路由
	registerBizRouter(g)
	return g
}

func registerBizRouter(v1 *gin.RouterGroup) {
	r := v1.Group("/biz")
	{
		// plan
		r.GET("/plans/:id", plan.GetPlan)
	}
}
