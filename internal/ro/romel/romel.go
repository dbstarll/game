package romel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dbstarll/game/internal/ro/dimension/weapon"
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

type Pager interface {
	getPage() int
	getDir() string
}

type Page struct {
	Page int `json:"page"`
}

func (p *Page) getPage() int {
	return p.Page
}

func (p *Page) getDir() string {
	return ""
}

type PageAndArms struct {
	Page int `json:"page"`
	Arms int `json:"arms"`
}

func (p *PageAndArms) getPage() int {
	return p.Page
}

func (p *PageAndArms) getDir() string {
	return fmt.Sprintf("/%d", p.Arms)
}

type Result struct {
	Code   int        `json:"code"`
	Status string     `json:"status"`
	Data   ResultData `json:"data"`
}

type ResultData struct {
	Page      int                      `json:"page"`
	PageSize  int                      `json:"pageSize"`
	PageCount int                      `json:"pageCount"`
	Total     int                      `json:"total"`
	List      []map[string]interface{} `json:"list"`
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
	return a.getListAndSave("card", &Page{Page: page}, Cards.Size())
}

func (a *RomelApi) GetHatList(page int) (*Result, error) {
	return a.getListAndSave("hat", &Page{Page: page}, Hats.Size())
}

func (a *RomelApi) GetEquipList(page int) (*Result, error) {
	return a.getListAndSave("equip", &Page{Page: page}, Equips.Size())
}

func (a *RomelApi) GetEquipListWithArms(page int, arms weapon.Weapon) (*Result, error) {
	return a.getListAndSave("equip", &PageAndArms{Page: page, Arms: arms.Code()}, Equips.SizeOfArms(arms))
}

func (a *RomelApi) GetPetList(page int) (*Result, error) {
	return a.getListAndSave2("pet", page, Pets.Size())
}

func (a *RomelApi) GetMonsterList(page int) (*Result, error) {
	return a.getListAndSave2("monster", page, Monsters.Size())
}

func (a *RomelApi) getListAndSave(list string, data Pager, check int) (*Result, error) {
	log.Printf("getListAndSave: %s-%+v", list, data)
	time.Sleep(time.Second * 5)
	query := &Query{
		apiId:    a.apiId,
		version:  a.version,
		language: a.language,
		data:     data,
	}

	if err := query.signature(a.apiSecret); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s/api/item/%sList?%s", a.domain, list, query.values().Encode())
	referer := fmt.Sprintf("http://%s/item/%s", a.domain, list)
	if req, err := a.newGetRequest(url, referer); err != nil {
		return nil, err
	} else if resp, err := http.DefaultClient.Do(req); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		result := &Result{}
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else if err := json.Unmarshal(body, result); err != nil {
			return nil, err
		} else if result.Data.Total == check {
			log.Printf("getListAndSave: %s-%+v -- not changed", list, data)
			result.Data.PageCount = 0
			return result, nil
		} else if err := ioutil.WriteFile(fmt.Sprintf("configs/romel/%s%s/%sList-%03d.json", list, data.getDir(), list, data.getPage()), body, 0644); err != nil {
			return nil, err
		} else {
			if data.getPage() == 1 {
				log.Printf("getListAndSave: %s-%+v -- updated[%d]: %d --> %d", list, data, result.Data.PageCount, check, result.Data.Total)
			}
			return result, nil
		}
	}
}

func (a *RomelApi) getListAndSave2(list string, page, check int) (*Result, error) {
	log.Printf("getListAndSave2: %s-%d", list, page)
	time.Sleep(time.Second * 5)
	query := &Query{
		apiId:    a.apiId,
		version:  a.version,
		language: a.language,
		data:     &Page{Page: page},
	}

	if err := query.signature(a.apiSecret); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://%s/api/%s/list?%s", a.domain, list, query.values().Encode())
	referer := fmt.Sprintf("http://%s/%s/", a.domain, list)
	if req, err := a.newGetRequest(url, referer); err != nil {
		return nil, err
	} else if resp, err := http.DefaultClient.Do(req); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		result := &Result{}
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		} else if err := json.Unmarshal(body, result); err != nil {
			return nil, err
		} else if result.Data.Total == check {
			log.Printf("getListAndSave2: %s-%d -- not changed", list, page)
			result.Data.PageCount = 0
			return result, nil
		} else if err := ioutil.WriteFile(fmt.Sprintf("configs/romel/%s/%sList-%03d.json", list, list, page), body, 0644); err != nil {
			return nil, err
		} else {
			if page == 1 {
				log.Printf("getListAndSave2: %s-%d -- updated[%d]: %d --> %d", list, page, result.Data.PageCount, check, result.Data.Total)
			}
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
