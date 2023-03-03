package buah

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Buah, error)
	FindByID(ID int) (Buah, error)
	Create(buah Buah) (Buah, error)
	Update(buah Buah) (Buah, error)
	Delete(buah Buah) (Buah, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Buah, error) {
	var buahs []Buah

	err := r.db.Find(&buahs).Error

	return buahs, err
}

func (r *repository) FindByID(ID int) (Buah, error) {
	var buah Buah

	err := r.db.Find(&buah, ID).Error

	return buah, err
}

func (r *repository) Create(buah Buah) (Buah, error) {
	err := r.db.Create(&buah).Error

	return buah, err
}

func (r *repository) Update(buah Buah) (Buah, error) {
	err := r.db.Save(&buah).Error
	return buah, err
}

func (r *repository) Delete(buah Buah) (Buah, error) {
	err := r.db.Delete(&buah).Error
	return buah, err
}
