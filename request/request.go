package request

import (
	"fmt"
	"github.com/mozhiyun/go-easy-conis/utils"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	baseUrl    string
	httpClient *http.Client
	token      string
	secretKey  string
}

func NewClient(baseUrl, token, secretKey string) *Client {
	return &Client{
		baseUrl:    baseUrl,
		token:      token,
		secretKey:  secretKey,
		httpClient: &http.Client{},
	}
}

// MakeReq HTTP request helper
func (client *Client) MakeReq(path string, data map[string]string) ([]byte, error) {
	var tmp []string
	for k, v := range data {
		tmp = append(tmp, k+"="+v)
	}
	url := client.baseUrl + path + "?" + strings.Join(tmp, "&")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = client.MakeHeader(data)
	resp, err := client.doReq(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// helper
// doReq HTTP client
func (client *Client) doReq(req *http.Request) ([]byte, error) {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

func (client *Client) MakeHeader(data map[string]string) http.Header {
	header := make(http.Header)
	nonce := strconv.Itoa(int(time.Now().Unix())) + utils.GetRandomString(5)
	signature := client.MakeSignature(nonce, data)
	header.Add("Nonce", nonce)
	header.Add("Token", client.token)
	header.Add("Signature", signature)
	return header
}

func (client *Client) MakeSignature(nonce string, data map[string]string) (signature string) {
	var tmp []string
	tmp = append(tmp, client.token, client.secretKey, nonce)
	for k, v := range data {
		tmp = append(tmp, k+"="+v)
	}
	sort.Strings(tmp)
	signatureBeforeSha1 := strings.Join(tmp, "")
	return utils.Sha1Hex(signatureBeforeSha1)
}
