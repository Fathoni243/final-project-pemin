package handler

import (
	"final-project-pemin/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) FindAllProdiHandler(c *gin.Context) {
	prodis, err := r.service.Prodi.FindAll()
	if err != nil {
		util.FailOrErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "grabbed all prodis", prodis)
}