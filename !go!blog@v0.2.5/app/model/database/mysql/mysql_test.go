package blogmysql

import (
	"goblog/app/model"
	"strings"
	"testing"

	"github.com/astaxie/beego/orm"
)

func TestMySql(t *testing.T) {
	db := &MysqlDB{UserName: userName, Password: password, IP: ip, Port: port, DBName: dbName}
	err := db.InitDB()
	if err != nil {
		t.Fatalf("Init db error :%v", err)
	}

	err = db.CreateDB("test")
	if err != nil {
		t.Fatalf("Create db error :%v", err)
	}
	db.CloseDB()

	// 执行orm操作
	// 设置默认数据库
	sql := strings.Join([]string{db.UserName, ":", db.Password, "@tcp(", db.IP, ":", db.Port, ")/", "test", "?charset=utf8"}, "")
	err = orm.RegisterDataBase("default", "mysql", sql, 30)
	if err != nil {
		t.Fatalf("RegisterDataBase Failed :%v", err)
	}

	err = orm.RegisterDataBase("tmp", "mysql", sql, 30)
	if err != nil {
		t.Fatalf("RegisterDataBase Failed :%v", err)
	}
	// 注册model
	orm.RegisterModel(new(model.User))
	// 创建 table
	err = orm.RunSyncdb("tmp", false, true)
	if err != nil {
		t.Fatalf("RunSyncdb Failed :%v", err)
	}

	o := orm.NewOrm()

	user := model.User{Name: "slene"}

	// 插入表
	_, err = o.Insert(&user)
	if err != nil {
		t.Fatalf("Insert Failed :%v", err)
	}

}
