package store

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
)

type ServerStore struct {
	db *gorm.DB
}

func NewServerStore(db *gorm.DB) *ServerStore {
	return &ServerStore{
		db: db,
	}
}

func (ss *ServerStore) GetByID(id uint) (*model.Server, error) {
	var m model.Server
	if err := ss.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (ss *ServerStore) Create(u *model.Server) (err error) {
	return ss.db.Create(u).Error
}

func (ss *ServerStore) Update(u *model.Server) error {
	return ss.db.Model(u).Update(u).Error
}
