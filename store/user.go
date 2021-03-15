package store

import (
	"github.com/ahmedashrafdev/reports/db"
	"github.com/ahmedashrafdev/reports/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
func (us *UserStore) GetByIDSec(id uint64) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByEmail(e string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Update(u).Error
}

func (us *UserStore) GetServer(id uint) (model.Server, error) {
	var m model.Server
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return m, nil
		}
		return m, err
	}
	return m, nil
}

func (us *UserStore) ConnectDb(id uint) error {
	u, err := us.GetByID(id)

	m, err := us.GetServer(u.ServerID)
	if err != nil {
		return err
	}
	err = db.InitDatabase(m)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) ListUsers() ([]model.User, int, error) {
	var (
		users []model.User
		count int
	)

	us.db.Model(&users).Count(&count)
	us.db.Order("created_at desc").Find(&users)

	return users, count, nil
}
func (us *UserStore) DeleteUser(u *model.User) error {
	return us.db.Delete(u).Error
}
