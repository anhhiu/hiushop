package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
// @Summary Thêm hàng hóa
// @Schemes
// @Description Thêm hàng hóa
// @Tags Hàng hóa
// @Accept json
// @Produce json
// @Param hanghoa body models.HangHoa true "HangHoa data"
// @Router /hanghoa/ [post]
func CreateHangHoa(c *gin.Context) {
	var input models.HangHoa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hanghoa := models.HangHoa{
		TenHH:     input.TenHH,
		TenAlias:  input.TenAlias,
		MaLoai:    input.MaLoai,
		MoTaDonVi: input.MoTaDonVi,
		DonGia:    input.DonGia,
		Hinh:      input.Hinh,
		NgaySX:    time.Now(),
		GiamGia:   input.GiamGia,
		SoLanXem:  input.SoLanXem,
		MoTa:      input.MoTa,
		MaNCC:     input.MaNCC,
	}

	if err := databases.DB.Create(&hanghoa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, hanghoa)
}
// @Summary Hiển thị hàng hóa
// @Schemes
// @Description Hiển thị hàng hóa
// @Tags Hàng hóa
// @Accept json
// @Produce json
// @Router /hanghoa/ [get]
func GetAllHangHoa(c *gin.Context) {
	var hanghoas []models.HangHoa
	if err := databases.DB.Preload("Loai").Preload("NhaCungCap").Find(&hanghoas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, hanghoas)
}
// @Summary Hiển thị hàng hóa theo mã hàng hóa
// @Schemes
// @Description Hiển thị hàng hóa theo mã hàng hóa
// @Tags Hàng hóa
// @Accept json
// @Produce json
// @Param ma_hh path int true "MaHH"
// @Router /hanghoa/{ma_hh} [get]
func GetAllHangHoaById(c *gin.Context) {
	var hanghoa models.HangHoa
	if err := databases.DB.Preload("Loai").Preload("NhaCungCap").Where("ma_hh = ?", c.Param("ma_hh")).First(&hanghoa).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, hanghoa)
}

func GetAllHangHoa2(c *gin.Context) {
	var hanghoas []models.HangHoa

	if err := databases.DB.
		Preload("Loai", "ten_loai_alias like ?", "%bimbim%").
		Preload("NhaCungCap").
		Find(&hanghoas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, hanghoas)
}
// @Summary Chỉnh sửa hàng hóa theo mã hàng hóa
// @Schemes
// @Description Chỉnh sửa hàng hóa theo mã hàng hóa
// @Tags Hàng hóa
// @Accept json
// @Produce json
// @Param ma_hh path int true "MaHH"
// @Param hanghoa body models.HangHoa true "HangHoa data"
// @Router /hanghoa/{ma_hh} [put]
func UpdateHangHoa(c *gin.Context) {
	var input models.HangHoa

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var hanghoa models.HangHoa

	if err := databases.DB.Where("ma_hh = ?", c.Param("ma_hh")).First(&hanghoa).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Model(&hanghoa).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, hanghoa)
}
// @Summary Xóa hàng hóa theo mã hàng hóa
// @Schemes
// @Description Xóa hàng hóa theo mãhàng hóa
// @Tags Hàng hóa
// @Accept json
// @Produce json
// @Param ma_hh path int true "MaHH"
// @Router /hanghoa/{ma_hh} [delete]
func DeleteHangHoa(c *gin.Context) {
	var hanghoa models.HangHoa
	if err := databases.DB.Where("ma_hh = ?", c.Param("ma_hh")).First(&hanghoa).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&hanghoa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
