package blogmysql

var createAdminTableSQL = `
CREATE TABLE admin (
        id INT(10) NOT NULL AUTO_INCREMENT,
        name VARCHAR(64) NULL DEFAULT NULL,
        password VARCHAR(64) NULL DEFAULT NULL,
		nick VARCHAR(64) NULL DEFAULT NULL,
		email TEXT NULL DEFAULT NULL,
		avator TEXT NULL DEFAULT NULL,
		url TEXT NULL DEFAULT NULL,
		bio TEXT NULL DEFAULT NULL,
		createTime DATETIME,
		lastLoginTime DATETIME,
		role TINYTEXT,
		PRIMARY KEY ( id )
	)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

// DataBase 创建数据库索引的结构体
type DataBase struct {
	dbName string
	sqls   []string
}

// InsertStruct 用于向数据表中插入数据的结构体
type InsertStruct struct {
	dbName string
	data   interface{}
}
