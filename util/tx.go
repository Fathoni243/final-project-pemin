package util

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) error {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		return errorRollback
	} else {
		errorCommit := tx.Commit().Error
		return errorCommit
	}
}
