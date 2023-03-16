package controller

import (
	"TheAdmin/src/models"
	"TheAdmin/src/setup"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.AbortWithStatus(404)
		return
	}
	students := models.Student{
		Name:         student.Name,
		Email:        student.Email,
		CoursesTaken: student.CoursesTaken,
	}
	res := setup.DB.Create(&students)

	if res.Error != nil {
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	result := setup.DB.First(&student, id)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func ListStudents(c *gin.Context) {
	var students []models.Student
	result := setup.DB.Find(&students)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var response []models.StudentResponse
	for _, s := range students {
		response = append(response, models.StudentResponse{
			ID:    s.ID,
			Name:  s.Name,
			Email: s.Email,
		})
	}
	c.JSON(http.StatusOK, response)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	result := setup.DB.First(&student, id)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var input models.Student
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	student.Name = input.Name
	student.Email = input.Email
	result = setup.DB.Save(&student)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	result := setup.DB.First(&student, id)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	result = setup.DB.Delete(&student)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func ListCourses(c *gin.Context) {
	var courses []models.Course
	result := setup.DB.Find(&courses)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var response []models.CourseResponse
	for _, c := range courses {
		response = append(response, models.CourseResponse{
			ID:    c.ID,
			Title: c.Title,
		})
	}
	c.JSON(http.StatusOK, response)
}

func ListStudentsByCourse(c *gin.Context) {
	courseName := c.Param("course_name")
	var students []string
	setup.DB.Table("students").Joins("JOIN student_courses ON student_courses.student_id = students.id").Joins("JOIN courses ON student_courses.course_id = courses.id").Where("courses.title = ?", courseName).Pluck("students.name", &students)

	c.JSON(http.StatusOK, gin.H{"students": students})
}
