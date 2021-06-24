package blogmysql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"database/sql"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "192.168.10.27"
	port     = "32806"
	dbName   = "mysql"
)

var ctx = context.Background()

// MysqlDB 定义的mysql结构体
type MysqlDB struct {
	// 基础配置
	UserName string
	Password string
	IP       string
	Port     string
	DBName   string

	// 性能配置
	MaxLifetime time.Duration
	MaxIdleCons int

	DB *sql.DB
}

func init() {
	var defaultSQL = strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultSQL, 30)
}

// InitDB 用于连接数据库
func (db *MysqlDB) InitDB() error {
	if db == nil {
		return fmt.Errorf("Input param is error")
	}
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{db.UserName, ":", db.Password, "@tcp(", db.IP, ":", db.Port, ")/", db.DBName, "?charset=utf8"}, "")
	DB, err := sql.Open("mysql", path)
	if err != nil {
		return fmt.Errorf("error when Open db:%v", err)
	}

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(db.MaxLifetime)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(db.MaxIdleCons)
	db.DB = DB
	return nil
}

// CloseDB 用于关闭数据库
func (db *MysqlDB) CloseDB() error {
	if db == nil || db.DB == nil {
		return fmt.Errorf("Input param is error")
	}

	return db.DB.Close()
}

// CreateDB 创建数据库
func (db *MysqlDB) CreateDB(dbname string) error {
	if db == nil || db.DB == nil {
		return fmt.Errorf("Input param is error")
	}

	sql := "create database if not exists " + dbname
	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDB 删除数据库
func (db *MysqlDB) DeleteDB(dbname string) error {
	if db == nil || db.DB == nil {
		return fmt.Errorf("Input param is error")
	}

	sql := "drop database" + dbname
	_, err := db.DB.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// ExecWriteSQLInDB 用户在指定的数据库中执行写入操作的sql
func (db *MysqlDB) ExecWriteSQLInDB(config interface{}) error {
	database, err := checkParam(db, config)
	if err != nil {
		return err
	}

	// 开启事务
	tx, err := db.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("Begin tx failed")
	}

	// 首先使用当前的数据库
	_, err = tx.Exec("use " + database.dbName + ";")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Exec sql use failed:%v", err)
	}
	// 在当前的数据库执行sql语句
	for _, v := range database.sqls {
		_, err = tx.Exec(v)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("Exec sql create failed:%v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Tx commit failed")
	}

	return nil
}
