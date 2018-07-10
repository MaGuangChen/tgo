package dboperator

import (
	"database/sql"
	"fmt"

	"github.com/G-Cool-ThanosGo/app"
)

// ScanAndGetResult : 下 raw sql 時, 公用的返回func
func ScanAndGetResult(rawDB *sql.DB, sqlSynx string) RecordType {
	rows, queryErr := rawDB.Query(sqlSynx)
	app.CheckError(queryErr)

	cols, _ := rows.Columns()
	values := make([]sql.RawBytes, len(cols))
	scans := make([]interface{}, len(cols))
	results := make(RecordType)
	for i := range values {
		scans[i] = &values[i]
	}

	defer rows.Close()
	i := 0
	for rows.Next() {
		if queryErr := rows.Scan(scans...); queryErr != nil {
			fmt.Println("Error")
		}
		row := make(map[string]string)

		for j, v := range values {
			key := cols[j]
			row[key] = string(v)
		}

		results[i] = row
		i++
	}

	return results
}
