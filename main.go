package main

import (
	"log"

	"github.com/jcasado94/gobuyright/pkg/mongo"
	"github.com/jcasado94/gobuyright/pkg/mongo/service"
	"github.com/jcasado94/gobuyright/pkg/server"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatal("Unable to connect to mongo")
	}
	defer ms.Close()

	s := server.NewServer(
		service.NewIUserService(ms.Copy(), "buyright", "iuser"),
		service.NewUsageSelectionService(ms.Copy(), "buyright", "usageselection"),
		service.NewUsageService(ms.Copy(), "buyright", "usage"),
		service.NewItemService(ms.Copy(), "buyright", "item"),
	)

	s.Start()
}
