package controller

import (
	"Bookee/controller/comm"
	"Bookee/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type bookController struct {
	bookService service.BookService
}

func RegisterBookController(bookGroup *gin.RouterGroup) (err error) {
	if nil == bookGroup {
		panic("Router for Book should not be nil")
	}

	controller := bookController{service.BookSvc()}

	bookGroup.GET(`/search`, controller.search)
	bookGroup.GET(`/detail`, controller.detail)

	return
}

func (this *bookController) search(c *gin.Context) {
	cond := struct {
		comm.PageRequest
		BookName string `json:"bookname" form:"bookname"`
	}{}
	c.ShouldBind(&cond)
	c.JSON(http.StatusOK, comm.Response{Data: cond})
}

func (this *bookController) detail(c *gin.Context) {
	bookId := struct {
		BookId int64 `form:"bookId" binding:"required"`
	}{}
	err := c.Bind(&bookId)

	if err == nil {
		book := this.bookService.Get(bookId.BookId)
		c.JSON(http.StatusOK, comm.Response{Data: book})
	} else {
		c.JSON(http.StatusBadRequest, comm.Response{Status: comm.ResponseStatus{Code: -1, Msg: err.Error()}})
	}
}
