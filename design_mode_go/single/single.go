package single

import "sync"

/*
*单例模式：用来控制类型实例的数量，确保一个类型只有一个实例。
?俄汉模式:使用于早期初始化时创建的已经确定要加载的类型实例；
?懒汉模式：延迟加载，适合程序执行过程中条件成立才创建的类型实例
*/

//&饿汉模式-使用init函数实现
//*连接数据库

type databaseConn struct {
	//...
}

var dbConn *databaseConn

func init() {
	dbConn = &databaseConn{}
}

//GetInstance
func Db() *databaseConn {
	return dbConn
}

//*懒汉模式
//^使用并发原语Once

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstanceOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
