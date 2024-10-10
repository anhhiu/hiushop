package models

type PhongBan struct {
	MaPB     int    `gorm:"primaryKey;autoIncrement" json:"ma_pb"`
	TenPB    string `gorm:"size:50;not null" json:"ten_pb"`
	ThongTin string `gorm:"type:text" json:"thong_tin"`
}
