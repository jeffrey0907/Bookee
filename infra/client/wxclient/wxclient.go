package wxclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	appId     = `wx9fa596a434fe6362`
	appSecret = `35647b17536a43f173697f5158c825c4`
)

func Code2Session(code string) (openid, sessionKey, unionid string, err error) {
	url := fmt.Sprintf(`https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`, appId, appSecret, code)
	resp, err := http.Get(url)
	if err == nil {
		res := struct {
			Openid     string `json:"openid"`
			SessionKey string `json:"session_key"`
			UnionId    string `json:"unionid"`
			ErrCode    int32  `json:"errcode"`
			ErrMsg     string `json:"errMsg"`
		}{}

		if res.ErrCode != 0 {
			err = errors.New(res.ErrMsg)
			return
		}
		rawbody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			err = json.Unmarshal(rawbody, &res)
			if err == nil {
				openid = res.Openid
				sessionKey = res.SessionKey
				unionid = res.UnionId
			}
		}
	}
	return
}
