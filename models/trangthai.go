package models

// Bảng TrangThai
type TrangThai struct {
	MaTrangThai  int    `gorm:"primaryKey;autoIncrement" json:"ma_trang_thai"`
	TenTrangThai string `gorm:"size:50;not null" json:"ten_trang_thai"`
	MoTa         string `gorm:"type:nvarchar(max)" json:"mo_ta"`
}
