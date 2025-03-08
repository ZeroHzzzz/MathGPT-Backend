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

	result := mysql.MysqlDB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByPhone 根据手机号查询用户
func GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	result := mysql.MysqlDB.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := mysql.MysqlDB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID 根据用户 ID 查询用户
func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := mysql.MysqlDB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByIDAndPass(userID, password string) (*models.User, error) {
	var user models.User
	result := mysql.MysqlDB.Where("id = ? and password = ?", userID, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByEmailAndPass(email, password string) (*models.User, error) {
	var user models.User
	result := mysql.MysqlDB.Where("email = ? and password = ?", email, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByPhoneAndPass(phone, password string) (*models.User, error) {
	var user models.User
	result := mysql.MysqlDB.Where("phone = ? and password = ?", phone, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(userID uint, updates map[string]interface{}) error {
	result := mysql.MysqlDB.Model(&models.User{}).Where("id = ?", userID).Updates(updates)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return apiException.UserNotFind
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(userID uint) error {
	result := mysql.MysqlDB.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
