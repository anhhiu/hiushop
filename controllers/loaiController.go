package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Thêm Loại
// @Schemes
// @Description Thêm loại hàng hóa
// @Tags Loại
// @Accept json
// @Produce json
// @Param loai body models.Loai true "Loai data"
// @Router /loai/ [post]
func CreateLoai(c *gin.Context) {
	var input models.Loai

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	loai := models.Loai{
		TenLoai:      input.TenLoai,
		TenLoaiAlias: input.TenLoaiAlias,
		MoTa:         input.MoTa,
		Hinh:         input.Hinh,
	}

	if err := databases.DB.Create(&loai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
		return
	}
	c.JSON(http.StatusCreated, loai)

}

// @Summary Lấy ra tất cả các loại hàng hóa
// @Schemes
// @Description Hiển thị tất cả các loại hàng hóa
// @Tags Loại
// @Accept json
// @Produce json
// @Router /loai/ [get]
func GetAllLoai(c *gin.Context) {
	var loais []models.Loai
	if err := databases.DB.Find(&loais).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error})
		return
	}
	c.JSON(http.StatusOK, loais)
}

// @Summary Lấy loại hàng hóa theo mã hàng hóa
// @Schemes
// @Description Hiển thị loại hàng hóa theo mã hàng hóa
// @Tags Loại
// @Accept json
// @Produce json
// @Param ma_loai path int true "MaLoai"
// @Router /loai/{ma_loai} [get]
func GetLoaiById(c *gin.Context) {
	var loai models.Loai
	if err := databases.DB.Where("ma_loai = ?", c.Param("ma_loai")).First(&loai).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, loai)
}

// @Summary Sửa loại hàng hóa theo mã hàng hóa
// @Schemes
// @Description Chỉnh sửa loại hàng hóa theo mã hàng hóa
// @Tags Loại
// @Accept json
// @Produce json
// @Param ma_loai path int true "MaLoai"
// @Param loai body models.Loai true "Loai data"
// @Router /loai/{ma_loai} [put]
func UpdateLoai(c *gin.Context) {
	var input models.Loai

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var loai models.Loai
	if err := databases.DB.Where("ma_loai = ?", c.Param("ma_loai")).First(&loai).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	loai.TenLoai = input.TenLoai
	loai.TenLoaiAlias = input.TenLoaiAlias
	loai.MoTa = input.MoTa
	loai.Hinh = input.Hinh

	if err := databases.DB.Save(&loai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, loai)
}

// @Summary Xóa loại hàng hóa
// @Schemes
// @Description Xóa loại hàng hóa theo mã hàng hóa
// @Tags Loại
// @Accept json
// @Produce json
// @Param ma_loai path int true "Maloai"
// @Router /loai/{ma_loai} [delete]
func DeleteLoai(c *gin.Context) {
	var loai models.Loai
	if err := databases.DB.Where("ma_loai = ?", c.Param("ma_loai")).First(&loai).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := databases.DB.Delete(&loai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Xóa thành công!")
}
