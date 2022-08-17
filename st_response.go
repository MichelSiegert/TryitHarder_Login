package main

type Response struct {
	Mail       string
	httpstatus int
	Message    string
	Data       string
	User       []userData
}
