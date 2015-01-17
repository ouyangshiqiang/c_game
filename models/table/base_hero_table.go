package table

import (
	"github.com/fhbzyc/c_game/libs/db"
)

func init() {
	db.DataBase.Sync(new(BaseHeroTable))
}

type BaseHeroTable struct {
	HeroId int    `xorm:"pk autoincr"`
	Name   string `xorm:"'hero_name' NOT NULL DEFAULT ''"`
	Star   int    `xorm:"'hero_star' NOT NULL DEFAULT 0 TINYINT(1)"`
	Skill1 int    `xorm:"'skill_id1' NOT NULL DEFAULT 0"`
	Skill2 int    `xorm:"'skill_id2' NOT NULL DEFAULT 0"`
	Skill3 int    `xorm:"'skill_id3' NOT NULL DEFAULT 0"`
	Skill4 int    `xorm:"'skill_id4' NOT NULL DEFAULT 0"`
}

func (this BaseHeroTable) TableName() string {
	return "base_hero"
}
