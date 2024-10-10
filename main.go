package main

import (
	"fmt"
	"haoshop/databases"
	"haoshop/models"
	"haoshop/routes"
	"log"
)

func main() {
	databases.ConnectDatabase()

	databases.DB.AutoMigrate(&models.Loai{})
	databases.DB.AutoMigrate(&models.NhaCungCap{})
	databases.DB.AutoMigrate(&models.TrangWeb{})
	databases.DB.AutoMigrate(&models.TrangThai{})
	databases.DB.AutoMigrate(&models.KhachHang{})
	databases.DB.AutoMigrate(&models.PhongBan{})
	databases.DB.AutoMigrate(&models.NhanVien{})
	databases.DB.AutoMigrate(&models.PhanQuyen{})
	databases.DB.AutoMigrate(&models.PhanCong{})
	databases.DB.AutoMigrate(&models.HangHoa{})
	databases.DB.AutoMigrate(&models.HoaDon{})
	databases.DB.AutoMigrate(&models.ChiTietHD{})

	r := routes.SetupRouter()
	fmt.Println("ket noi voi cong 8083")
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Lỗi không thể khởi chạy server: %v", err)
	}

}
