package singlee_sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func New(key string) *Client {
	return &Client{
		HttpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Key: key,
		// NotifyKey:  notifyKey,
		Host:   "https://paygate.leshuazf.com",
		Logger: logrus.New(),
	}
}

type Client struct {
	HttpClient *http.Client
	Key        string // sign key
	NotifyKey  string
	Host       string
	Logger     *logrus.Logger
}

func (c *Client) Payment(params map[string]string) (resp PaymentResponse, err error) {
	url := "https://dz.singlee.com.cn:19023/Payment"
	err = c.Request(url, params, &resp)
	return
}

func (c *Client) OrderQuery(shopId, posId, mchOrderNo string) (resp OrderQueryResponse, err error) {
	url := "https://dz.singlee.com.cn:19023/OrderQuery"
	params := map[string]string{
		"shop_id":      shopId,
		"pos_id":       posId,
		"mch_order_no": mchOrderNo,
	}
	err = c.Request(url, params, &resp)
	return
}

func (c *Client) getLoggerEntity(keyword string) *logrus.Entry {
	return c.Logger.WithField("singleeSdk", keyword)
}

func (c *Client) Request(url string, params map[string]string, v interface{}) (err error) {

	params["timestamp"] = time.Now().Format("20060102150405")
	params["sign"] = GetSign(params, c.Key)
	jStr, _ := json.Marshal(params)
	c.getLoggerEntity("requestParams").Info(string(jStr))
	req, err := http.NewRequest("POST", url, bytes.NewReader(jStr))
	if err != nil {
		c.getLoggerEntity("requestError").Error(err)
		return
	}

	rsp, err := c.HttpClient.Do(req)
	if err != nil {
		c.getLoggerEntity("doRequestError").Error(err)
		return
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		c.getLoggerEntity("readBodyError").Error(err)
		return
	}
	err = json.Unmarshal(body, v)
	c.Logger.WithField("singlee-sdk", "resp").Info(string(body))
	return
}
