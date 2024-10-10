package controllers

import (
	"fmt"
	"haoshop/databases"
	"haoshop/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Tạo Hóa Đơn
// @Tags Hóa đơn
// @Param hoadon body models.HoaDon true "HoaDon data"
// @Router /hoadon/ [post]
func CreateHoadon(c *gin.Context) {
	var input struct {
		MaKH          int     `json:"ma_kh"`
		NgayDat       string  `json:"ngay_dat"`
		NgayCan       string  `json:"ngay_can"`
		NgayGiao      string  `json:"ngay_giao"`
		HoTen         string  `json:"ho_ten"`
		DiaChi        string  `json:"dia_chi"`
		CachThanhToan string  `json:"cach_thanh_toan"`
		CachVanChuyen string  `json:"cach_van_chuyen"`
		PhiVanChuyen  float64 `json:"phi_van_chuyen"`
		MaTrangThai   int     `json:"ma_trang_thai"`
		MaNV          int     `json:"ma_nv"`
		GhiChu        string  `json:"ghi_chu"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 02/01/2006
	ngaydat, err := time.Parse("02/01/2006", input.NgayDat)

	if err != nil {
		c.JSON(http.StatusBadRequest, "loi khong dung dinh dang dd/MM/yyyy")
		return
	}
	ngaycan, err := time.Parse("02/01/2006", input.NgayCan)

	if err != nil {
		c.JSON(http.StatusBadRequest, "loi khong dung dinh dang dd/MM/yyyy")
		return
	}
	ngaygiao, err := time.Parse("02/01/2006", input.NgayGiao)

	if err != nil {
		c.JSON(http.StatusBadRequest, "loi khong dung dinh dang dd/MM/yyyy")
		return
	}

	hoadon := models.HoaDon{
		MaKH:          input.MaKH,
		HoTen:         input.HoTen,
		DiaChi:        input.DiaChi,
		NgayDat:       ngaydat,
		NgayCan:       ngaycan,
		NgayGiao:      ngaygiao,
		CachVanChuyen: input.CachVanChuyen,
		CachThanhToan: input.CachThanhToan,
		PhiVanChuyen:  input.PhiVanChuyen,
		GhiChu:        input.GhiChu,
		MaNV:          input.MaNV,
		MaTrangThai:   input.MaTrangThai,
	}
	fmt.Println(input)
	if err := databases.DB.Create(&hoadon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, hoadon)
}

// @Summary Hiển thị tất cả hóa đơn
// @Tags Hóa đơn
// @Router /hoadon/ [get]
func GetAllHoaDon(c *gin.Context) {
	var hoadons []models.HoaDon
	if err := databases.DB.
		Preload("TrangThai").Preload("NhanVien").Preload("KhachHang").
		Find(&hoadons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hoadons)
}

// @Summary Hiển thị hóa đơn theo mã hóa đơn
// @Tags Hóa đơn
// @Param ma_hd path int true "MaHD"
// @Router /hoadon/{ma_hd} [get]
func GetHoaDonById(c *gin.Context) {
	var hoadon models.HoaDon
	if err := databases.DB.Preload("NhanVien").Preload("TrangThai").Preload("KhachHang").
		Where("ma_hd = ?", c.Param("ma_hd")).First(&hoadon).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, hoadon)
}

// @Summary Sửa hóa đơn theo mã hóa đơn
// @Tags Hóa đơn
// @Param ma_hd path int true "MaHD"
// @Param hoadon body models.HoaDon true "HoaDon data"
// @Router /hoadon/{ma_hd} [put]
func UpdateHoaDon(c *gin.Context) {
	var input struct {
		MaKH          int     `json:"ma_kh"`
		NgayDat       string  `json:"ngay_dat"`
		NgayCan       string  `json:"ngay_can"`
		NgayGiao      string  `json:"ngay_giao"`
		HoTen         string  `json:"ho_ten"`
		DiaChi        string  `json:"dia_chi"`
		CachThanhToan string  `json:"cach_thanh_toan"`
		CachVanChuyen string  `json:"cach_van_chuyen"`
		PhiVanChuyen  float64 `json:"phi_van_chuyen"`
		MaTrangThai   int     `json:"ma_trang_thai"`
		MaNV          int     `json:"ma_nv"`
		GhiChu        string  `json:"ghi_chu"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// 02/01/2006
	ngaydat, err := time.Parse("02/01/2006", input.NgayDat)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, "loi khong dung dinh dang dd/MM/yyyy")
		return
	}
	ngaycan, err := time.Parse("02/01/2006", input.NgayCan)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, "loi khong dung dinh dang dd/MM/yyyy")
		return
	}
	ngaygiao, err := time.Parse("02/01/2006", input.NgayGiao)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, "loi khong dung dinh dang dd/MM/yyyy")
		return
	}

	var hoadon models.HoaDon

	if err := databases.DB.Where("ma_hd = ?", c.Param("ma_hd")).First(&hoadon).Error; err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	hoadon.MaKH = input.MaKH
	hoadon.HoTen = input.HoTen
	hoadon.DiaChi = input.DiaChi
	hoadon.NgayDat = ngaydat
	hoadon.NgayCan = ngaycan
	hoadon.NgayGiao = ngaygiao
	hoadon.CachThanhToan = input.CachThanhToan
	hoadon.CachVanChuyen = input.CachVanChuyen
	hoadon.PhiVanChuyen = input.PhiVanChuyen
	hoadon.MaNV = input.MaNV
	hoadon.MaTrangThai = input.MaTrangThai
	hoadon.GhiChu = input.GhiChu

	if err := databases.DB.Save(&hoadon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, hoadon)

}

// @Summary Xóa hóa đơn theo mã hóa đơn
// @Tags Hóa đơn
// @Param ma_hd path int true "MaHD"
// @Router /hoadon/{ma_hd} [delete]
func DeleteHoaDon(c *gin.Context) {
	var hoadon models.HoaDon
	if err := databases.DB.Where("ma_hd = ?", c.Param("ma_hd")).First(&hoadon).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&hoadon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Xóa thành công!")

}
