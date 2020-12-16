package global

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var GinEngine *gin.Engine

var Eloquent *gorm.DB
var (
	Source string
	Driver string
	DBName string
)

var (
	Version   = "The make command should be used to compile the application"
	BuildTime = "The make command should be used to compile the application"
	GitHash   = "The make command should be used to compile the application"
	GoVersion = "The make command should be used to compile the application"
)
