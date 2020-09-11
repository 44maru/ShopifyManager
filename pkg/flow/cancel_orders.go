package flow

import (
	"log"

	"shopify-manager/pkg/api/shopify"
	"shopify-manager/pkg/config"
)

func CancelOrders(config *config.Config) {

	err := shopify.CancelOrders(config)
	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		return
	}

	log.Println("キャンセル処理成功")
}
