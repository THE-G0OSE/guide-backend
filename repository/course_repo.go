package repository

import (
	"github.com/THE-G0OSE/guide-backend/models"
	"gorm.io/gorm"
)

type CourseRepo struct{ DB *gorm.DB }

func (r *CourseRepo) Find(id uint) (*models.Course, error) {
	var course models.Course
	err := r.DB.First(&course, id).Error
	return &course, err
}

func (r *CourseRepo) Create(course *models.Course) error {
	err := r.DB.Create(course).Error
	return err
}

func (r *CourseRepo) FindMy(creator uint) (*[]models.Course, error) {
	var courses []models.Course
	err := r.DB.Where("Creator_ID = ?", creator).Find(&courses).Error
	return &courses, err
}

func (r *CourseRepo) Delete(course *models.Course) error {
	return r.DB.Delete(course, "ID = ?", course.ID).Error
}

func (r *CourseRepo) Patch(course *models.Course) error {
	return r.DB.Save(course).Error
}
