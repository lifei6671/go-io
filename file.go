package iohelper

import (
	"os"
	"io/ioutil"
	"bytes"
)

//判断文件是否存在
func FileExists (filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 删除文件
func DeleteFile(filepath string) (bool) {
	if !FileExists(filepath) {
		return true;
	}
	if err := os.Remove(filepath); err != nil {
		panic(err);
	}
	return  true;
}

//读取指定路径的文件
func ReadFileContent(filePath string) (string) {
	fi,err := os.Open(filePath)
	if err != nil{
		panic(err);
	}
	defer fi.Close()

	fd,err := ioutil.ReadAll(fi)

	return bytes.NewBuffer(fd).String();
}