package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error) //Save nama method ,user nama package,User nama model/entity struct,(User, error) nilai yg harus direturn
}

// tidak bisa diakses dari luar /internal struct
type repository struct {
	db *gorm.DB //panggil database
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
