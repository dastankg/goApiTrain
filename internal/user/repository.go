package user

import "apiProject/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	res := repo.Database.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
func (repo *UserRepository) GetByEmail(email string) (*User, error) {
	user := &User{}
	res := repo.Database.Where("email = ?", email).First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
