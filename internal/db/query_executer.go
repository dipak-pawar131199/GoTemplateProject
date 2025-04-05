package db

import (
	"errors"
	"log"
)

// ExecuteQuery runs any SQL query and returns the result as a slice of maps.
func ExecuteQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		log.Println("Query execution failed:", err)
		return nil, err
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Prepare a slice for result storage
	var results []map[string]interface{}

	for rows.Next() {
		// Create a slice to hold column values
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))

		// Assign addresses to column pointers
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Scan row values
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Map column values into a dictionary
		rowData := make(map[string]interface{})
		for i, colName := range columns {
			val := columnValues[i]

			// Convert raw bytes to string (for better JSON compatibility)
			if b, ok := val.([]byte); ok {
				rowData[colName] = string(b)
			} else {
				rowData[colName] = val
			}
		}

		results = append(results, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.New("no records found")
	}

	return results, nil
}
