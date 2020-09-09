package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"shopify-manager/pkg/config"
	"shopify-manager/pkg/constants"
	"shopify-manager/pkg/flow"
	"shopify-manager/pkg/infrastructure/util"
)

const LogFile = "./info.log"

func main() {
	flowType := flag.String("flow", constants.FLOW_TYPE_CREATE_INSTANCE, "flow type")
	flag.Parse()

	logfile, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("%sのオープンに失敗しました : %s\n", LogFile, err.Error())
		util.WaitEnter()
		return
	}
	defer logfile.Close()
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	config, err := config.LoadConfig()
	if err != nil {
		log.Printf("コンフィグファイルのロードに失敗しました : %s", err.Error())
		util.WaitEnter()
		return
	}

	if *flowType == constants.FLOW_TYPE_CREATE_INSTANCE {
		flow.CancelOrders(config)
	}

	util.WaitEnter()
}
