package cv

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Resume, error)
	GetByID(ID int) (Resume, error)
	Save(resume Resume) (Resume, error)
	Update(resume Resume) (Resume, error)
	Delete(resume Resume) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Resume, error) {
	var resume []Resume
	err := r.db.Find(&resume).Error
	if err != nil {
		return resume, err
	}

	return resume, nil
}

func (r *repository) GetByID(ID int) (Resume, error) {
	var resume Resume

	err := r.db.Where("id = ?", ID).Find(&resume).Error
	if err != nil {
		return resume, err
	}

	return resume, nil
}

func (r *repository) Save(resume Resume) (Resume, error) {
	err := r.db.Create(&resume).Error
	if err != nil {
		return resume, err
	}

	return resume, nil
}

func (r *repository) Update(resume Resume) (Resume, error) {
	err := r.db.Save(&resume).Error
	if err != nil {
		return resume, err
	}

	return resume, nil
}

func (r *repository) Delete(resume Resume) error {
	err := r.db.Delete(&resume).Error
	if err != nil {
		return err
	}

	return nil
}
