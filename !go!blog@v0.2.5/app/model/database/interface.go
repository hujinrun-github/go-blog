package dbinterface

/*
*该文件主要定义一些数据库操作的接口
*
**/

// AbstraceDB 抽象的数据库操作
type AbstraceDB interface {
	InitDB() error                            // 初始化数据库
	CreateDB(config interface{}) error        // 创建数据库
	InitTableOfDB(config interface{}) error   // 初始化数据表
	DeleteTableOfDB(config interface{}) error // 删除数据表
	InsertDataIntoDB(data interface{}) error  // 往数据库插入一条数据
	WriteToDB(content interface{}) error      // 将指定的内容写入数据库
	ReadFromDB() (interface{}, error)         // 从数据库读取指定的内容
	CloseDB() error                           // 关闭数据库
}
