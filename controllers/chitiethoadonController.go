package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Hiển thị tất cả các hóa đơn
// @Tags Chi tiết hóa đơn
// @Router /chitiethoadon/ [get]
func GetAllChiTietHoaDon(c *gin.Context) {
	var chitiethd []models.ChiTietHD
	if err := databases.DB.
		Preload("HangHoa").
		Preload("HoaDon").
		Find(&chitiethd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chitiethd)
}

// @Summary Thêm hóa đơn chi tiết
// @Tags Chi tiết hóa đơn
// @Param chitiethd body models.ChiTietHD true "ChiTietHD data"
// @Router /chitiethoadon/ [post]
func CreateChiTietHoaDon(c *gin.Context) {
	var input struct {
		MaHD    int     `json:"ma_hd"`
		MaHH    int     `json:"ma_hh"`
		DonGia  float64 `json:"don_gia"`
		SoLuong int     `json:"so_luong"`
		GiamGia float64 `json:"giam_gia"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	chitiethd := models.ChiTietHD{
		MaHD:    input.MaHD,
		MaHH:    input.MaHH,
		DonGia:  input.DonGia,
		SoLuong: input.SoLuong,
		GiamGia: input.GiamGia,
	}
	if err := databases.DB.Create(&chitiethd).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, chitiethd)

	
}

// @Summary Hiển thị tất cả các chi tiết hóa đơn theo id
// @Tags Chi tiết hóa đơn
// @Param ma_ct path int true "MaHD"
// @Router /chitiethoadon/{ma_ct} [get]
func GetChiTietHoaDonById(c *gin.Context) {
	var chitiethd models.ChiTietHD
	err := databases.DB.
		Preload("HoaDon").
		Preload("HangHoa").
		Where("ma_ct = ?", c.Param("ma_ct")).
		First(&chitiethd).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chitiethd)
}
