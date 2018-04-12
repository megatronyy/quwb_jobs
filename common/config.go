package common

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type Conf struct {
	Db map[string]string
}

//读取配置文件
func ReadConf(path string) (Conf, error) {
	var c Conf

	if fi, err := os.Open(path); err == nil {
		defer fi.Close()

		//读取配置文件
		if fd, err := ioutil.ReadAll(fi); err == nil {
			err = json.Unmarshal(fd, &c)
			return c, err
		} else {
			return c, err
		}
	} else {
		return c, err
	}
}
