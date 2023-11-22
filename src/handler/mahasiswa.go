package handler

import (
	"final-project-pemin/middleware"
	"final-project-pemin/src/model"
	"final-project-pemin/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *rest) RegisterMahasiswa(c *gin.Context) {
	var req model.MahasiswaInputRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	newMahasiswa, err := r.service.Mahasiswa.Create(&req)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "create mahasiswa success", "mahasiswa", newMahasiswa)
}

func (r *rest) LoginMahasiswa(c *gin.Context) {
	var req model.MahasiswaLoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	mahasiswaLogin, err := r.service.Mahasiswa.Login(&req)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tokenJwt, err := middleware.GenerateToken(mahasiswaLogin.NIM)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "mahasiswa success login", "token", tokenJwt)
}

func (r *rest) FindMahasiswaByNIM(c *gin.Context) {
	nim := c.Param("nim")

	mahasiswa, err := r.service.Mahasiswa.FindByNIM(nim)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "success get data mahasiswa", "mahasiswa", mahasiswa)
}

func (r *rest) FindAllMahasiswa(c *gin.Context) {
	mahasiswas, err := r.service.Mahasiswa.FindAll()
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "grabbed all mahasiswa", "mahasiswa", mahasiswas)
}

func (r *rest) GetProfileMahasiswa(c *gin.Context) {
	nimToken := c.MustGet("nim").(string)

	mahasiswa, err := r.service.Mahasiswa.FindByNIM(nimToken)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "success get data mahasiswa", "mahasiswa", mahasiswa)
}

func (r *rest) SaveMatkulMahasiswa(c *gin.Context) {
	nimToken := c.MustGet("nim").(string)

	idString := c.Param("mkId")
	mkId, err := strconv.Atoi(idString)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	mahasiswa, errorSave := r.service.Mahasiswa.SaveMatkul(nimToken, int64(mkId))
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, errorSave.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "success add mata kuliah", "mahasiswa", mahasiswa)
}

func (r *rest) DeleteMatkulMahasiswa(c *gin.Context) {
	nimToken := c.MustGet("nim").(string)

	idString := c.Param("mkId")
	mkId, err := strconv.Atoi(idString)
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	mahasiswa, errorSave := r.service.Mahasiswa.DeleteMatkul(nimToken, int64(mkId))
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, errorSave.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "success delete mata kuliah", "mahasiswa", mahasiswa)
}
