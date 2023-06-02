package main

type Message struct {
	HttpCode int    `json:"http_code"`
	Tag      string `json:"tag"`
	Message  string `json:"message"`
}
