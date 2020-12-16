package utils

import "strconv"

/**
 * File :   string.go
 * Author:  xuxueyun
 * Version: 1.0.0
 * Date:    2020/12/16 19:49
 * Copyright: 2020 DanielXU<i@xuxueyun.com>
 * Description:
 */

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}
