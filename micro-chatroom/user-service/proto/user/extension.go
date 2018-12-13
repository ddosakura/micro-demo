package go_micro_srv_user

import (
	"log" // 或可使用　"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// 函数 BeforeCreate() 指定了 GORM 库使用 uuid 作为 ID 列值
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("created uuid error: %v\n", err)
	}
	// 选择主键列
	return scope.SetColumn("UserId", uuid.String())
}
