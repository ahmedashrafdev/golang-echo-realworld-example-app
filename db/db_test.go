package db

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

func TestAutoMigrate(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AutoMigrate(tt.args.db)
		})
	}
}
