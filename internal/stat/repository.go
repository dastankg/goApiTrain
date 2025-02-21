package stat

import (
	"apiProject/pkg/db"
	"gorm.io/datatypes"
	"time"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{db}
}

func (repo *StatRepository) AddClick(linkId uint) {

	stat := Stat{}
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, datatypes.Date(time.Now()))
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   datatypes.Date(time.Now()),
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string
	if by == "day" {
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	} else {
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	repo.Db.Table("stats").Select(selectQuery).Where("date BETWEEN ? AND ?", from, to).
		Group("period").Order("period").Scan(&stats)

	return stats
}
