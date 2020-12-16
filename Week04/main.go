package main

import "Week04/cmd"

/**
 * File :   main.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/12/16 20:11
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

// @title galaxy-mail
// @version 1.0.0
// @description MyApp 系统
// @license.name MIT
// @license.url https://www.xuxueyun.com

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
