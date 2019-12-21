package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	c.String(http.StatusOK, "GetUser")
}

func Create(c *gin.Context) {
	c.String(http.StatusOK, "CreateUser")
}

func Update(c *gin.Context) {
	c.String(http.StatusOK, "GetUser")
}

func Delete(c *gin.Context) {
	c.String(http.StatusOK, "CreateUser")
}

func Search(c *gin.Context) {
	c.String(http.StatusOK, "GetUser")
}
