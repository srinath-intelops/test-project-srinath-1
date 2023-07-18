package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/srinath-intelops/test-project-srinath-1/user-svc-1/pkg/rest/server/models"
	"github.com/srinath-intelops/test-project-srinath-1/user-svc-1/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type UserdetailsController struct {
	userdetailsService *services.UserdetailsService
}

func NewUserdetailsController() (*UserdetailsController, error) {
	userdetailsService, err := services.NewUserdetailsService()
	if err != nil {
		return nil, err
	}
	return &UserdetailsController{
		userdetailsService: userdetailsService,
	}, nil
}

func (userdetailsController *UserdetailsController) CreateUserdetails(context *gin.Context) {
	// validate input
	var input models.Userdetails
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger userdetails creation
	if _, err := userdetailsController.userdetailsService.CreateUserdetails(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Userdetails created successfully"})
}

func (userdetailsController *UserdetailsController) UpdateUserdetails(context *gin.Context) {
	// validate input
	var input models.Userdetails
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger userdetails update
	if _, err := userdetailsController.userdetailsService.UpdateUserdetails(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Userdetails updated successfully"})
}

func (userdetailsController *UserdetailsController) FetchUserdetails(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger userdetails fetching
	userdetails, err := userdetailsController.userdetailsService.GetUserdetails(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userdetails)
}

func (userdetailsController *UserdetailsController) DeleteUserdetails(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger userdetails deletion
	if err := userdetailsController.userdetailsService.DeleteUserdetails(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Userdetails deleted successfully",
	})
}

func (userdetailsController *UserdetailsController) ListUserdetails(context *gin.Context) {
	// trigger all userdetails fetching
	userdetails, err := userdetailsController.userdetailsService.ListUserdetails()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userdetails)
}

func (*UserdetailsController) PatchUserdetails(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*UserdetailsController) OptionsUserdetails(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*UserdetailsController) HeadUserdetails(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
