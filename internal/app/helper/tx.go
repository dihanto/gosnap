package helper

import (
	"database/sql"
	"log"
)

func CommitOrRollback(tx *sql.Tx, err *error) {
	if *err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			log.Println(errRollback)
		}
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			log.Println(errCommit)
		}
	}
}
