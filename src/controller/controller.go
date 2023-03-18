package controller

import (
	"TheAdmin/src/models"
	"TheAdmin/src/setup"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateStudent godoc
// @Summary      Create Student Account
// @Description  create student account
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  models.StudentResponse{}
// @Router        /student [post]
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

// GetStudentById godoc
// @Summary      Get Student
// @Description  get student by Id
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  models.StudentResponse{}
// @Router        /student/:id [get]
// @Param        id   path    int  true  "Student ID"
func GetStudentById(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := setup.DB.Where("id = ?", id).Preload("CoursesTaken").First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// ListStudents godoc
// @Summary      List Students
// @Description  List total students in the db
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  []models.StudentResponse{}
// @Router        /list/student [get]
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

// UpdateStudent godoc
// @Summary      Update Student account
// @Description  Update student Name/email
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  models.StudentResponse{}
// @Router        /student/:id [put]
// @Param        id   path    int  true  "Student ID"
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

// DeleteStudent godoc
// @Summary      Delete Student Account
// @Description  Delete student account
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  models.StudentResponse{}
// @Router        /student/:id [delete]
// @Param        id   path    int  true  "Student ID"
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

// ListCourses godoc
// @Summary      List courses
// @Description  List courses total courses taken by all students
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  models.CourseResponse{}
// @Router        /list/course [get]
// @Param        course_name   path    string  true  "Student ID"
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

// ListStudentsByCourse godoc
// @Summary      Create Student Account
// @Description  create student account
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Success       200 {object}  []string
// @Router        /list/student/:course_name [get]
// @Param        course_name   path    string  true  "Student ID"
func ListStudentsByCourse(c *gin.Context) {
	courseName := c.Param("course_name")
	var students []string
	setup.DB.Table("students").Joins("JOIN student_courses ON student_courses.student_id = students.id").Joins("JOIN courses ON student_courses.course_id = courses.id").Where("courses.title = ?", courseName).Pluck("students.name", &students)

	c.JSON(http.StatusOK, gin.H{"students": students})
}

// UpdateStudentCourseById godoc
// @Summary      Create Student Account
// @Description  create student account
// @Tags         Student
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      int  true  "Student ID"
// @Success       200 {object} models.StudentResponse{}
// @Router        /course/:id/student [put]
func UpdateStudentCourseById(c *gin.Context) {
	studentID := c.Param("id")
	var req models.UpdateCoursesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	if err := setup.DB.Where("id = ?", studentID).Preload("CoursesTaken").First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
		return
	}

	var courses []models.Course
	if err2 := setup.DB.Where("id IN (?)", req.CourseIds).Find(&courses).Error; err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No course models"})
	}

	if student.CoursesTaken == nil {
		student.CoursesTaken = []models.Course{}
	}

	student.CoursesTaken = courses
	if err := setup.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "courses updated successfully"})
}
