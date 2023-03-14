package router

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
