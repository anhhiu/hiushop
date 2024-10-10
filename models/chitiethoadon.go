package models

type ChiTietHD struct {
	MaCT    int     `gorm:"primaryKey;autoIncrement" json:"ma_ct"`
	MaHD    int     `gorm:"not null" json:"ma_hd"`
	MaHH    int     `gorm:"not null" json:"ma_hh"`
	DonGia  float64 `gorm:"not null" json:"don_gia"`
	SoLuong int     `gorm:"not null" json:"so_luong"`
	GiamGia float64 `gorm:"not null" json:"giam_gia"`

	HoaDon  *HoaDon  `gorm:"foreignKey:MaHD;references:MaHD" json:"hoa_don"`
	HangHoa *HangHoa `gorm:"foreignKey:MaHH;references:MaHH" json:"hang_hoa"`
}
