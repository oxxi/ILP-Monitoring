package models

import "time"

type RequestEntity struct {
	UUID       string    `json:"uuid"  redis:"uuid"`
	Resource   string    `json:"resource"  redis:"resource"`
	StatusCode uint64    `json:"statusCode" redis:"statusCode"`
	Status     string    `json:"status" redis:"status"`
	TimeStart  time.Time `json:"timeStart" redis:"timeStart"`
	TimeEnd    time.Time `json:"timeEnd" redis:"timeEnd"`
	Duration   uint64    `json:"duration" redis:"duration"`
}

type ResourceConfig struct {
	UUID       string `json:"uuid" redis:"uuid"`
	Name       string `json:"name" redis:"name"`
	IsLogin    bool   `json:"isLogin" redis:"isLogin"`
	Url        string `json:"url" redis:"url"`
	Parameters string `json:"parameters" redis:"parameters"`
	Method     string `json:"method" redis:"method"`
	User       string `json:"user" redis:"user"`
	Password   string `json:"password" redis:"password"`
	HasToken   string `json:"hasToken" redis:"hasToken"`
}
