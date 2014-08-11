package mngo

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const BaseURI = "https://api.mndigital.com/"

var APIKey = ""
var APISecret = ""

type stringMap map[string]interface{}

func GetMP3(id int64, ip string) string {
	if ip == "" {
		ip = "127.0.0.1"
	}

	v := url.Values{}
	v.Set("method", "Radio.GetMediaLocation")
	v.Set("format", "json")
	v.Set("trackId", strconv.FormatInt(id, 10))
	v.Set("assetCode", "014")
	v.Set("protocol", "http")
	v.Set("userIP", ip)
	v.Set("apiKey", APIKey)
	v.Set("timestamp", time.Now().String())

	b := []byte(v.Encode())

	key := hmac.New(md5.New, []byte(APISecret))
	key.Write(b)

	query := v.Encode() + "&signature=" + hex.EncodeToString(key.Sum(nil))

	mnURI := BaseURI + "?" + query

	res, err := http.Get(mnURI)
	if err != nil {
		return ""
	}

	var results stringMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&results)
	if err != nil {
		return ""
	}

	fmt.Printf("it is %s\n", results)

	return mnURI
}
