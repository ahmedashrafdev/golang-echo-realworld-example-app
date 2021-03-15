package store

import (
	"github.com/ahmedashrafdev/reports/model"
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

func (ss *ServerStore) GetByID(id uint64) (*model.Server, error) {
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

func (ss *ServerStore) Update(s *model.Server) error {
	return ss.db.Model(s).Update(s).Error
}

func (ss *ServerStore) ListServers() ([]model.Server, int, error) {
	var (
		servers []model.Server
		count   int
	)

	ss.db.Model(&servers).Count(&count)
	ss.db.Order("created_at desc").Find(&servers)

	return servers, count, nil
}
func (ss *ServerStore) DeleteServer(s *model.Server) error {
	return ss.db.Delete(s).Error
}
