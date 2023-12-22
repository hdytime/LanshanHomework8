package main

import (
	"LanshanTeamwork8/api"
	"LanshanTeamwork8/dao"
)

func main() {
	dao.InitDatabase()
	api.InitRouter()
}
