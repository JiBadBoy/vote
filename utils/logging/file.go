package logging

import (
	"fmt"
	"os"
	"time"
	"vote/utils/file"
	"vote/utils/setting"
)

//获取日志文件存储路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

//获取日志文件文件名
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
}

//打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err :%v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
