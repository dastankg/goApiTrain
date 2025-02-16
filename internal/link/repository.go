package link

import "apiProject/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	res := repo.Database.DB.Create(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) Get(hash string) (*Link, error) {
	link := &Link{}
	res := repo.Database.DB.Where("hash = ?", hash).First(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	res := repo.Database.DB.Updates(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	res := repo.Database.DB.Where("id = ?", id).Delete(&Link{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
