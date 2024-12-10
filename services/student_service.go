package services

import (
	"gorm.io/gorm"
	"gorm_api/models"
)

type StudentService struct {
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{db: db}
}

// fetch for all students
func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	result := s.db.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil
}

// Get student by id
func (s *StudentService) GetStudentById(id uint) (*models.Student, error) {
	var student models.Student
	result := s.db.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, result.Error
}

// add a student to the databse
func (s *StudentService) CreateStudent(student *models.Student) (*models.Student, error) {
	result := s.db.Create(student)
	return student, result.Error
}

// update values in the students table
func (s *StudentService) UpdateStudent(id uint, updatedData *models.Student) (*models.Student, error) {
	var student models.Student
	result := s.db.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}

	student.Name = updatedData.Name
	student.Age = updatedData.Age
	student.Grade = updatedData.Grade
	result = s.db.Save(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil

}

// Delete student from the database table
func (s *StudentService) DeleteStudent(id uint) error {
	result := s.db.Delete(&models.Student{}, id)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
