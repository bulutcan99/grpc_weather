package fetch

import (
	"errors"
	"fmt"
	"github.com/bulutcan99/grpc_weather/dto"
	"github.com/bulutcan99/grpc_weather/pkg/env"
	decoder "github.com/goccy/go-json"
	"go.uber.org/zap"
	"io"
	"net/http"
)

var (
	apiURL      = &env.Env.WeatherUrl
	apiKey      = &env.Env.WeatherApiKey
	defaultCity = &env.Env.DefaultWeatherCity
)

func buildURL(city string) string {
	return fmt.Sprintf("%s?key=%s&q=%s&aqi=no", *apiURL, *apiKey, city)
}

func fetchData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseData(data []byte) (dto.WeatherData, error) {
	var result dto.WeatherData
	if err := decoder.Unmarshal(data, &result); err != nil {
		return dto.WeatherData{}, err
	}

	return result, nil
}

func FetchWeather(city string) (dto.WeatherData, error) {
	url := buildURL(city)

	body, err := fetchData(url)
	if err != nil {
		zap.S().Error("Failed to fetch data: ", err)
		return dto.WeatherData{}, errors.New("failed to fetch data")
	}

	data, err := parseData(body)
	if err != nil {
		zap.S().Error("Failed to parse data: ", err)
		return dto.WeatherData{}, errors.New("failed to parse data")
	}

	return data, nil
}

func FetchDefaultWeather() (dto.WeatherData, error) {
	return FetchWeather(*defaultCity)
}
