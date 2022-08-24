package wechat

import (
	"sync"
	"time"
)

type void struct{}

var member void

type Wechat struct {
	Appid string `json:"appid"`

	Secret string `json:"secret"`

	user map[string]void `json:"map"`

	tokenInfo *TokenInfo

	tokenLock *sync.RWMutex

	timer *time.Timer
}

func NweWechat(appid, secret string) *Wechat {
	return &Wechat{
		user:      make(map[string]void),
		tokenLock: new(sync.RWMutex),
		tokenInfo: nil,
		timer:     nil,
		Appid:     appid,
		Secret:    secret,
	}
}

func (w *Wechat) GetToken() string {
	if w.tokenInfo == nil {
		w.RefreshToken()
	}
	w.tokenLock.RLock()
	defer w.tokenLock.RUnlock()
	return w.tokenInfo.Token
}

func (w *Wechat) RefreshToken() {
	w.tokenLock.Lock()
	defer w.tokenLock.Unlock()
	w.tokenInfo = GetToken(w.Appid, w.Secret)
}

/**
 * 定时跟新token
 */
func (w *Wechat) renew(expiresIn int) {
	exp := time.Duration(expiresIn) * time.Second
	if w.timer != nil {
		w.timer.Reset(exp)
		return
	}
	w.timer = time.NewTimer(exp)
	go func() {
		for {
			select {
			case <-w.timer.C:
				w.RefreshToken()
			}
		}
	}()

}
