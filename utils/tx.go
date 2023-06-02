package utils

import "gorm.io/gorm"

func CommitOrRollBack(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
		PanicIfError(err.(error))
	} else {
		tx.Commit()
	}
}
