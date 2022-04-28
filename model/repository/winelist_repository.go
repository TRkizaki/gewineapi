package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TRKizaki/gewineapi/model/entity"
)

type WinelistRepository interface {
	GetWinelists() (winelists []entity.WinelistEntity, err error)
	InsertWinelist(winelist entity.WinelistEntity) (id int, err error)
	UpdateWinelist(winelist entity.WinelistEntity) (err error)
	DeleteWinelist(id int) (err error)
}

type winelistRepository struct {
}

func NewWinelistRepository() WinelistRepository {
	return &winelistRepository{}
}

func (wr *winelistRepository) GetWinelists() (winelists []entity.WinelistEntity, err error) {
	winelists = []entity.WinelistEntity{}
	rows, err := Db.
		Query("SELECT id, title, content FROM winelist ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		winelist := entity.WinelistEntity{}
		err = rows.Scan(&winelist.Id, &winelist.Title, &winelist.Brand, &winelist.Price)
		if err != nil {
			log.Print(err)
			return
		}
		winelists = append(winelists, winelist)
	}

	return
}

func (wr *winelistRepository) InsertWinelist(winelist entity.WinelistEntity) (id int, err error) {
	_, err = Db.Exec("INSERT INTO winelist (title, Brand) VALUES (?, ?)", winelist.Title, winelist.Brand, winelist.Price)
	if err != nil {
		log.Print(err)
		return
	}
	err = Db.QueryRow("SELECT id FROM winelist ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}

func (wr *winelistRepository) UpdateWinelist(winelist entity.WinelistEntity) (err error) {
	_, err = Db.Exec("UPDATE winelist SET title = ?, brand = ?, price = ? WHERE id = ?", winelist.Title, winelist.Brand, winelist.Price, winelist.Id)
	return
}

func (wr *winelistRepository) DeleteWinelist(id int) (err error) {
	_, err = Db.Exec("DELETE FROM winelist WHERE id = ?", id)
	return
}