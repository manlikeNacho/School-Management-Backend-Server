package main

import (
	"TheAdmin/src/router"
	"TheAdmin/src/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	setup.Connection()
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "welcome to deez administration",
		})
	})
	r.POST("/createStudent", router.CreateStudent)
	r.GET("/getStudentById/:id", router.GetStudentById)
	r.GET("/listStudents", router.ListStudents)
	r.PUT("/updateStudentById/:id", router.DeleteStudent)
	r.DELETE("/deleteStudentById/:id", router.DeleteStudent)
	r.Run()
}
