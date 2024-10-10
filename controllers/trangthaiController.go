package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Trangthai
// @Summary Thêm trạng thái
// @Schemes
// @Description Thêm trạng thái
// @Tags Trạng thái
// @Accept json
// @Produce json
// @Param trangthai body models.TrangThai true "TrangThai data"
// @Router /trangthai/ [post]
func CreateTrangthai(c *gin.Context) {
	var input models.TrangThai

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	trangthai := models.TrangThai{
		TenTrangThai: input.TenTrangThai,
		MoTa:         input.MoTa,
	}

	if err := databases.DB.Create(&trangthai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, trangthai)
}

// @Summary Hiển thị trạng thái
// @Schemes
// @Description Hiển thị trạng thái
// @Tags Trạng thái
// @Accept json
// @Produce json
// @Router /trangthai/ [get]
func GetAllTrangthai(c *gin.Context) {
	var trangthais []models.TrangThai
	if err := databases.DB.Find(&trangthais).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, trangthais)
}

// @Summary Hiển thị trạng thái theo mã trạng thái
// @Schemes
// @Description Hiển thị trạng thái theo mã trạng thái
// @Tags Trạng thái
// @Accept json
// @Produce json
// @Param ma_trang_thai path int true "MaTrangThai"
// @Router /trangthai/{ma_trang_thai} [get]
func GetAllTrangthaiById(c *gin.Context) {
	var trangthai models.TrangThai
	if err := databases.DB.Where("ma_trang_thai = ?", c.Param("ma_trang_thai")).First(&trangthai).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, trangthai)
}

// @Summary Chỉnh sửa trạng thái theo mã trạng thái
// @Schemes
// @Description Chỉnh sửa trạng thái theo mã trạng thái
// @Tags Trạng thái
// @Accept json
// @Produce json
// @Param ma_trang_thai path int true "ma_trang_thai"
// @Param trangthai body models.TrangThai true "TrangThai data"
// @Router /trangthai/{ma_trang_thai} [put]
func UpdateTrangthai(c *gin.Context) {
	var input models.TrangThai

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var trangthai models.TrangThai

	if err := databases.DB.Where("ma_trang_thai = ?", c.Param("ma_trang_thai")).First(&trangthai).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Model(&trangthai).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, trangthai)
}

// @Summary Xóa trạng thái theo mã trạng thái
// @Schemes
// @Description Xóa trạng thái theo mã trạng thái
// @Tags Trạng thái
// @Accept json
// @Produce json
// @Param ma_trang_thai path int true "MaTrangThai"
// @Router /trangthai/{ma_trang_thai} [delete]
func DeleteTrangthai(c *gin.Context) {
	var trangthai models.TrangThai
	if err := databases.DB.Where("ma_trang_thai = ?", c.Param("ma_trang_thai")).First(&trangthai).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if err := databases.DB.Delete(&trangthai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, true)
}
