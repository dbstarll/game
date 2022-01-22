package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RomelApi struct {
	version   string
	domain    string
	language  string
	apiId     string
	apiSecret string
}

type Query struct {
	apiId    string
	version  string
	language string
	sign     string
	time     int64
	data     interface{}
	dataJson string
}

type Page struct {
	Page int `json:"page"`
}

type Result struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Data   ResultData `json:"data"`
}

type ResultData struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

func NewRomelApi(secret string) *RomelApi {
	return &RomelApi{
		version:   "1.0.0",
		domain:    "romel.wiki",
		language:  "cn",
		apiId:     "rowiki",
		apiSecret: secret,
	}
}

func (a *RomelApi) GetCardList(page int) (*Result, error) {
	log.Printf("GetCardList: %d", page)
	query := &Query{
		apiId:    a.apiId,
		version:  a.version,
		language: a.language,
		data:     &Page{Page: page},
	}

	if err := query.signature(a.apiSecret); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s/api/item/cardList?%s", a.domain, query.values().Encode())
	referer := fmt.Sprintf("http://%s/item/card", a.domain)
	if req, err := a.newGetRequest(url, referer); err != nil {
		return nil, err
	} else if resp, err := http.DefaultClient.Do(req); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		result := &Result{}
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else if err := ioutil.WriteFile(fmt.Sprintf("configs/romel/card/cardList-%03d.json", page), body, 0644); err != nil {
			return nil, err
		} else if err := json.Unmarshal(body, result); err != nil {
			return nil, err
		} else {
			return result, nil
		}
	}
}

func (a *RomelApi) newGetRequest(url, referer string) (*http.Request, error) {
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	} else {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")
		req.Header.Set("Referer", referer)
		return req, nil
	}
}

func (q *Query) signature(secret string) error {
	if dataJson, err := json.Marshal(q.data); err != nil {
		return err
	} else {
		q.dataJson = string(dataJson)
		q.time = time.Now().Unix()
		m := md5.Sum([]byte(fmt.Sprintf("appid=%sdata=%ssecret=%stime=%d", q.apiId, dataJson, secret, q.time)))
		q.sign = hex.EncodeToString(m[:])
		return nil
	}
}

func (q *Query) values() *url.Values {
	values := &url.Values{}
	values.Set("appid", q.apiId)
	values.Set("version", q.version)
	values.Set("language", q.language)
	values.Set("sign", q.sign)
	values.Set("time", strconv.FormatInt(q.time, 10))
	values.Set("data", q.dataJson)
	return values
}
