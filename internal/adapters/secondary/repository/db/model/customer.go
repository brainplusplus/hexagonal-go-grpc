package model

import (
	"gorm.io/plugin/soft_delete"
)

const TableNameCustomer = "customers"

type Customer struct {
	ID        int32                 `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string                `gorm:"column:name;type:varchar(255);not null"`
	Email     string                `gorm:"column:email;type:varchar(255);unique;not null"`
	Phone     string                `gorm:"column:phone;type:varchar(20);not null"`
	Gender    string                `gorm:"column:gender;type:varchar(10);not null"`
	IsActive  bool                  `gorm:"column:is_active;default:true"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;softDelete:milli"`
}

func (*Customer) TableName() string {
	return TableNameCustomer
}
