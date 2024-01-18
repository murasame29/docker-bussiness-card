package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// LoadEnv 環境変数を読み込む (引数に.envファイルのパスを指定するとそのファイルを読み込む)
func LoadEnv(path ...string) {
	if path != nil {
		if err := godotenv.Load(path...); err != nil {
			panic(err)
		}
	}

	config := &config{}

	if err := env.Parse(&config.Server); err != nil {
		panic(err)
	}

	Config = config
}
