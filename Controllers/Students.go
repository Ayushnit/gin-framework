package Controllers

import (
	"fmt"
	"gin-framework/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)
func GetStudents(c *gin.Context) {
	var st []Models.Student
	err := Models.GetAllStudents(&st)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, st)
	}
}
func CreateStudent(c *gin.Context) {
	var st Models.Student
	c.BindJSON(&st)
	err := Models.CreateStudent(&st)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, st)
	}
}
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var st Models.Student
	err := Models.GetStudentByID(&st, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, st)
	}
}
func UpdateStudent(c *gin.Context) {
	var st Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&st, id)
	if err != nil {
		c.JSON(http.StatusNotFound, st)
	}
	c.BindJSON(&st)
	err = Models.UpdateStudent(&st, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, st)
	}
}
func DeleteStudent(c *gin.Context) {
	var st Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&st, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
