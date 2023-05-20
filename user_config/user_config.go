package userconfig

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type UserConfig struct {
	MusicPath string `json:"musicPath"`
}

func New() *UserConfig {
	userConfig := UserConfig{}

	return parseConfig(&userConfig)
}

func (uc *UserConfig) SetMusicPath(path string) {
	uc.MusicPath = path

	jsonBytes, err := json.Marshal(uc)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join("user_config", "user-config.json"), jsonBytes, fs.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}
}

func (uc *UserConfig) GetMusicPath() string {
	return uc.MusicPath
}

func parseConfig(userConfig *UserConfig) *UserConfig {
	file, err := os.Open(filepath.Join("user_config", "user-config.json"))

	defer file.Close()

	if err != nil {
		return userConfig
	}

	fileStat, err := file.Stat()

	if err != nil {
		return userConfig
	}

	jsonBytes := make([]byte, fileStat.Size())

	_, err = bufio.NewReader(file).Read(jsonBytes)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(jsonBytes, &userConfig)

	return userConfig
}
