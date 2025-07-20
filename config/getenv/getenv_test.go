package getenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	assert.True(t, LoadConfig("../../.env", ".env"))
	assert.Equal(t, "http://viacep.com.br/ws/%s/json/", ApiCepUrl)
	assert.Equal(t, "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", ApiWeatherUrl)
}
