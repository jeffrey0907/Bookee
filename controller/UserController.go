package controller

import (
    "Bookee/service"
    "math"
    "net/http"

    "Bookee/controller/comm"
    "github.com/gin-gonic/gin"
)

type userController struct {
    userService service.UserService
}

func RegistUserController(userGroup *gin.RouterGroup) (err error) {
    ctrl := userController{userService: service.UserSvc()}

    userGroup.POST("/login", ctrl.login)
    userGroup.POST("/login/wx", ctrl.loginWX)
    userGroup.POST("/register", ctrl.register)

    userGroup.GET("/books", service.SessionSvc().CheckJWT, ctrl.listBooks)
    return
}

func (uc *userController) register(c *gin.Context) {

}

func (uc *userController) loginWX(c *gin.Context) {
    reqParams := struct {
        Code string `json:"code"`
    }{}
    err := c.BindJSON(&reqParams)
    if err == nil {
        token, err := uc.userService.LoginWX(reqParams.Code)
        if err == nil {
            c.JSON(http.StatusOK, comm.ResponseOk(token))
        } else {
            c.JSON(http.StatusOK, comm.ResponseErr(nil, -1, err.Error()))
        }
    } else {
        c.JSON(http.StatusOK, comm.ResponseErr(nil, -1, "Code不能为空"))
    }
}

func (uc *userController) login(c *gin.Context) {
    reqParams := struct {
        Uid int64  `json:"uid" form:"uid" binding:"required"`
        Pwd string `json:"password" form:"password" binding:"required"`
    }{}
    err := c.BindJSON(&reqParams)

    if err == nil {
        token, err := uc.userService.Login(reqParams.Uid, reqParams.Pwd)
        if err == nil {
            c.SetCookie("accessToken", token, math.MaxInt32, "", "", false, true)
            user := uc.userService.Get(reqParams.Uid)
            c.JSON(http.StatusOK, comm.ResponseOk(user))
        } else {
            c.JSON(http.StatusOK, comm.ResponseErr(nil, 1, err.Error()))
        }
    } else {
        c.JSON(http.StatusOK, comm.ResponseErr(nil, -1, "用户名或密码不能为空"))
    }
}

func (uc *userController) listBooks(c *gin.Context) {

}
