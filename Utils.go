package goex

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0.0
	}

	switch v.(type) {
	case float64:
		return v.(float64)
	case string:
		vStr := v.(string)
		vF, _ := strconv.ParseFloat(vStr, 64)
		return vF
	default:
		panic("to float64 error.")
	}
}

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case string:
		vStr := v.(string)
		vInt, _ := strconv.Atoi(vStr)
		return vInt
	case int:
		return v.(int)
	case float64:
		vF := v.(float64)
		return int(vF)
	default:
		panic("to int error.")
	}
}

func ToUint64(v interface{}) uint64 {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case int:
		return uint64(v.(int))
	case float64:
		return uint64((v.(float64)))
	case string:
		uV, _ := strconv.ParseUint(v.(string), 10, 64)
		return uV
	default:
		panic("to uint64 error.")
	}
}

func ToInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case float64:
		return int64(v.(float64))
	default:
		vv := fmt.Sprint(v)

		if vv == "" {
			return 0
		}

		vvv, err := strconv.ParseInt(vv, 0, 64)
		if err != nil {
			return 0
		}

		return vvv
	}
}

func FloatToString(v float64, precision int) string {
	return fmt.Sprint(FloatToFixed(v, precision))
}

func FloatToFixed(v float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return math.Round(v*p) / p
}

func ValuesToJson(v url.Values) ([]byte, error) {
	parammap := make(map[string]interface{})
	for k, vv := range v {
		if len(vv) == 1 {
			parammap[k] = vv[0]
		} else {
			parammap[k] = vv
		}
	}
	return json.Marshal(parammap)
}

func GzipDecompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}

func FlateDecompress(data []byte) ([]byte, error) {
	return ioutil.ReadAll(flate.NewReader(bytes.NewReader(data)))
}

func GenerateOrderClientId(size int) string {
	uuidStr := strings.Replace(uuid.New().String(), "-", "", 32)
	return "goex" + uuidStr[0:size-5]
}

func MergeDepths(oldDepths DepthRecords, newDepths DepthRecords) (DepthRecords, error) {
	newRecord := DepthRecords{}
	oldIdx, newIdx := 0, 0

	for oldIdx < oldDepths.Len() && newIdx < newDepths.Len() {
		oldItem := oldDepths[oldIdx]
		newItem := newDepths[newIdx]

		if oldItem.Price == newItem.Price {
			if newItem.Amount > 0 {
				newRecord = append(newRecord, newItem)
			}
			oldIdx++
			newIdx++
		} else if oldItem.Price > newItem.Price {
			if newItem.Amount > 0 {
				newRecord = append(newRecord, newItem)
			}
			newIdx++
		} else if oldItem.Price < newItem.Price {
			newRecord = append(newRecord, oldItem)
			oldIdx++
		}
	}
	for ; oldIdx < oldDepths.Len(); oldIdx++ {
		newRecord = append(newRecord, oldDepths[oldIdx])
	}
	for ; newIdx < newDepths.Len(); newIdx++ {
		if newDepths[newIdx].Amount > 0 {
			newRecord = append(newRecord, newDepths[newIdx])
		}
	}
	return newRecord, nil
}

//CorrectDepth 使用Trade修正Orderbook 盘口
func CorrectDepths(depths DepthRecords, IsAskDepth bool, trade *Trade) DepthRecords {
	newRecord := DepthRecords{}
	for i := 0; i < depths.Len(); i++ {
		dr := depths[i]
		if IsAskDepth {
			if dr.Price > trade.Price { // 高于成交价的卖盘保留
				newRecord = append(newRecord, dr)
			} else if dr.Price == trade.Price && trade.Type != SELL {
				dr.Amount -= trade.Amount
				if dr.Amount > 0 {
					newRecord = append(newRecord, dr)
				}
			}
		} else {
			if dr.Price < trade.Price { // 低于成交价的买盘保留
				newRecord = append(newRecord, dr)
			} else if dr.Price == trade.Price && trade.Type != BUY {
				dr.Amount -= trade.Amount
				if dr.Amount > 0 {
					newRecord = append(newRecord, dr)
				}
			}
		}

	}
	return newRecord
}
