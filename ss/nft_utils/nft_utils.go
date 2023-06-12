package nft_utils

import (
	"log"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

var NftServerUrl string = "/nftServer"

func GetGas() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit":   "20",
			"sort":    "name",
			"order":   "asc",
			"random":  strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get(NftServerUrl)
	log.Println(resp, err)
}
