package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Upgrade struct {
	OsCode   string
	OsName   string
	OsIcon   string
	CmsCode  string
	TvOsCode string
}

func GetUpgrade(id string) (upgrade Upgrade, err error) {

	upgrade = Upgrade{OsCode: id}
	// logger.Debug("upgrade.Id :{%v}", upgrade.OsCode)
	c, err := redis.Dial("tcp", "10.154.252.153:6379")
	if err != nil {
		logger.Debug("Connect to redis error:{%v}", err)
		return
	}
	defer c.Close()
	// "pkg:"..appkey..":id"
	key := fmt.Sprintf("pkg:%s:id", id)
	// logger.Debug("key :{%v}", key)
	pkg, err := redis.String(c.Do("GET", key))
	upgrade.TvOsCode = pkg
	// logger.Debug("orm.RegisterDataBase error:{%v}", upgrade.TvOsCode)
	// logger.Error("orm.RegisterDataBase error:{%v}", err)
	return upgrade, err
}
