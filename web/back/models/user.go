package models

import (
	"back/pkg/crypto"
)

type Message struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique;not null;type:varchar(12)"`
	Password string `json:"password" gorm:"not null;type:varchar(128)"`
	Level    int    `json:"level" gorm:"not null;type:tinyint"`
	Name     string `json:"name" gorm:"not null;type:varchar(20)"`
}

// AddUser 完整的用户信息进行注册
func AddUser(user *Message) (err error) {
	if err = db.Omit("id").Create(user).Error; err != nil {
		return err
	}
	return nil
}

// CheckUser  通过username，password 检查用户是否存在及密码是否正确
func CheckUser(username, password string) uint {
	var message Message
	db.Select("id").Where(Message{Username: username, Password: password}).First(&message)
	return message.ID
}

// GetUser 通过username获取用户权限等级
func GetUser(username string) (msg Message) {
	db.Debug().Select("level").Where("username = ? ", username).First(&msg)
	return
}

// GetUserInfo id获取用户信息
func GetUserInfo(id uint) (Message, error) {
	var msg Message
	if err := db.Select("id,username,level,name").Where("id = ? ", id).First(&msg).Error; err != nil {
		return Message{}, err
	}
	return msg, nil
}

// AddMessage 加密密码，并保存用户信息
func AddMessage(message Message) error {
	message.Password = crypto.Encrypt(message.Password)
	if err := db.Omit("id").Create(&message).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 传入id删除，若不存在用户则返回错误
func DeleteUser(id uint) error {
	if err := db.Where("id = ?", id).Delete(&Message{}).Error; err != nil {
		return err
	}
	return nil
}

// GetUsers 获取所有权限小于level的用户
func GetUsers(level int) (users []Message) {
	db.Where("level > ?", level).Find(&users)
	return
}

// UpdateUser 传递修改后的信息，ID不存在、或格式不正确则返回错误
func UpdateUser(message Message) error {
	// 修改后信息进行加密
	if err := db.Debug().Model(&Message{}).Where("id = ?", message.ID).Omit("id", "username", "password").Updates(message).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePassword(password string, id uint) error {
	if err := db.Debug().Model(&Message{ID: id}).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
