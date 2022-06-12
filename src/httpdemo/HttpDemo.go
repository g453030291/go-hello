package httpdemo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func GetRequest() {
	client := resty.New()

	response, error := client.R().EnableTrace().Get("http://www.baidu.com")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(response.Body()))
}

func PostRequest() {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"username":"testuser", "password":"testpass"}`).
		SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
		Post("https://myapp.com/login")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp.Body()))
}
