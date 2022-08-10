package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/cloudinary-gin-api/dtos"
	"github.com/thitiphongD/cloudinary-gin-api/models"
	"github.com/thitiphongD/cloudinary-gin-api/services"
)

func FileUpload() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		formfile, _, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.MediaDto{
				Code:    http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": "Select a file to upload"},
			})
			return
		}

		uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.MediaDto{
				Code:    http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": "Error uploading file"},
			})
			return
		}

		ctx.JSON(http.StatusOK, dtos.MediaDto{
			Code:    http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": uploadUrl},
		})
	}
}

func RemoteUpload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var url models.Url

		if err := ctx.BindJSON(&url); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.MediaDto{
				Code:    http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.MediaDto{
				Code:    http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": "Error uploading file"},
			})
			return
		}

		ctx.JSON(http.StatusOK, dtos.MediaDto{
			Code:    http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": uploadUrl},
		})
	}
}
