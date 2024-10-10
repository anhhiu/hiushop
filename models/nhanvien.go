package models

type NhanVien struct {
	MaNV     int `gorm:"primaryKey;autoiIncrement" json:"ma_nv"`
	HoTen    string `gorm:"size:50;not null" json:"ho_ten"`
	Email    string `gorm:"size:50" json:"email"`
	TaiKhoan string `gorm:"size:50;not null" json:"tai_khoan"`
	MatKhau  string `gorm:"size:50;not null" json:"mat_khau"`
}
