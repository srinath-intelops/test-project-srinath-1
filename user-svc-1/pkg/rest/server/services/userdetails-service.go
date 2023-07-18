package services

import (
	"github.com/srinath-intelops/test-project-srinath-1/user-svc-1/pkg/rest/server/daos"
	"github.com/srinath-intelops/test-project-srinath-1/user-svc-1/pkg/rest/server/models"
)

type UserdetailsService struct {
	userdetailsDao *daos.UserdetailsDao
}

func NewUserdetailsService() (*UserdetailsService, error) {
	userdetailsDao, err := daos.NewUserdetailsDao()
	if err != nil {
		return nil, err
	}
	return &UserdetailsService{
		userdetailsDao: userdetailsDao,
	}, nil
}

func (userdetailsService *UserdetailsService) CreateUserdetails(userdetails *models.Userdetails) (*models.Userdetails, error) {
	return userdetailsService.userdetailsDao.CreateUserdetails(userdetails)
}

func (userdetailsService *UserdetailsService) UpdateUserdetails(id int64, userdetails *models.Userdetails) (*models.Userdetails, error) {
	return userdetailsService.userdetailsDao.UpdateUserdetails(id, userdetails)
}

func (userdetailsService *UserdetailsService) DeleteUserdetails(id int64) error {
	return userdetailsService.userdetailsDao.DeleteUserdetails(id)
}

func (userdetailsService *UserdetailsService) ListUserdetails() ([]*models.Userdetails, error) {
	return userdetailsService.userdetailsDao.ListUserdetails()
}

func (userdetailsService *UserdetailsService) GetUserdetails(id int64) (*models.Userdetails, error) {
	return userdetailsService.userdetailsDao.GetUserdetails(id)
}
