package wechat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	token = "https://api.weixin.qq.com/cgi-bin/token"
)

type TokenInfo struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

func GetToken(appid, secret string) *TokenInfo {
	params := url.Values{}
	Url, _ := url.Parse(token)
	params.Set("grant_type", "client_credential")
	params.Set("appid", appid)
	params.Set("secret", secret)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res TokenInfo
	_ = json.Unmarshal(body, &res)
	return &res
}
