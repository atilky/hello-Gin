package models

import "gindemo02/util"

type User struct {
	ID       int    `gorm:"column:id;primaryKey"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  int64  `json:"add_time"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) GetUsersByAccAndPwd(account, pwd string) (users User, err error) {
	err = DB.Select("*").
		Where("account = ?", account).Where("password = ?", pwd).
		First(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

func (u User) GetUsers() (users []User, err error) {
	err = DB.Select("*").
		Where("age > ?", u.Age).
		Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

// 创建一个用户
func (u User) CreateUser(user User) {
	//ORM
	if err := DB.Create(&user).Error; err != nil { //必须传指针，因为要给user的主键赋值
		util.LogRus.Errorf("create user %s failed: %s", user.Name, err)
	} else {
		util.LogRus.Infof("create user id %d", user.ID)
	}
}

func (u User) DeleteUserById(id int) (err error) {
	//ORM
	if err := DB.Where("id=?", id).Delete(User{}).Error; err != nil { //Delete操作必须有where条件
		util.LogRus.Errorf("delete user %d failed: %s", id, err)
		return err
	}
	return nil
}
