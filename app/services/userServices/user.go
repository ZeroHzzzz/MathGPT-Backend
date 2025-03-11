package userservices

import (
	"mathgpt/app/apiException"
	"mathgpt/app/models"
	"mathgpt/configs/database/mysql"
)

func CreateUser(email, phone, password string) (*models.User, error) {
	user := &models.User{
		Email:    email,
		Phone:    phone,
		Password: password,
	}

	result := mysql.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByPhone 根据手机号查询用户
func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	result := mysql.DB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := mysql.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID 根据用户 ID 查询用户
func GetUserByID(userID string) (*models.User, error) {
	var user models.User
	result := mysql.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByIDAndPass(userID, password string) (*models.User, error) {
	var user models.User
	result := mysql.DB.Where("id = ? and password = ?", userID, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByEmailAndPass(email, password string) (*models.User, error) {
	var user models.User
	result := mysql.DB.Where("email = ? and password = ?", email, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByPhoneAndPass(phone, password string) (*models.User, error) {
	var user models.User
	result := mysql.DB.Where("phone = ? and password = ?", phone, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(userID string, updates map[string]interface{}) error {
	result := mysql.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return apiException.UserNotFind
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(userID string) error {
	result := mysql.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
