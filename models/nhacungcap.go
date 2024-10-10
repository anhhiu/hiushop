package models

//import "gorm.io/gorm"

// Bảng NhaCungCap
type NhaCungCap struct {
	MaNCC        int    `gorm:"primaryKey;autoIncrement" json:"ma_ncc"`
	TenCongTy    string `gorm:"size:50;not null" json:"ten_cong_ty"`
	Logo         string `gorm:"size:200;not null" json:"logo"`
	NguoiLienLac string `gorm:"size:50" json:"nguoi_lien_lac"`
	Email        string `gorm:"size:50;not null" json:"email"`
	DienThoai    string `gorm:"size:50" json:"dien_thoai"`
	DiaChi       string `gorm:"size:50" json:"dia_chi"`
	MoTa         string `gorm:"type:nvarchar(max)" json:"mo_ta"`
}
