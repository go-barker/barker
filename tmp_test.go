package main

import (
	"fmt"
	"testing"

	"github.com/corporateanon/barker/pkg/types"
	"github.com/go-resty/resty/v2"
)

func TestMain(t *testing.T) {

	type ClientError struct {
		Message string `json:"error,omitempty"`
	}

	client := resty.New()
	client.EnableTrace()
	client.SetHostURL("https://httpbin.org/")

	bot := &types.Bot{
		ID:    1,
		Title: "dfs",
		Token: "sdfg",
	}

	resultWrapper := &struct {
		Origin string `json:"data,omitempty"`
	}{Origin: ""}
	res, err := client.R().
		SetBody(bot).
		SetResult(resultWrapper).
		SetError(&ClientError{}).
		Post("/post")
	fmt.Println(resultWrapper)
	if err != nil {
		panic(err)
	}
	httpErr := res.Error()
	fmt.Println(httpErr)
	if httpErr != nil {
		panic(err)
	}

}
