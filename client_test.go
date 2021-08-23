package singlee_sdk

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

var client = New("AcoHRL6UdeUgCzeF0bxoJbDT6ZjP6pU0")

func init() {
	client.Logger.Formatter = new(logrus.JSONFormatter)
}

func TestClient(t *testing.T) {
	params := map[string]string{
		"shop_id":         "czgs001",
		"pos_id":          "00254428",
		"mch_order_no":    "test012",
		"pay_qrcode":      "134514428510372310",
		"amount":          "1",
		"mch_info_string": "测试收款1分",
	}
	resp, err := client.Payment(params)
	fmt.Println(resp, err)

	fmt.Println(resp.GetPayStatus(), resp.GetFailedMsg())
}

func TestClient_OrderQuery(t *testing.T) {
	resp, err := client.OrderQuery(
		"czgs001",
		"00254428",
		"test012",
	)
	fmt.Println(resp, err)
	fmt.Println(resp.GetPayStatus(), resp.GetAmount(), resp.GetFailedMsg())
}
