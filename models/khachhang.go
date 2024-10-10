package models

import "time"

// Bảng KhachHang
type KhachHang struct {
	MaKH      int       `gorm:"primaryKey;autoIncrement" json:"ma_kh"`
	HoTen     string    `gorm:"size:50;not null" json:"ho_ten"`
	GioiTinh  string    `gorm:"not null" json:"gioi_tinh"`
	NgaySinh  time.Time `gorm:"not null" json:"ngay_sinh"`
	DiaChi    string    `gorm:"size:60" json:"dia_chi"`
	DienThoai string    `gorm:"size:24" json:"dien_thoai"`
	Email     string    `gorm:"size:50;not null" json:"email"`
}
