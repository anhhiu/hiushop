package models

// Bảng PhanQuyen
type PhanQuyen struct {
	MaPQ    int    `gorm:"primaryKey;autoIncrement" json:"ma_pq"`
	MaPB    string `gorm:"size:7" json:"ma_pb"`
	MaTrang int    `json:"ma_trang"`

	Them bool `gorm:"not null" json:"them"`
	Sua  bool `gorm:"not null" json:"sua"`
	Xoa  bool `gorm:"not null" json:"xoa"`
	Xem  bool `gorm:"not null" json:"xem"`

	PhongBan *PhongBan `gorm:"foreignKey:MaPB;references:MaPB;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"phong_ban"`
	TrangWeb *TrangWeb `gorm:"foreignKey:MaTrang;references:MaTrang;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"trang_web"`
}
