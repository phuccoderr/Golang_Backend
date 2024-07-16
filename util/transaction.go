package util

import "gorm.io/gorm"

func ExcuteInTransaction(db *gorm.DB, fn func(db *gorm.DB) error) error {
	return db.Transaction(fn)
}
