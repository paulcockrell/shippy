package go_micro_service_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate - Generate UUID in place of auto-incrementing DB ID column value
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuidGen := uuid.NewV4()
	return scope.SetColumn("id", uuidGen)
}
