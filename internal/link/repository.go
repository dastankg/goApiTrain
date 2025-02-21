package link

import (
	"apiProject/pkg/db"
)

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

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	link := &Link{}
	res := repo.Database.DB.Where("hash = ?", hash).First(link)

	if res.Error != nil {
		return nil, res.Error
	}

	return link, nil
}
func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	link := &Link{}
	res := repo.Database.DB.Where("id = ?", id).First(link)
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
		return res.Error
	}
	return nil
}

func (repo *LinkRepository) GetLinks(limit, offset int) []Link {
	links := []Link{}

	repo.Database.Table("links").
		Where("deleted_at is null").Limit(limit).Offset(offset).Scan(&links)
	return links
}

func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.Database.Table("links").Where("deleted_at is null").Count(&count)
	return count
}
