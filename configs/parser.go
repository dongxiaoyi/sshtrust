package configs

import (
	"github.com/dongxiaoyi/toolBox/pkg"
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
	"path/filepath"
)


func ParserConfig() *goconfig.ConfigFile {
	dir := pkg.AbsPath()
	ini := filepath.Join(dir, "configs/settings.ini")
	config, err := goconfig.LoadConfigFile(ini)
	if err != nil {
		fmt.Println("Get config file error")
		os.Exit(-1)
	}
	return config
}
