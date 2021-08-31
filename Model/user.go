package Model

import (
	emsg "Exe/Utils/ErrorMessage"
	"Exe/Utils/GenerateId"
	"Exe/Utils/SensitiveWordFilter"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var IM = gid.NewIdMaker()

type User struct {
	Id       int    `json:"id" gorm:"type:int; primary_key"`
	UserName string `json:"username" gorm:"type:varchar(25)" binding:"required,valid"`
	PassWord string `json:"password" gorm:"type:varchar(25)" binding:"required,min=8,max=20"`
}

//判断用户名是否合法
func valid(fl validator.FieldLevel) bool {
	if _, ok := SensitiveWordFilter.T.Check(fl.Field().Interface().(string)); ok {
		return false
	}
	return true
}

func init() {
	//注册表单参数验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("valid", valid)
		if err != nil {
			log.Println(err, "表单参数验证注册失败")
		}
	}
}

//用户注册
func Register(username, password string) (*User, int) {
	var id int
	DB.Model(User{}).Select("id").Where("user_name = ?", username).First(&id)
	if id == 0 {
		//该用户名可用
		user := &User{
			Id:       IM.NewId(),
			UserName: username,
			PassWord: password,
		}
		DB.Create(user)
		return user, emsg.Success
	}
	return nil, emsg.UsernameExist
}

//用户登录
func Login(username, password string) int {
	var pw string
	DB.Model(User{}).Select("pass_word").Where("user_name = ?", username).First(&pw)
	//判断用户是否存在
	if pw == "" {
		return emsg.UsernameNoExist
	} else if pw == password { //判断密码是否正确
		return emsg.Success
	}
	return emsg.PasswordWrong
}
