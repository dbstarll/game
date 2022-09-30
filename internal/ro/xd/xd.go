package xd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode/utf8"
)

const dataFile = "/Users/dbstar/git/github.com/dbstarll/game/web/static/moniqi_data_file_1637826700.js"

func init() {
	if file, err := os.Open(dataFile); err != nil {
		zap.S().Fatalf("open data file failed: %+v", errors.WithStack(err))
	} else {
		defer file.Close()

		if err := loadAll(file); err != nil {
			zap.S().Fatalf("load data failed: %+v", err)
		}
	}
}

func loadAll(file *os.File) error {
	scanner, bufferSize := bufio.NewScanner(file), 10*1024*1024
	scanner.Buffer(make([]byte, bufferSize), bufferSize)
	for scanner.Scan() {
		line := scanner.Bytes()
		if idx := bytes.IndexRune(line, '='); idx > 0 {
			if idxStart := bytes.IndexRune(line, '{'); idxStart < 0 {
				continue
			} else if idxEnd := bytes.LastIndexFunc(line, func(r rune) bool {
				return r == '}'
			}); idxEnd <= idxStart {
				continue
			} else {
				head := strings.TrimSpace(string(line[:idx]))
				data := line[idxStart : idxEnd+1]
				switch head {
				case "MONIQI_DATA.CommonFun.BaseLvRate":
					fmt.Printf("BaseLvRate: %s\n", data)
				case "MONIQI_DATA.CommonFun.HpRate":
					fmt.Printf("HpRate: %s\n", data)
				case "MONIQI_DATA.CommonFun.BaseHp":
					fmt.Printf("BaseHp: %s\n", data)
				//case "MONIQI_DATA.attrratio_data":
				//	fmt.Printf("attrratio_data: %s\n", data)
				//case "MONIQI_DATA.attrvalue_data":
				//	fmt.Printf("attrvalue_data: %s\n", data)
				//case "MONIQI_DATA.attr220value_data":
				//	fmt.Printf("attr220value_data: %s\n", data)
				case "MONIQI_DATA.baselevel_data":
					if levels, err := loadBaseLevels(data); err != nil {
						return err
					} else {
						BaseLevels = levels
					}
				case "MONIQI_DATA.job_data":
					if jobs, err := loadJobs(data); err != nil {
						return err
					} else {
						Jobs = jobs
					}
				case "MONIQI_DATA.type_data":
					if jobTypes, err := loadJobTypes(data); err != nil {
						return err
					} else {
						JobTypes = jobTypes
					}
				case "MONIQI_DATA.role_data":
					if roles, err := loadRoles(data); err != nil {
						return err
					} else {
						Roles = roles
					}
				case "MONIQI_DATA.site_data":
					if sites, err := loadSites(data); err != nil {
						return err
					} else {
						Sites = sites
					}
				case "MONIQI_DATA.equip_data":
					if equips, err := loadEquips(data); err != nil {
						return err
					} else {
						Equips = equips
					}
				case "MONIQI_DATA.card_data":
					if cards, err := loadCards(data); err != nil {
						return err
					} else {
						Cards = cards
					}
				case "MONIQI_DATA.buff_data":
					if buffs, err := loadBuffs(data); err != nil {
						return err
					} else {
						Buffs = buffs
					}
				default:
					zap.S().Debugf("ignore [%d] %s", len(data), head)
				}
			}
		}
	}
	return scanner.Err()
}

type PropertyDetector struct {
	count map[string]int
}

func NewPropertyDetector() *PropertyDetector {
	return &PropertyDetector{count: make(map[string]int)}
}

func (d *PropertyDetector) Detect(item interface{}) {
	if mapItem, ok := item.(map[string]interface{}); ok {
		for k, v := range mapItem {
			key := fmt.Sprintf("%s: %s", k, reflect.TypeOf(v))
			if oc, ok := d.count[key]; ok {
				d.count[key] = oc + 1
			} else {
				d.count[key] = 1
			}
		}
	}
}

func (d *PropertyDetector) Show() {
	var items []string
	for k, v := range d.count {
		items = append(items, fmt.Sprintf("%s - %d", k, v))
	}
	sort.Strings(items)
	for idx, item := range items {
		fmt.Printf("unknown buff property: [%d]%s\n", idx, item)
	}
}

func unmarshalDisallowUnknownFields(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(v); err != nil {
		return errors.Wrapf(err, "decode failed: [%s] - %s", reflect.TypeOf(v), data)
	} else {
		return nil
	}
}

type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) error {
	if r, _ := utf8.DecodeRune(data); r == '1' {
		*b = true
		return nil
	} else if r == '0' {
		*b = false
		return nil
	} else {
		return errors.Errorf("unknown Bool: %s", data)
	}
}
