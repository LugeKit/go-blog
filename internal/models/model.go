package models

type Model struct {
	ID        uint32 `gorm:"primary_key" json:"id"`
	CreatedAt uint32 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt uint32 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
	DeletedAt uint32 `json:"deleted_at"`
}
