package tx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Transactions struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		LastItemIndex int `json:"last_item_index"`
		Pi            struct {
			Balance              int64 `json:"balance"`
			CurrentHeight        int   `json:"curent_height"`
			TransferEntriesCount int   `json:"transfer_entries_count"`
			TransfersCount       int   `json:"transfers_count"`
			UnlockedBalance      int64 `json:"unlocked_balance"`
		} `json:"pi"`
		TotalTransfers int `json:"total_transfers"`
		Transfers      []struct {
			Amount          int64    `json:"amount"`
			Comment         string   `json:"comment"`
			Fee             int64    `json:"fee"`
			Height          int      `json:"height"`
			IsIncome        bool     `json:"is_income"`
			IsMining        bool     `json:"is_mining"`
			IsMixing        bool     `json:"is_mixing"`
			IsService       bool     `json:"is_service"`
			PaymentId       string   `json:"payment_id"`
			RemoteAddresses []string `json:"remote_addresses,omitempty"`
			RemoteAliases   []string `json:"remote_aliases,omitempty"`
			ShowSender      bool     `json:"show_sender"`
			Td              struct {
				Spn []int64 `json:"spn,omitempty"`
				Rcv []int64 `json:"rcv,omitempty"`
			} `json:"td"`
			Timestamp             int    `json:"timestamp"`
			TransferInternalIndex int    `json:"transfer_internal_index"`
			TxBlobSize            int    `json:"tx_blob_size"`
			TxHash                string `json:"tx_hash"`
			TxType                int    `json:"tx_type"`
			UnlockTime            int    `json:"unlock_time"`
			ServiceEntries        []struct {
				Body        string `json:"body"`
				Flags       int    `json:"flags"`
				Instruction string `json:"instruction"`
				ServiceId   string `json:"service_id"`
			} `json:"service_entries,omitempty"`
		} `json:"transfers"`
	} `json:"result"`
}

func GetTxs(url string) {
	//walletUrl := url
	jsonBody := fmt.Sprintln(`{
    "jsonrpc": "2.0",
    "id": 0,
    "method": "get_recent_txs_and_info",
    "params": {
      "offset": 0,
      "update_provision_info": true,
      "exclude_mining_txs": true,
      "count": 50,
      "order": "FROM_END_TO_BEGIN",
      "exclude_unconfirmed": true
    }
  }`)
	request, err := http.NewRequest("POST", "http://localhost:11212/json_rpc", bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		fmt.Println("error") // return meaningful statement
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("error") // return meaningful statement
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	data := Transactions{}

	_ = json.Unmarshal([]byte(body), &data)

	for _, tx := range data.Result.Transfers {
		fmt.Println(tx.Amount)
	}
}
