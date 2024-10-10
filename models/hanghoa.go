package models

import "time"

type HangHoa struct {
	MaHH      int       `gorm:"primaryKey;autoIncrement" json:"ma_hh"`
	TenHH     string    `gorm:"type:nvarchar(50);not null" json:"ten_hh"`
	TenAlias  string    `gorm:"type:nvarchar(50)" json:"ten_alias"`
	MaLoai    int       `gorm:"not null" json:"ma_loai"`
	MoTaDonVi string    `gorm:"type:nvarchar(50)" json:"mo_ta_don_vi"`
	DonGia    float64   `json:"don_gia"`
	Hinh      string    `gorm:"type:nvarchar(max)" json:"hinh"`
	NgaySX    time.Time `gorm:"not null" json:"ngay_sx"`
	GiamGia   float64   `gorm:"not null" json:"giam_gia"`
	SoLanXem  int       `gorm:"not null" json:"so_lan_xem"`
	MoTa      string    `gorm:"type:nvarchar(max)" json:"mo_ta"`
	MaNCC     int       `gorm:"not null" json:"ma_ncc"`

	NhaCungCap *NhaCungCap `gorm:"foreignKey:MaNCC;references:MaNCC" json:"nha_cung_cap"`
	Loai       *Loai       `gorm:"foreignKey:MaLoai;references:MaLoai" json:"loai"`
}
