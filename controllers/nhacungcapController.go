package controllers

import (
	"haoshop/databases"
	"haoshop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Tạo nhà cung cấp
// @Schemes
// @Description Thêm nhà cung cấp cho hàng hóa
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param nhacungcap body models.NhaCungCap true "NhaCungCap data"
// @Router /nhacungcap/ [post]
func CreateNhacungcap(c *gin.Context) {
	var input models.NhaCungCap

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	nhacungcap := models.NhaCungCap{
		TenCongTy:    input.TenCongTy,
		Logo:         input.Logo,
		NguoiLienLac: input.NguoiLienLac,
		Email:        input.Email,
		DienThoai:    input.DiaChi,
		DiaChi:       input.DiaChi,
		MoTa:         input.MoTa,
	}

	if err := databases.DB.Create(&nhacungcap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
		return
	}
	c.JSON(http.StatusCreated, nhacungcap)

}
// @Summary Hiển thị nhà cung cấp
// @Schemes
// @Description Hiển thị tất nhà cung cấp cho hàng hóa
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Router /nhacungcap/ [get]
func GetAllNhacungcap(c *gin.Context) {
	var nhacunngcaps []models.NhaCungCap
	if err := databases.DB.Find(&nhacunngcaps).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error})
		return
	}
	c.JSON(http.StatusOK, nhacunngcaps)
}
// @Summary Hiển thị nhà cung cấp theo mã nhà cung cấp
// @Schemes
// @Description Hiển thị nhà cung cấp theo mã nhà cung cấp
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param nhacungcap path int true "MaNCC"
// @Router /nhacungcap/{ma_ncc} [get]
func GetNhacungcapById(c *gin.Context) {
	var nhacungcap models.NhaCungCap
	if err := databases.DB.Where("ma_ncc = ?", c.Param("ma_ncc")).First(&nhacungcap).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, nhacungcap)
}
// @Summary Chỉnh sửa nhà cung cấp theo mã nhà cung cấp
// @Schemes
// @Description Chỉnh sửa nhà cung cấp theo mã nhà cung cấp
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param nhacungcap path int true "MaNCC"
// @Param nhacungcap body models.NhaCungCap true "NhaCungCap data"
// @Router /nhacungcap/{ma_ncc} [put]
func UpdateNhacungcap(c *gin.Context) {
	var input models.NhaCungCap

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var nhacungcap models.NhaCungCap
	if err := databases.DB.Where("ma_ncc = ?", c.Param("ma_ncc")).First(&nhacungcap).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	nhacungcap.TenCongTy = input.TenCongTy
	nhacungcap.Logo = input.Logo
	nhacungcap.NguoiLienLac = input.NguoiLienLac
	nhacungcap.Email = input.Email
	nhacungcap.DienThoai = input.DienThoai
	nhacungcap.DiaChi = input.DiaChi
	nhacungcap.MoTa = input.MoTa

	if err := databases.DB.Save(&nhacungcap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nhacungcap)
}
// @Summary Xóa nhà cung cấp theo mã nhà cung cấp
// @Schemes
// @Description Xóa nhà cung cấp theo mã nhà cung cấp
// @Tags Nhà Cung Cấp
// @Accept json
// @Produce json
// @Param nhacungcap path int true "MaNCC"
// @Router /nhacungcap/{ma_ncc} [delete]
func DeleteNhacungcap(c *gin.Context) {
	var nhacungcap models.NhaCungCap
	if err := databases.DB.Where("ma_ncc = ?", c.Param("ma_ncc")).First(&nhacungcap).Error; err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if err := databases.DB.Delete(&nhacungcap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Xóa thành công!")
}
