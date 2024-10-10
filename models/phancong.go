package models

import "time"

type PhanCong struct {
	MaPC    int        `gorm:"primaryKey;autoIncrement" json:"ma_pc"`
	MaNV    string     `gorm:"size:50;not null" json:"ma_nv"`
	MaPB    string     `gorm:"size:7;not null" json:"ma_pb"`
	NgayPC  *time.Time `json:"ngay_pc"`
	HieuLuc *bool      `json:"hieu_luc"`

	NhanVien *NhanVien `gorm:"foreignKey:MaNV;references:MaNV;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"nhan_vien"`
	PhongBan *PhongBan `gorm:"foreignKey:MaPB;references:MaPB;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"phong_ban"`
}
