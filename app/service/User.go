package service

import (
	"errors"
	"gin-mall/app/common/request"
	"gin-mall/app/dao"
	"gin-mall/app/model"
	"gin-mall/global"
	"gin-mall/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
)

/**
 * @author: biao
 * @date: 2025/9/2 21:36
 * @code: 彼方尚有荣光在
 * @description: User处理服务
 */

type userService struct {
}

var UserService = new(userService)

func (u *userService) UserRegister(params request.UserRegister) (err error, user *model.User) {
	if params.Key == "" || len(params.Key) != 16 {
		return errors.New("Register-密钥长度不足"), user
	}

	// 对key进行加密操作
	utils.Encrypt.SetKey(params.Key)
	//通过用户名查询数据库是否存在（用户名已存在 错误码为30001）
	_, exist, err := dao.UserDao.FindUserByUserName(params.UserName)
	if err != nil {
		global.App.Log.Error("Register-用户信息查询失败：", zap.Any("err", err))
		return errors.New("Register-用户信息查询失败"), user
	}
	if exist {
		global.App.Log.Error("Register-用户已存在：", zap.String("UserName", params.UserName))
		return errors.New("Register-用户名已存在"), user
	}
	//创建用户（金钱默认10000，并且加密）
	user = &model.User{
		UserName:       params.UserName,
		NickName:       params.NickName,
		PasswordDigest: utils.BcryptMake([]byte(params.Password)), // 密码加密
		Avatar:         "avatar.jpg",
		Status:         model.Active,
		Money:          utils.Encrypt.AesEncoding("10000"), //初始金额加密
	}

	//添加用户
	err = dao.UserDao.CreateUser(user)
	return
}

func (u *userService) UserLogin(params request.UserLogin) (err error, user *model.User) {
	// 查看对应的用户名是否存在
	user, exist, err := dao.UserDao.FindUserByUserName(params.UserName)

	if err != nil {
		global.App.Log.Error("Login-用户信息查询失败", zap.Any("err", err))
		return
	}

	if !exist {
		global.App.Log.Error("Login-用户不存在", zap.String("user_name", params.UserName))
		err = errors.New("Login-用户不存在，请先注册")
		return
	}

	global.App.Log.Info("用户信息查询成功", zap.String("user_name", params.UserName))

	// 验证密码
	if !utils.BcryptMakeCheck([]byte(params.Password), user.PasswordDigest) {
		global.App.Log.Error("密码验证失败", zap.String("user_name", params.UserName))
		err = errors.New("密码验证失败，请重试")
		return
	}
	global.App.Log.Info("密码验证成功", zap.String("user_name", params.UserName))
	return
}

func (u *userService) UserUpdateInfo(params request.UserUpdateInfo, id uint) (err error, user *model.User) {
	// 根据id查找对应用户
	user, err = dao.UserDao.GetUserById(id)
	if err != nil {
		global.App.Log.Error("根据id查询用户信息失败", zap.Any("err", err))
		return
	}

	if params.Email != "" {
		user.Email = params.Email
	}
	if params.UserName != "" {
		user.UserName = params.UserName
	}
	if params.Password != "" {
		user.PasswordDigest = utils.BcryptMake([]byte(params.Password))
	}
	if params.NickName != "" {
		user.NickName = params.NickName
	}

	// 用户相关信息
	err = dao.UserDao.UpdateUserInfo(user)
	if err != nil {
		global.App.Log.Error("更新用户信息失败")
	}
	return
}

// 上传头像
func (u *userService) UploadAvatarToLocal(id uint, file multipart.File) (err error, filePath string) {
	// 根据id查找用户
	user, err := dao.UserDao.GetUserById(id)
	if err != nil {
		global.App.Log.Error("根据id查询用户信息失败", zap.Any("err", err))
		return
	}
	// 上传图片至本地
	err, filePath = UploadService.UploadAvatarToLocalStatic(id, file)
	if err != nil {
		global.App.Log.Error("用户上传图片至本地失败", zap.Any("err", err))
		return
	}

	// 修改数据库信息
	user.Avatar = filePath
	err = dao.UserDao.UpdateUserInfo(user)
	if err != nil {
		global.App.Log.Error("更新用户信息失败", zap.Any("err", err))
		return err, ""
	}
	return nil, filePath
}

// 发送邮件
func (u *userService) SendEmail(id uint, params request.SendEmail) (err error) {

	tokenData, err, _ := JwtService.CreateEamilToken(id, params.Email, params.Password, params.OperationType, AppGuardName)
	if err != nil {
		global.App.Log.Error("签发EmailToken失败", zap.Any("err", err))
		return
	}

	err, notice := dao.NoticeDao.GetNoticeById(params.OperationType)
	if err != nil {
		global.App.Log.Error("获取Notice数据失败", zap.Any("err", err))
		return
	}

	address := global.App.Config.Email.ValidEmail + tokenData.AccessToken //发送方地址
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)
	// 信（头部：发送、接收地址，标题，体：信的内容）
	m := mail.NewMessage()
	m.SetHeader("From", global.App.Config.Email.SmtpEmail)
	m.SetHeader("To", params.Email)
	m.SetHeader("Subject", "Gin-Mall")
	m.SetBody("text/html", mailText)
	// 邮件发送工具 （使用的哪个邮局， 邮局地址， 寄件人， 寄件人的授权码）
	d := mail.NewDialer(global.App.Config.Email.SmtpHost, 465, global.App.Config.Email.SmtpEmail, global.App.Config.Email.SmtPass)
	// 启用邮件加密（TLS加密）
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// 通过发送工具连接到邮局，再将信发送出去
	if err = d.DialAndSend(m); err != nil {
		global.App.Log.Error("邮件发送失败", zap.Any("err", err))
		return
	}

	return nil
}

func (u *userService) ValidEmail(c *gin.Context) (err error, user *model.User) {
	id, _ := c.Get("id")
	password, _ := c.Get("password")
	email, _ := c.Get("email")
	operationType, _ := c.Get("operationType")

	user, err = dao.UserDao.GetUserById(id.(uint))
	if err != nil {
		global.App.Log.Error("获取用户信息失败", zap.Any("err", err))
		return err, nil
	}

	switch operationType.(uint) {
	case 1: // 绑定邮箱
		user.Email = email.(string)
		break
	case 2: // 解绑邮箱
		user.Email = ""
		break
	case 3: // 修改密码
		user.PasswordDigest = utils.BcryptMake([]byte(password.(string)))
		break
	}

	err = dao.UserDao.UpdateUserInfo(user)
	if err != nil {
		global.App.Log.Error("更新用户信息失败", zap.Any("err", err))
		return err, nil
	}

	return nil, user
}
