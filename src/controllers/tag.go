package controllers

import (
	"net/http"

	"github.com/Vlad-Peresta/todo_list_go/src/models"
	"github.com/Vlad-Peresta/todo_list_go/src/schemas"
	"github.com/gin-gonic/gin"
)

// CreateTag godoc
//
//	@Summary		Create Tag record
//	@Description	Create Tag record
//	@Tags			tags
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			Request Body 	body		schemas.TagRequest  	true	"Request Body"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Router			/tags [POST]
//
// CreateTag creates Tag record in the database
func CreateTag(context *gin.Context) {
	var tag models.Tag

	if err := context.ShouldBindJSON(&tag); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
		return
	}

	if err := models.CreateRecord(&tag); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
		return
	}

	context.JSON(http.StatusCreated, schemas.Response{Status: "success", Data: tag})
}

// GetAllTags godoc
//
//	@Summary		Get all Tag records
//	@Description	Get all Tag records
//	@Tags			tags
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
// @Param        sort    query     string  false  "Sorting parameter"  default(id DESC)
// @Param        limit    query     string  false  "Records per page" default(100)
// @Param        page    query     string  false  "Current page" default(1)
//
//	@Produce		json
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Router			/tags [GET]
//
// GetAllTags finds all Tag records
func GetAllTags(context *gin.Context) {
	var tags []models.Tag
	var pagination models.Pagination

	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
		return
	}

	if err := models.GetAllRecords(&tags, &pagination); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, schemas.Response{Status: "success", Data: tags})
}

// GetTag godoc
//
//	@Summary		Get Tag record by ID
//	@Description	Get Tag record by ID
//	@Tags			tags
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			id	path		int	true	"Tag ID"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Router			/tags/{id} [GET]
//
// GetTag finds Tag record by ID
func GetTag(context *gin.Context) {
	var tag models.Tag

	if err := models.GetRecordByID(&tag, context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, schemas.Response{Status: "success", Data: tag})
}

// UpdateTag godoc
//
//	@Summary		Update Tag record
//	@Description	Update Tag record
//	@Tags			tags
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			id	path		int	true	"Tag ID"
//	@Param			Request Body 	body		schemas.TagRequest  	true	"Request Body"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Router			/tags/{id} [PUT]
//
// UpdateTag updates Tag record by ID
func UpdateTag(context *gin.Context) {
	var data schemas.TagRequest
	var tag models.Tag

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
	}

	if err := models.PatchUpdateTagByID(&tag, &data, context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
	}

	context.JSON(http.StatusOK, schemas.Response{Status: "success", Data: tag})
}

// DeleteTag godoc
//
//	@Summary		Delete Tag record
//	@Description	Delete Tag record
//	@Tags			tags
//
// @Param Authorization header string true "Insert your access token" default(Bearer <Access token>)
//
//	@Produce		json
//	@Param			id	path		int	true	"Tag ID"
//	@Success		200	{object}	schemas.Response
//	@Failure		400	{object}	schemas.Response
//	@Router			/tags/{id} [DELETE]
//
// DeleteTag deletes Tag record by ID
func DeleteTag(context *gin.Context) {
	var tag models.Tag
	id := context.Param("id")

	if err := models.DeleteRecordByID(&tag, id); err != nil {
		context.JSON(http.StatusBadRequest, schemas.Response{Status: "error", Message: err.Error()})
	}

	context.JSON(http.StatusOK, schemas.Response{Status: "success", Message: "record was deleted successfully", Data: id})
}
