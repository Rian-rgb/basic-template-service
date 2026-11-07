package payment_type

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"type:varchar(50);not null" json:"name"`
	IsActive bool      `gorm:"type:boolean;not null;default:true" json:"is_active"`
}

func (*Entity) TableName() string {
	return "payment_types"
}
