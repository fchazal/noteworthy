package server

import (
	"github.com/fchazal/noteworthy/server/api"
	"github.com/fchazal/noteworthy/server/storage"
)

func Start() {
	storage.Open("./noteworthy.json")
	storage.Save()

	api.Initialize(":3000")
}
