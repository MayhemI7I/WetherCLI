package weather

import (
	"fmt"
	"io"
	"local/pkg/geo"
	"net/http"
	"net/url"
	"strconv"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return ""
	}
	defer resp.Body.Close()
	return string(body)
}
