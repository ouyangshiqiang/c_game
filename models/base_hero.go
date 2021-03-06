package models

import (
	"fmt"
	"github.com/fhbzyc/c_game/libs/db"
	"github.com/fhbzyc/c_game/models/table"
)

var BaseHero baseHeroModel

type baseHeroModel struct {
}

func (this baseHeroModel) FindAll() ([]table.BaseHeroTable, error) {
	var list []table.BaseHeroTable
	return list, db.DataBase.Find(&list)
}

func (this baseHeroModel) FindOne(heroId int) table.BaseHeroTable {
	list, _ := this.FindAll()

	for _, item := range list {
		if item.HeroId == heroId {
			return item
		}
	}

	panic(fmt.Sprintf("不存在的英雄: %d", heroId))
}
