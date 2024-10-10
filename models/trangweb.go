package models

type TrangWeb struct {
	MaTrang  int    `gorm:"primaryKey;autoIncrement" json:"ma_trang"`
	TenTrang string `gorm:"size:200" json:"ten_trang"`
	URL      string `gorm:"size:500" json:"url"`
}
