package main

import (
	"TheAdmin/src/controller"
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
	r.POST("/student", controller.CreateStudent)
	r.GET("/student/:id", controller.GetStudentById)
	r.PUT("/student/:id", controller.UpdateStudent)
	r.DELETE("student/:id", controller.DeleteStudent)
	r.GET("/list/student", controller.ListStudents)
	r.GET("/list/course", controller.ListCourses)
	r.GET("/list/student/:course_name", controller.ListStudentsByCourse)
	r.Run()
}
