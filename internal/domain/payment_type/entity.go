package payment_type

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"type:varchar(50);not null"`
	IsActive bool      `gorm:"type:boolean;not null;default:true"`
}

func (*Entity) TableName() string {
	return "payment_types"
}
