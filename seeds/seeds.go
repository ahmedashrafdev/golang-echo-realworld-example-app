package seeds

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/seed"
	"github.com/jinzhu/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateServer",
			Run: func(db *gorm.DB) error {
				return CreateServer(db)
			},
		},
		seed.Seed{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				return CreateUser(db)
			},
		},
	}
}
