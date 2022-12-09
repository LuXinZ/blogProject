package blogApp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBlogs(c *gin.Context) {
	blogs, err := GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": blogs,
	})
}
func GetBlogById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	blog, err := Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": blog,
	})
}
func CreateBlog(c *gin.Context) {
	var blog Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
	}
	err = Insert(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": blog,
	})
}
func DeleteBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}
func UpdateBlog(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
		return
	}
	err = Update(id, &blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})

}
