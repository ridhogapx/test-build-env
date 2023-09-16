package controller

import (
	"test-build-env/config"
	"test-build-env/model"

	"github.com/gin-gonic/gin"
)

func SaveBook(ctx *gin.Context) {
  var book model.SaveBookRequest
  
  ctx.ShouldBindJSON(&book)

  newData := model.Book{
    Title: book.Title,
    Author: book.Author,
  }

  config.DB.Create(&newData)

  ctx.JSON(201, model.WebResponse{
    Status: "success",
    Message: "Book was saved!",
  })

}
