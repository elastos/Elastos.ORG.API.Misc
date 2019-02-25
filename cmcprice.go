package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const CMC_ENDPOINT_URL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=%d&convert=%s"

type Status struct {
	Timestamp     string
	Error_code    int
	Error_message string
	Elapsed       int
	Credit_count  int
}

type Price struct {
	Price              float64
	Volume_24h         float64
	Percent_change_1h  float64
	Percent_change_24h float64
	Percent_change_7d  float64
	Market_cap         float64
	Last_updated       string
}

type Quote struct {
	USD, CNY,BTC Price
}

type Plateform struct {
	Id int64
	Name string
	Symbol string
	Slug   			string
	Token_Address string
}

type Data struct {
	Id                 int64
	Name               string
	Symbol             string
	Slug               string
	Circulating_supply float64
	Total_supply       float64
	Max_supply         float64
	Date_added         string
	Num_market_pairs   int64
	Tags               []string
	Platform           Plateform
	Cmc_rank           int
	Last_updated       string
	Quote              Quote
}

type CmcResponse struct {
	Status Status
	Data   []Data
}

var dba = db.NewInstance()

func init() {
	go func() {
		i := 0
		for{
			sleepy , err := time.ParseDuration(strconv.Itoa(config.Conf.Cmc.Inteval)+"m")
			if err != nil {
				fmt.Printf("%s",err.Error())
				os.Exit(-1)
			}
			<- time.After(sleepy)
			cmcResponseUSD , err := fetchPrice(i,"USD")
			if err != nil {
				fmt.Printf("Error init cmc price %s", err.Error())
				continue
			}
			cmcResponseCNY , err := fetchPrice(i,"CNY")
			if err != nil {
				fmt.Printf("Error init cmc price %s", err.Error())
				continue
			}
			cmcResponseBTC , err := fetchPrice(i,"BTC")
			if err != nil {
				fmt.Printf("Error init cmc price %s", err.Error())
				continue
			}
			err = saveToDb(cmcResponseUSD,cmcResponseCNY,cmcResponseBTC)
			if err != nil {
				fmt.Printf("Error init cmc price %s", err.Error())
			}
			if i == len(config.Conf.Cmc.ApiKey) -1 {
				i = 0
			}else{
				i++
			}
		}
	}()
}

func saveToDb(cmcResponseUSD , cmcResponseCNY , cmcResponseBTC CmcResponse) error{
	tx , err := dba.Begin()
	if err != nil {
		return err
	}
	data := cmcResponseUSD.Data
	for i:=0;i<len(data);i++ {
		_ , err = tx.Exec("insert into chain_cmc_price(id,name,symbol,`rank`,price_usd,price_cny,price_btc,24h_volume_usd,market_cap_usd,available_supply,total_supply,max_supply,percent_change_1h,percent_change_24h,percent_change_7d,last_updated,24h_volume_btc,market_cap_btc,local_system_time,24h_volume_cny,market_cap_cny,platform_symbol,platform_token_address,num_market_pairs) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			strconv.Itoa(int(data[i].Id)),
			data[i].Name,
			data[i].Symbol,
			strconv.Itoa(data[i].Cmc_rank),
			strconv.FormatFloat(data[i].Quote.USD.Price,'f',8,64),
			strconv.FormatFloat(cmcResponseCNY.Data[i].Quote.CNY.Price,'f',8,64),
			strconv.FormatFloat(cmcResponseBTC.Data[i].Quote.BTC.Price,'f',8,64),
			strconv.FormatFloat(data[i].Quote.USD.Volume_24h,'f',8,64),
			strconv.FormatFloat(data[i].Quote.USD.Market_cap,'f',8,64),
			strconv.FormatFloat(data[i].Circulating_supply,'f',8,64),
			strconv.FormatFloat(data[i].Total_supply,'f',8,64),
			strconv.FormatFloat(data[i].Max_supply,'f',8,64),
			strconv.FormatFloat(data[i].Quote.USD.Percent_change_1h,'f',8,64),
			strconv.FormatFloat(data[i].Quote.USD.Percent_change_24h,'f',8,64),
			strconv.FormatFloat(data[i].Quote.USD.Percent_change_7d,'f',8,64),
			data[i].Quote.USD.Last_updated,
			strconv.FormatFloat(cmcResponseBTC.Data[i].Quote.BTC.Volume_24h,'f',8,64),
			strconv.FormatFloat(cmcResponseBTC.Data[i].Quote.BTC.Market_cap,'f',8,64),
			time.Now(),
			strconv.FormatFloat(cmcResponseCNY.Data[i].Quote.CNY.Volume_24h,'f',8,64),
			strconv.FormatFloat(cmcResponseCNY.Data[i].Quote.CNY.Market_cap,'f',8,64),
			data[i].Platform.Symbol,
			data[i].Platform.Token_Address,
			data[i].Num_market_pairs)
		if err != nil {
			dba.Rollback(tx)
			return err
			return nil
		}
	}
	dba.Commit(tx)
	return nil
}

func fetchPrice(i int,curr string) (CmcResponse,error){
	url := fmt.Sprintf(CMC_ENDPOINT_URL, 200, curr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CmcResponse{},err
	}
	println(config.Conf.Cmc.ApiKey[i])
	req.Header = map[string][]string{
		"X-CMC_PRO_API_KEY": []string{config.Conf.Cmc.ApiKey[i]},
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CmcResponse{},err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CmcResponse{},err
	}
	cmcResp := CmcResponse{}
	err = json.Unmarshal(body, &cmcResp)
	if err != nil {
		return CmcResponse{},err
	}
	return cmcResp,nil
}