package orm

import "github.com/FourWD/middleware/orm"

type Callback struct {
	ID string `json:"id" query:"id" db:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ArticleTypeID string `json:"article_type_id" query:"article_type_id" db:"article_type_id" gorm:"type:varchar(36)"`
	Description   string `json:"description" query:"description" db:"description" gorm:"type:varchar(100)"`
}
