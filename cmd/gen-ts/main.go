package main

import (
	"os"

	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	converter := typescriptify.New().
		Add(types.Bot{}).
		Add(types.Campaign{}).
		Add(types.User{}).
		Add(types.Delivery{}).
		AddEnum(types.AllDeliveryStates).
		Add(dao.DeliveryTakeResult{})

	converter.CreateInterface = true
	converter.BackupDir = os.TempDir()

	err := converter.ConvertToFile("ts/types.ts")
	if err != nil {
		panic(err.Error())
	}

}
