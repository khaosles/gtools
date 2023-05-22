package g

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelInterface interface{}

type BaseModelInt struct {
	ID        int            `json:"id" gorm:"type:int;comment:主键"` // 主键ID
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`   // 删除时间
	Remarks   string         `json:"remarks" gorm:"comment:备注"`     // 备注
}

type BaseModel struct {
	ID         string         `json:"id" gorm:"type:varchar(32);comment:主键"`         // 主键ID
	CreateTime time.Time      `json:"createTime" gorm:"autoCreateTime;comment:创建时间"` // 创建时间
	UpdateTime time.Time      `json:"updateTime" gorm:"autoUpdateTime;comment:更新时间"` // 更新时间
	DeleteTime gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                   // 删除标记
	Remarks    string         `json:"remarks" gorm:"comment:备注"`                     // 备注
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = strings.Replace(uuid.New().String(), "-", "", 4)
	return nil
}
