package model

const TableNameDynamic = "dynamic"

const (
	DynamicCategoryUnknown = iota
	DynamicCategoryImage
	DynamicCategoryVideo
)

type Dynamic struct {
	BaseModel
	Title     string      `json:"title" gorm:"column:title;type:varchar(255);index:index_title;not null;default:'';comment:标题"`
	Content   string      `json:"content" gorm:"column:content;type:text;comment:内容"`
	Resources StringSlice `json:"resources" gorm:"column:resources;type:text;not null;comment:资源(图片或视频)"`
	Category  int8        `json:"category" gorm:"column:category;type:tinyint(1);not null;comment:1图片 2视频"`
}

func (Dynamic) TableName() string {
	return TableNameDynamic
}
