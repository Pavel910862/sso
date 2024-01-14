// парсит config.yaml

package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct { // создаем структуру для config.yaml
	Env         string        `yaml:"env" env-default:"local"` // стракттеги нужны чтобы искать их в config.yaml по имени
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// теперь нужно научить наш код превращать config.yaml в структуру Config используес cleanenv
func MustLoad() *Config { // используем Must чтобы не возвращать ошибку приложение упадет при запуске
	path := fetchConfigPath() // читаем путь
	if path == "" {
		panic("config path is empty")
	}

	// проверяем что по этому пути что-то есть
	//if _, err := os.Stat(path); os.IsNotExist(err) {
	//	panic("config file does not exist" + path)
	//}

	//var cfg Config // переменная в которую будет сохранен объект конфига

	//if err := cleanenv.ReadConfig("C:/Users/Алёна Валерьевна/Desktop/sso/сonfig/local.yaml", &cfg); err != nil { //непосредственно для парсинга файла
	//	panic("failed to read config" + err.Error())
	//}

	return MustLoadByPath(path) //возвращаем объект конфига
}

func MustLoadByPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist111: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string { //получает конфиг сначала из флага потом из переменных окружения
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config file") // sso --config=./path...
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH") // CONFIG_PATH=./path/to/config/config.yaml sso
	}
	return res
}
