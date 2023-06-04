package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alitdarmaputra/fims-be/src/common"
	"github.com/alitdarmaputra/fims-be/src/config"
	seeds "github.com/alitdarmaputra/fims-be/src/docs/sql/seeder"
	"github.com/alitdarmaputra/fims-be/utils"
)

func main() {
	cfg := config.LoadConfigAPI(".")
	handleArgs(cfg)
}

func handleArgs(cfg *config.Api) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			db, err := common.NewMySQL(&cfg.Database)
			utils.PanicIfError(err)
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
	fmt.Println("done")
}
