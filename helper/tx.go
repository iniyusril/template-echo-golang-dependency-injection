package helper

import (
	"github.com/jinzhu/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Rollback()

	} else {
		tx.Commit()
	}
}
