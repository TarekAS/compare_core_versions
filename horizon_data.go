package main

import (
	"strings"
	"time"
)

// JSON Schema of Horizon response.
type HorizonData struct {
	Links struct {
		Account struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"account"`
		Accounts struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"accounts"`
		AccountTransactions struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"account_transactions"`
		ClaimableBalances struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"claimable_balances"`
		Assets struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"assets"`
		Effects struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"effects"`
		FeeStats struct {
			Href string `json:"href"`
		} `json:"fee_stats"`
		Ledger struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"ledger"`
		Ledgers struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"ledgers"`
		Offer struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"offer"`
		Offers struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"offers"`
		Operation struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"operation"`
		Operations struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"operations"`
		OrderBook struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"order_book"`
		Payments struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"payments"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		StrictReceivePaths struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"strict_receive_paths"`
		StrictSendPaths struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"strict_send_paths"`
		TradeAggregations struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"trade_aggregations"`
		Trades struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"trades"`
		Transaction struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"transaction"`
		Transactions struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"transactions"`
	} `json:"_links"`
	HorizonVersion               string    `json:"horizon_version"`
	CoreVersion                  string    `json:"core_version"`
	IngestLatestLedger           int       `json:"ingest_latest_ledger"`
	HistoryLatestLedger          int       `json:"history_latest_ledger"`
	HistoryLatestLedgerClosedAt  time.Time `json:"history_latest_ledger_closed_at"`
	HistoryElderLedger           int       `json:"history_elder_ledger"`
	CoreLatestLedger             int       `json:"core_latest_ledger"`
	NetworkPassphrase            string    `json:"network_passphrase"`
	CurrentProtocolVersion       int       `json:"current_protocol_version"`
	CoreSupportedProtocolVersion int       `json:"core_supported_protocol_version"`
}

// Returns the Semantic Version.
func (h *HorizonData) CoreSemVer() string {
	return strings.Fields(h.CoreVersion)[1]
}

// Returns the Hash Version.
func (h *HorizonData) CoreVersionHash() string {
	return strings.Trim(strings.Fields(h.CoreVersion)[2], "()")
}
