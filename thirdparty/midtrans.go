package thirdparty

import (
	"midtrans-gopay-test/config"

	"github.com/veritrans/go-midtrans"
)

func GetMidtransCoreGateway() midtrans.CoreGateway {
	midclient := midtrans.NewClient()
	midclient.ServerKey = config.Conf.MidtransServerKey
	midclient.ClientKey = config.Conf.MidtransClientKey
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway := midtrans.CoreGateway{
		Client: midclient,
	}

	return coreGateway
}
