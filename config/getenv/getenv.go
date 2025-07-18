package getenv

import (
	"os"
)

var (
	ApiCepUrl     = os.Getenv("API_CEP_URL")
	ApiWeatherUrl = os.Getenv("API_WEATHER_URL")
)
