package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Khachhang
// @Summary Thêm khách hàng
// @Schemes
// @Description Thêm khách hàng
// @Tags Khách hàng
// @Accept json
// @Produce json
// @Param khachhang body models.KhachHang true "khachhang data"
// @Router /khachhang/ [post]
func CreateKhachhang(c *gin.Context) {

	var input struct {
		HoTen     string `json:"ho_ten"`
		GioiTinh  string `json:"gioi_tinh"`
		NgaySinh  string `json:"ngay_sinh"`
		DienThoai string `json:"dien_thoai"`
		Email     string `json:"email"`
		DiaChi    string `json:"dia_chi"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Chuyển đổi định dạng ngày sinh
	ngaySinh, err := time.Parse("02/01/2006", input.NgaySinh)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Ngày sinh không hợp lệ. Vui lòng sử dụng định dạng dd/MM/yyyy")
		return
	}
	khachhang := models.KhachHang{
		HoTen:     input.HoTen,
		GioiTinh:  input.GioiTinh,
		NgaySinh:  ngaySinh,
		DienThoai: input.DienThoai,
		Email:     input.Email,
		DiaChi:    input.DiaChi,
	}

	if err := databases.DB.Create(&khachhang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, khachhang)
}

// @Summary Hiển thị khách hàng
// @Schemes
// @Description Hiển thị khách hàng
// @Tags Khách hàng
// @Accept json
// @Produce json
// @Router /khachhang/ [get]
func GetAllKhachhang(c *gin.Context) {
	var khachhangs []models.KhachHang
	if err := databases.DB.Find(&khachhangs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, khachhangs)
}

// @Summary Hiển thị khách hàng theo mã khách hàng
// @Schemes
// @Description Hiển thị khách hàng theo mã khách hàng
// @Tags Khách hàng
// @Accept json
// @Produce json
// @Param ma_kh path int true "MaHH"
// @Router /khachhang/{ma_kh} [get]
func GetAllKhachhangById(c *gin.Context) {
	var khachhang models.KhachHang
	if err := databases.DB.Where("ma_kh = ?", c.Param("ma_kh")).First(&khachhang).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, khachhang)
}

// @Summary Chỉnh sửa khách hàng theo mã khách hàng
// @Schemes
// @Description Chỉnh sửa khách hàng theo mã khách hàng
// @Tags Khách hàng
// @Accept json
// @Produce json
// @Param ma_kh path int true "MaHH"
// @Param khachhang body models.KhachHang true "khachhang data"
// @Router /khachhang/{ma_kh} [put]
func UpdateKhachhang(c *gin.Context) {
	var input models.KhachHang

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var khachhang models.KhachHang

	if err := databases.DB.Where("ma_kh = ?", c.Param("ma_kh")).First(&khachhang).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Model(&khachhang).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, khachhang)
}

// @Summary Xóa khách hàng theo mã khách hàng
// @Schemes
// @Description Xóa khách hàng theo mã khách hàng
// @Tags Khách hàng
// @Accept json
// @Produce json
// @Param ma_kh path int true "MaHH"
// @Router /khachhang/{ma_kh} [delete]
func DeleteKhachhang(c *gin.Context) {
	var khachhang models.KhachHang
	if err := databases.DB.Where("ma_kh = ?", c.Param("ma_kh")).First(&khachhang).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&khachhang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
