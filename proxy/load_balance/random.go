package load_balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	curIndex int
	rss      []string
	//观察主题
	conf LoadBalanceConf
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	r.rss = append(r.rss, params...)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}

func (r *RandomBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

func (r *RandomBalance) SetConf(conf LoadBalanceConf) {
	r.conf = conf
}

func (r *RandomBalance) Update() {
}
