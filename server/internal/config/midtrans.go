package config

import (
	"log"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var CoreClient coreapi.Client
var SnapClient snap.Client

func InitMidtrans() {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	env := os.Getenv("NODE_ENV")

	snapClient := snap.Client{}
	snapClient.New(serverKey, getMidtransEnv(env))
	SnapClient = snapClient

	coreClient := coreapi.Client{}
	coreClient.New(serverKey, getMidtransEnv(env))
	CoreClient = coreClient

	log.Println("Midtrans Snap & CoreAPI client initialized")
}

func getMidtransEnv(env string) midtrans.EnvironmentType {
	if env == "production" {
		return midtrans.Production
	}
	return midtrans.Sandbox
}
