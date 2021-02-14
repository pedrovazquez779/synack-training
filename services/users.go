package services

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"synack/models"
)

type IUserService interface {
	GetUserByID(userID int) (*models.User, error)
}

type UserService struct{}

func (uc *UserService) GetUserByID(userID int) (*models.User, error) {

	dbService := DBService{}
	db, err := dbService.GetDB()
	if err != nil {
		return nil, err
	}

	foundUser, err := models.FindUser(context.TODO(), db, userID)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uc *UserService) SaveUser(user *models.User) (*models.User, error) {

	dbService := DBService{}
	db, err := dbService.GetDB()
	if err != nil {
		return nil, err
	}

	err = user.Insert(context.TODO(), db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserService) UpdateUser(userID int, user *models.User) (*models.User, error) {

	dbService := DBService{}
	db, err := dbService.GetDB()
	if err != nil {
		return nil, err
	}

	foundUser, err := models.FindUser(context.TODO(), db, userID)
	if err != nil {
		return nil, err
	}

	foundUser.Name = user.Name
	_, err = foundUser.Update(context.TODO(), db, boil.Whitelist(models.UserColumns.Name))
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

func (uc *UserService) DeleteUser(userID int) (*models.User, error) {

	dbService := DBService{}
	db, err := dbService.GetDB()
	if err != nil {
		return nil, err
	}

	deleteUser, err := models.FindUser(context.TODO(), db, userID)
	if err != nil {
		return nil, err
	}

	_, err = deleteUser.Delete(context.TODO(), db)
	if err != nil {
		return nil, err
	}

	return deleteUser, nil
}
