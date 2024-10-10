package models

import "time"

type HoaDon struct {
	MaHD          int        `gorm:"primaryKey;autoIncrement" json:"ma_hd"`
	MaKH          int        `gorm:"not null" json:"ma_kh"`
	NgayDat       time.Time `gorm:"not null" json:"ngay_dat"`
	NgayCan       time.Time `json:"ngay_can"`
	NgayGiao      time.Time `json:"ngay_giao"`
	HoTen         string     `gorm:"size:50" json:"ho_ten"`
	DiaChi        string     `gorm:"size:60;not null" json:"dia_chi"`
	CachThanhToan string     `gorm:"size:100;not null" json:"cach_thanh_toan"`
	CachVanChuyen string     `gorm:"size:100;not null" json:"cach_van_chuyen"`
	PhiVanChuyen  float64    `gorm:"not null" json:"phi_van_chuyen"`
	MaTrangThai   int        `gorm:"not null" json:"ma_trang_thai"`
	MaNV          int        `gorm:"not null" json:"ma_nv"`
	GhiChu        string     `gorm:"size:50" json:"ghi_chu"`

	KhachHang *KhachHang `gorm:"foreignKey:MaKH;references:MaKH;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"khach_hang"`
	NhanVien  *NhanVien  `gorm:"foreignKey:MaNV;references:MaNV;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"nhan_vien"`
	TrangThai *TrangThai `gorm:"foreignKey:MaTrangThai;references:MaTrangThai;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"trang_thai"`
}
