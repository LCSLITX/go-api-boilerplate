package models

import (
	"time"
	"context"

	"github.com/lucassauro/go-api-boilerplate/db"
	"github.com/lucassauro/go-api-boilerplate/utils"
	"github.com/lucassauro/go-api-boilerplate/structs"
)

// This function query db.table and return only the first correspondence.
func TestDB() structs.DBItem {
	db := db.MySQL()
	defer db.Close()

	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()

	var r structs.DBItem

	err := db.QueryRowContext(ctx, "SELECT * FROM `tableA`;").Scan(&r.ColumnA, &r.ColumnB)
	utils.Check(err)

	return r
}

// This function query db.table and return all correspondences.
func TestsDB() (result []structs.DBItem) {
	db := db.MySQL()
	defer db.Close()

	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()

	result = make([]structs.DBItem, 0)

	s, err := db.QueryContext(ctx, "SELECT * FROM `tableA`;")
	utils.Check(err)
	defer s.Close()

	for s.Next() {
		var n structs.DBItem
		err := s.Scan(&n.ColumnA, &n.ColumnB)
		utils.Check(err)
		result = append(result, n)
	}
	err = s.Err()
	utils.Check(err)

	return result
}
