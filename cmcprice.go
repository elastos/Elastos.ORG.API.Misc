package main

import (
	"encoding/json"
	"fmt"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"github.com/pkg/errors"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	CMC_ENDPOINT_URL = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=%d&convert=%s"
	BGX_ENDPOINT_URL = "https://www.gaex.com/svc/portal/api/v2/publicinfo"
	HBG_ENDPOINT_URL = "https://api.huobi.pro/market/history/trade?symbol=elabtc"
)

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
	USD, CNY, BTC Price
}

type Plateform struct {
	Id            int64
	Name          string
	Symbol        string
	Slug          string
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
var dbaforela = db.NewInstance()

func init() {
	if config.Conf.Cmc.Enable {
		go func() {
			i := -1
			sleepy, err := time.ParseDuration(config.Conf.Cmc.Inteval)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				os.Exit(-1)
			}
			for {
				if i == len(config.Conf.Cmc.ApiKey)-1 {
					i = 0
				} else {
					i++
				}
				cmcResponseUSD, err := fetchPrice(i, "USD")
				if err != nil {
					fmt.Printf("Error in cmc price %s\n", err.Error())
					<-time.After(sleepy)
					continue
				}
				cmcResponseCNY, err := fetchPrice(i, "CNY")
				if err != nil {
					fmt.Printf("Error in cmc price %s\n", err.Error())
					<-time.After(sleepy)
					continue
				}
				cmcResponseBTC, err := fetchPrice(i, "BTC")
				if err != nil {
					fmt.Printf("Error in cmc price %s\n", err.Error())
					<-time.After(sleepy)
					continue
				}
				cmcResponseBGX, err := fetchBGXPrice()
				if err != nil {
					fmt.Printf("Error in bgx price %s\n", err.Error())
				}
				err = saveToDb(cmcResponseUSD, cmcResponseCNY, cmcResponseBTC, cmcResponseBGX)
				if err != nil {
					fmt.Printf("Error in cmc price %s\n", err.Error())
					<-time.After(sleepy)
					continue
				}
				<-time.After(sleepy)
			}
		}()
		go func() {
			for {
				<-time.After(time.Second * 10)
				tx, err := dbaforela.Begin()
				if err != nil {
					fmt.Printf("Error fetching ela price from hbg: %s\n", err.Error())
					tx.Rollback()
					continue
				}
				btcPrice, err := getPriceFromHbg()
				if err != nil {
					tx.Rollback()
					fmt.Printf("Error fetching ela price from hbg: %s\n", err.Error())
					continue
				}
				_, err = tx.Exec("update chain_cmc_price set price_btc = '" + btcPrice + "' where symbol = 'ELA' order by _id desc limit 1")
				if err != nil {
					tx.Rollback()
					fmt.Printf("Error fetching ela price from hbg 111 : %s\n", err.Error())
					continue
				}
				tx.Commit()
			}
		}()
	}
}

type hbg_price struct {
	Status string
	Ch     string
	Ts     int64
	Data   []hbg_price_data
}

type hbg_price_data struct {
	Id   int64
	Ts   int64
	Data []hg_price_data_data
}

type hg_price_data_data struct {
	Amount    float64
	Ts        int64
	Id        float64
	Price     float64
	Direction string
}

func getPriceFromHbg() (string, error) {
	resp, err := http.Get(HBG_ENDPOINT_URL)
	if err != nil {
		fmt.Printf("Error fetching price from hbg\n")
		return "", err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var hbg_price hbg_price
		err = json.Unmarshal(body, &hbg_price)
		if err != nil {
			return "", err
		}
		if len(hbg_price.Data) > 0 && len(hbg_price.Data[0].Data) > 0 {
			return strconv.FormatFloat(hbg_price.Data[0].Data[0].Price, 'f', 8, 64), nil
		}
		return "", errors.New("Error fetching price from hbg, data structure is changed")
	}
}

func saveToDb(cmcResponseUSD, cmcResponseCNY, cmcResponseBTC, cmcResponseBGX CmcResponse) error {
	tx, err := dba.Begin()
	if err != nil {
		return err
	}
	data := cmcResponseUSD.Data
	if len(cmcResponseCNY.Data) != len(cmcResponseUSD.Data) || len(cmcResponseUSD.Data) != len(cmcResponseBTC.Data) {
		fmt.Printf("Invalid Key fetch Cmc Data m CNY :%v, BTC :%v, USD :%v", cmcResponseCNY, cmcResponseBTC, cmcResponseUSD)
		return nil
	}
	tx.Exec("delete from chain_cmc_price")
	for i := 0; i < len(data); i++ {
		var btcPrice string
		if data[i].Symbol == "ELA" {
			btcPrice, err = getPriceFromHbg()
			if err != nil {
				dba.Rollback(tx)
				return err
			}
			fmt.Printf("Getting Price From Hbg " + btcPrice + "\n")
		} else {
			btcPrice = strconv.FormatFloat(cmcResponseBTC.Data[i].Quote.BTC.Price, 'f', 8, 64)
		}
		_, err = tx.Exec("insert into chain_cmc_price(id,name,symbol,`rank`,price_usd,price_cny,price_btc,24h_volume_usd,market_cap_usd,available_supply,total_supply,max_supply,percent_change_1h,percent_change_24h,percent_change_7d,last_updated,24h_volume_btc,market_cap_btc,local_system_time,24h_volume_cny,market_cap_cny,platform_symbol,platform_token_address,num_market_pairs) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			strconv.Itoa(int(data[i].Id)),
			data[i].Name,
			data[i].Symbol,
			strconv.Itoa(data[i].Cmc_rank),
			strconv.FormatFloat(data[i].Quote.USD.Price, 'f', 8, 64),
			strconv.FormatFloat(cmcResponseCNY.Data[i].Quote.CNY.Price, 'f', 8, 64),
			btcPrice,
			strconv.FormatFloat(data[i].Quote.USD.Volume_24h, 'f', 8, 64),
			strconv.FormatFloat(data[i].Quote.USD.Market_cap, 'f', 8, 64),
			strconv.FormatFloat(data[i].Circulating_supply, 'f', 8, 64),
			strconv.FormatFloat(data[i].Total_supply, 'f', 8, 64),
			strconv.FormatFloat(data[i].Max_supply, 'f', 8, 64),
			strconv.FormatFloat(data[i].Quote.USD.Percent_change_1h, 'f', 8, 64),
			strconv.FormatFloat(data[i].Quote.USD.Percent_change_24h, 'f', 8, 64),
			strconv.FormatFloat(data[i].Quote.USD.Percent_change_7d, 'f', 8, 64),
			data[i].Quote.USD.Last_updated,
			strconv.FormatFloat(cmcResponseBTC.Data[i].Quote.BTC.Volume_24h, 'f', 8, 64),
			strconv.FormatFloat(cmcResponseBTC.Data[i].Quote.BTC.Market_cap, 'f', 8, 64),
			time.Now(),
			strconv.FormatFloat(cmcResponseCNY.Data[i].Quote.CNY.Volume_24h, 'f', 8, 64),
			strconv.FormatFloat(cmcResponseCNY.Data[i].Quote.CNY.Market_cap, 'f', 8, 64),
			data[i].Platform.Symbol,
			data[i].Platform.Token_Address,
			data[i].Num_market_pairs)
		if err != nil {
			dba.Rollback(tx)
			return err
		}
		// put price that not in the cmc at rank 100
		if i == 99 && len(cmcResponseBGX.Data) > 0 {
			_, err = tx.Exec("insert into chain_cmc_price(id,name,symbol,`rank`,price_usd,price_cny,price_btc,24h_volume_usd,market_cap_usd,available_supply,total_supply,max_supply,percent_change_1h,percent_change_24h,percent_change_7d,last_updated,24h_volume_btc,market_cap_btc,local_system_time,24h_volume_cny,market_cap_cny,platform_symbol,platform_token_address,num_market_pairs) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
				strconv.Itoa(int(cmcResponseBGX.Data[0].Id)),
				cmcResponseBGX.Data[0].Name,
				cmcResponseBGX.Data[0].Symbol,
				strconv.Itoa(cmcResponseBGX.Data[0].Cmc_rank),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.USD.Price, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.CNY.Price, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.BTC.Price, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.USD.Volume_24h, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.USD.Market_cap, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Circulating_supply, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Total_supply, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Max_supply, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.USD.Percent_change_1h, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.USD.Percent_change_24h, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.USD.Percent_change_7d, 'f', 8, 64),
				data[i].Quote.USD.Last_updated,
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.BTC.Volume_24h, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.BTC.Market_cap, 'f', 8, 64),
				time.Now(),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.CNY.Volume_24h, 'f', 8, 64),
				strconv.FormatFloat(cmcResponseBGX.Data[0].Quote.CNY.Market_cap, 'f', 8, 64),
				cmcResponseBGX.Data[0].Platform.Symbol,
				cmcResponseBGX.Data[0].Platform.Token_Address,
				cmcResponseBGX.Data[0].Num_market_pairs)
		}
		if err != nil {
			dba.Rollback(tx)
			return err
		}
	}
	dba.Commit(tx)
	return nil
}

func fetchBGXPrice() (CmcResponse, error) {
	url := fmt.Sprintf(BGX_ENDPOINT_URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CmcResponse{}, err
	}
	req.Header["Accept-Language"] = []string{"*"}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CmcResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CmcResponse{}, err
	}
	id := 5000
	name := "BIT GAME EXCHANGE"
	symbol := "BGX"
	slug := "BIG GAME"
	circulating_supply := 2000000000
	total_supply := 5000000000
	max_supply := 0
	date_added := "2019-02-27T13:53:00.000Z"
	num_market_pairs := 1
	platform_symbol := "ETH"
	platform_token_address := "0xbf3f09e4eba5f7805e5fac0ee09fd6ee8eebe4cb"
	bgxRespMap := new(map[string]interface{})
	err = json.Unmarshal(body, &bgxRespMap)
	if err != nil {
		return CmcResponse{}, err
	}
	data, ok := (*bgxRespMap)["data"].(map[string]interface{})
	if !ok {
		return CmcResponse{}, errors.New("BGX Price Error")
	}
	rate, ok := data["rate"].(map[string]interface{})
	if !ok {
		return CmcResponse{}, errors.New("BGX Price Error")
	}
	enUS, ok := rate["en_US"].(map[string]interface{})
	if !ok {
		return CmcResponse{}, errors.New("BGX Price Error")
	}
	zhCN, ok := rate["zh_CN"].(map[string]interface{})
	if !ok {
		return CmcResponse{}, errors.New("BGX Price Error")
	}
	bgxUs, ok1 := enUS["BGX"].(float64)
	btcUs, ok2 := enUS["BTC"].(float64)
	bgxCn, ok3 := zhCN["BGX"].(float64)
	if !(ok1 && ok2 && ok3) {
		return CmcResponse{}, errors.New("BGX Price Error")
	}
	bgxbtc := math.Round(bgxUs/btcUs*100000000) / 100000000
	now := time.Now().Format("2006-01-02T15:04:05.000Z")
	return CmcResponse{
		Status: Status{
			Timestamp:     now,
			Error_code:    0,
			Error_message: "",
			Elapsed:       0,
			Credit_count:  0,
		},
		Data: []Data{
			Data{
				Id:                 int64(id),
				Name:               name,
				Symbol:             symbol,
				Slug:               slug,
				Circulating_supply: float64(circulating_supply),
				Total_supply:       float64(total_supply),
				Max_supply:         float64(max_supply),
				Date_added:         date_added,
				Num_market_pairs:   int64(num_market_pairs),
				Tags:               nil,
				Platform: Plateform{
					Symbol:        platform_symbol,
					Token_Address: platform_token_address,
				},
				Cmc_rank:     0,
				Last_updated: date_added,
				Quote: Quote{
					CNY: Price{
						Price:              bgxCn,
						Volume_24h:         float64(0),
						Percent_change_1h:  float64(0),
						Percent_change_24h: float64(0),
						Percent_change_7d:  float64(0),
						Market_cap:         float64(0),
						Last_updated:       now,
					},
					USD: Price{
						Price:              bgxUs,
						Volume_24h:         float64(0),
						Percent_change_1h:  float64(0),
						Percent_change_24h: float64(0),
						Percent_change_7d:  float64(0),
						Market_cap:         float64(0),
						Last_updated:       now,
					},
					BTC: Price{
						Price:              bgxbtc,
						Volume_24h:         float64(0),
						Percent_change_1h:  float64(0),
						Percent_change_24h: float64(0),
						Percent_change_7d:  float64(0),
						Market_cap:         float64(0),
						Last_updated:       now,
					},
				},
			},
		},
	}, nil
}

func fetchPrice(i int, curr string) (CmcResponse, error) {
	url := fmt.Sprintf(CMC_ENDPOINT_URL, config.Conf.Cmc.NumOfCoin, curr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CmcResponse{}, err
	}
	println(config.Conf.Cmc.ApiKey[i])
	req.Header = map[string][]string{
		"X-CMC_PRO_API_KEY": []string{config.Conf.Cmc.ApiKey[i]},
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CmcResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CmcResponse{}, err
	}
	cmcResp := CmcResponse{}
	err = json.Unmarshal(body, &cmcResp)
	if err != nil {
		return CmcResponse{}, err
	}
	return cmcResp, nil
}
