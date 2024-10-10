package routes

import (
	"haoshop/controllers"
	"haoshop/docs"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// programmatically set swagger info
	docs.SwaggerInfo.BasePath = "/api"
	loai := r.Group("/api/loai")
	{
		loai.GET("/", controllers.GetAllLoai)
		loai.POST("/", controllers.CreateLoai)
		loai.GET("/:ma_loai", controllers.GetLoaiById)
		loai.PUT("/:ma_loai", controllers.UpdateLoai)
		loai.DELETE("/:ma_loai", controllers.DeleteLoai)
	}

	nhacungcap := r.Group("/api/nhacungcap")
	{
		nhacungcap.GET("/", controllers.GetAllNhacungcap)
		nhacungcap.POST("/", controllers.CreateNhacungcap)
		nhacungcap.GET("/:ma_ncc", controllers.GetNhacungcapById)
		nhacungcap.PUT("/:ma_ncc", controllers.UpdateNhacungcap)
		nhacungcap.DELETE("/:ma_ncc", controllers.DeleteNhacungcap)
	}

	hanghoa := r.Group("/api/hanghoa")
	{
		hanghoa.GET("/", controllers.GetAllHangHoa)
		hanghoa.POST("/", controllers.CreateHangHoa)
		hanghoa.GET("/:ma_hh", controllers.GetAllHangHoaById)
		hanghoa.PUT("/:ma_hh", controllers.UpdateHangHoa)
		hanghoa.DELETE("/:ma_hh", controllers.DeleteHangHoa)

	}

	khachhang := r.Group("/api/khachhang")
	{
		khachhang.GET("/", controllers.GetAllKhachhang)
		khachhang.POST("/", controllers.CreateKhachhang)
		khachhang.GET("/:ma_kh", controllers.GetAllKhachhangById)
		khachhang.PUT("/:ma_kh", controllers.UpdateKhachhang)
		khachhang.DELETE("/:ma_kh", controllers.DeleteKhachhang)
	}

	Trangthai := r.Group("/api/trangthai")
	{
		Trangthai.GET("/", controllers.GetAllTrangthai)
		Trangthai.POST("/", controllers.CreateTrangthai)
		Trangthai.GET("/:ma_trang_thai", controllers.GetAllTrangthaiById)
		Trangthai.PUT("/:ma_trang_thai", controllers.UpdateTrangthai)
		Trangthai.DELETE("/:ma_trang_thai", controllers.DeleteTrangthai)
	}

	nhanvien := r.Group("/api/nhanvien")
	{
		nhanvien.GET("/", controllers.GetAllNhanvien)
		nhanvien.POST("/", controllers.CreateNhanvien)
	}

	hoadon := r.Group("/api/hoadon")
	{
		hoadon.GET("/", controllers.GetAllHoaDon)
		hoadon.POST("/", controllers.CreateHoadon)
		hoadon.GET("/:ma_hd", controllers.GetHoaDonById)
		hoadon.PUT("/:ma_hd", controllers.UpdateHoaDon)
		hoadon.DELETE("/:ma_hd", controllers.DeleteHoaDon)
	}

	chitiethoadon := r.Group("/api/chitiethoadon")
	{
		chitiethoadon.GET("/", controllers.GetAllChiTietHoaDon)
		chitiethoadon.POST("/", controllers.CreateChiTietHoaDon)
		chitiethoadon.GET("/:ma_ct", controllers.GetChiTietHoaDonById)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
