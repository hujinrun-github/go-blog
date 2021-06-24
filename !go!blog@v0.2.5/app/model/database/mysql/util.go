package blogmysql

import "fmt"

func checkParam(db *MysqlDB, config interface{}) (DataBase, error) {
	database := DataBase{}

	if db == nil || db.DB == nil {
		return database, fmt.Errorf("Input param error")
	}

	if tmpDatabase, ok := config.(DataBase); ok {
		database = tmpDatabase
	} else {
		return database, fmt.Errorf("Input param error")
	}
	return database, nil
}
