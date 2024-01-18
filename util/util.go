package util

import (
	"fmt"
	"github.com/cloudflare/cfssl/log"
	"github.com/pkg/errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ValidateAndReturnAbsConf 구성 파일 경로와 홈 디렉터리 간에 충돌이 없는지 확인합니다. 충돌이 없으면 구성 파일과 홈 디렉터리의 절대 경로를 반환합니다
func ValidateAndReturnAbsConf(configFilePath, homeDir, cmdName string) (string, string, error) {
	var err error
	var homeDirSet bool
	var configFileSet bool

	defaultConfig := GetDefaultConfigFile(cmdName)

	if configFilePath == "" {
		configFilePath = defaultConfig
	} else {
		configFileSet = true
	}

	if homeDir == "" {
		homeDir = filepath.Dir(defaultConfig)
	} else {
		homeDirSet = true
	}

	homeDir, err = filepath.Abs(homeDir)
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to get full path of config file")
	}
	homeDir = strings.TrimRight(homeDir, "/")

	if configFileSet && homeDirSet {
		log.Warning("Using both --config and --home CLI flags; --config will take precedence")
	}

	if configFileSet {
		configFilePath, err = filepath.Abs(configFilePath)
		if err != nil {
			return "", "", errors.Wrap(err, "Failed to get full path of configuration file")
		}
		return configFilePath, filepath.Dir(configFilePath), nil
	}

	configFile := filepath.Join(homeDir, filepath.Base(defaultConfig))
	return configFile, homeDir, nil
}

// GetDefaultConfigFile 사용 메시지에 표시할 구성 파일의 기본 경로를 가져옵니다.
func GetDefaultConfigFile(cmdName string) string {
	if cmdName == "fabric-ca-server" {
		var fname = fmt.Sprintf("%s-config.yaml", cmdName)
		home := "."
		envs := []string{"FABRIC_CA_SERVER_HOME", "FABRIC_CA_HOME", "CA_CFG_PATH"}
		for _, env := range envs {
			envVal := os.Getenv(env)
			if envVal != "" {
				home = envVal
				break
			}
		}
		return path.Join(home, fname)
	}

	var fname = fmt.Sprintf("%s-config.yaml", cmdName)
	var home string
	envs := []string{"FABRIC_CA_CLIENT_HOME", "FABRIC_CA_HOME", "CA_CFG_PATH"}
	for _, env := range envs {
		envVal := os.Getenv(env)
		if envVal != "" {
			home = envVal
			return path.Join(home, fname)
		}
	}
	return path.Join(os.Getenv("HOME"), ".fabric-ca-client", fname)
}
