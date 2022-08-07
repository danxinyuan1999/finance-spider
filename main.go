package main

import (
	"github/danxinyuan1999/finace-spider/global"
	"github/danxinyuan1999/finace-spider/infra"
)

func main() {
	// init db
	global.DB = infra.NewDB()
	// init redis
	global.RBD = infra.NewRedisClient()
	// init resty
	global.HttpReq = infra.NewHttpClient()
	// auto migrate tables
	// start schedule job
}
