package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL       string
	Platform    string
	tokenSecret string
}

// home 디렉토리에 있는 .gatorconfig.json 파일을 읽어와 Config 구조체에 저장하고 반환하는 함수
func Read() (Config, error) {
	// .env 파일 load해서 리눅스 환경변수에 추가
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading config file: %w", err)
	} // godotenv.Load(filenames ...string) 함수에 불러들일 파일들의 path들을 입력해도 된다. (입력하지 않으면 기본값 .env 파일 로드)

	// Getenv 함수로 환경변수를 불러올 수 있음
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	config := Config{
		DBURL:       dbURL,
		Platform:    platform,
		tokenSecret: tokenSecret,
	}

	return config, nil
}
