package config

import (
	"os"
	"path/filepath"
)

var (
	// _              = os.Setenv("CONFIG_DIR", "./.proglog/")
	CAFile         = configFile("ca.pem")
	ServerCertFile = configFile("server.pem")
	ServerKeyFile  = configFile("server-key.pem")
)

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	panic(err)
	// }
	return filepath.Join("/home/ndnhuy2504/workspace/throwaway/go-book-code/LetsGo/", ".proglog", filename)
}
