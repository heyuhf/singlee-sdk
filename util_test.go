package singlee_sdk

import (
	"fmt"
	"testing"
)

func TestGetSign(t *testing.T) {
	data := map[string]string{
		"shop_id":         "czyhgtcsp1",
		"pos_id":          "10415992",
		"mch_order_no":    "210820200435124194598",
		"pay_qrcode":      "135165195081678886",
		"amount":          "1670",
		"mch_info_string": "沧州市运河区甘太泰食品店收款16.7元",
		"timestamp":       "20210820200435",
		"sign":            "6F859AFFADE528B57BBFDBD86AA1089DC591B08577256396B728DA90B92870BB",
	}
	sign := GetSign(data, "A")
	if sign != "6F859AFFADE528B57BBFDBD86AA1089DC591B08577256396B728DA90B92870BB" {
		t.Fail()
	}
	fmt.Println(sign)
}
