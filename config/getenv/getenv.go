package getenv

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiCepUrl     string
	ApiWeatherUrl string
)

func LoadConfig(filenames ...string) bool {
	configLoaded := false
	for _, file := range filenames {
		if err := godotenv.Load(file); err == nil {
			configLoaded = true

			ApiCepUrl = os.Getenv("API_CEP_URL")
			ApiWeatherUrl = os.Getenv("API_WEATHER_URL")
			break
		}
	}
	return configLoaded
}
