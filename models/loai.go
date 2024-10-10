package models

// Bảng Loai
type Loai struct {
	MaLoai       int    `gorm:"primaryKey;autoIncrement" json:"ma_loai"`
	TenLoai      string `gorm:"size:50;not null" json:"ten_loai"`
	TenLoaiAlias string `gorm:"size:50" json:"ten_loai_alias"`
	MoTa         string `gorm:"type:nvarchar(max)" json:"mo_ta"`
	Hinh         string `gorm:"size:100" json:"hinh"`
}
