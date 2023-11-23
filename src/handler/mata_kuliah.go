package handler

import (
	"final-project-pemin/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) FindAllMataKuliah(c *gin.Context) {
	matakuliahs, err := r.service.MataKuliah.FindALl()
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "grabbed all matakuliah", "matakuliah", matakuliahs)
}
