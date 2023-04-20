package categorydto

type CategoryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name"binding:"required, name" gorm:"unique; not null"`
}

type CategoryUpdate struct {
	Name string `json:"name" gorm:"type: varchar(255)"`
}
