package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
//@Summary Hiển thị danh sách nhân viên
//@Tags Nhân Viên
//@Router /nhanvien/ [get]
func GetAllNhanvien(c *gin.Context) {
	var nhanviens []models.NhanVien

	if err := databases.DB.Find(&nhanviens).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nhanviens)
}
//@Summary Tạo tài khoản cho nhân viên
//@Tags Nhân Viên
//@Param nhanvien body models.NhanVien true "NhanVien data"
//@router /nhanvien/ [post]
func CreateNhanvien(c *gin.Context) {
	var input models.NhanVien

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	nhanvien := models.NhanVien{
		HoTen:    input.HoTen,
		Email:    input.Email,
		TaiKhoan: input.TaiKhoan,
		MatKhau:  input.MatKhau,
	}

	if err := databases.DB.Create(&nhanvien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nhanvien)
}
