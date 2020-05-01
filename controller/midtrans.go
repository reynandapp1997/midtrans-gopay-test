package controller

import (
	"midtrans-gopay-test/thirdparty"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veritrans/go-midtrans"
)

func RequestChart(c *gin.Context) {
	midtransClient := thirdparty.GetMidtransCoreGateway()

	price := 200000
	qty := 2
	grossAmt := price * qty

	orderID := c.Query("order_id")
	charge := &midtrans.ChargeReq{
		PaymentType: midtrans.SourceGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(grossAmt),
		},
		Items: &[]midtrans.ItemDetail{
			midtrans.ItemDetail{
				Brand:        "test-brand",
				Category:     "test-category",
				ID:           "test-id",
				MerchantName: "test-merchant",
				Name:         "test-name",
				Price:        int64(price),
				Qty:          int32(qty),
			},
		},
		Gopay: &midtrans.GopayDetail{
			EnableCallback: true,
			CallbackUrl:    "http://example.com/",
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: "test@test.com",
			FName: "first",
			LName: "last",
			Phone: "+62",
		},
	}

	response, err := midtransClient.Charge(charge)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"error":        true,
			"errorMessage": err.Error(),
			"data":         map[string]interface{}{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":      false,
		"error":        true,
		"errorMessage": nil,
		"data":         response,
	})
}

func GetTransactionStatus(c *gin.Context) {
	midtransClient := thirdparty.GetMidtransCoreGateway()

	orderID := c.Query("order_id")

	response, err := midtransClient.Status(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"error":        true,
			"errorMessage": err.Error(),
			"data":         map[string]interface{}{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":      false,
		"error":        true,
		"errorMessage": nil,
		"data":         response,
	})
}
