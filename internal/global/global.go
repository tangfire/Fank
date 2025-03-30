package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB // 全局 db 对象
)

var (
	SysLog *logrus.Logger // 全局系统级日志对象
	BizLog *logrus.Entry  // 全局业务级日志对象
)
