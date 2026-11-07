package payment_type

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(50);not null" json:"name"`
	IsActive  bool       `gorm:"type:boolean;not null;default:true" json:"is_active"`
	CreatedAt time.Time  `gorm:"type:timestamptz;not null" json:"created_at"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;autoUpdateTime:false" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"type:uuid" json:"updated_by,omitempty"`
}

func (*Entity) TableName() string {
	return "payment_types"
}
