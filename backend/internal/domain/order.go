package domain

import (
	"time"
)

type OrderDirection string

const (
	DirectionCall OrderDirection = "CALL"
	DirectionPut  OrderDirection = "PUT"
)

type TradeOrder struct {
	ID            string         `json:"id"`
	UserID        string         `json:"user_id"`
	AssetSymbol   string         `json:"asset_symbol"`
	Direction     OrderDirection `json:"direction"`
	StakeAmount   float64        `json:"stake_amount"`
	StrikePrice   float64        `json:"strike_price"`
	PayoutRate    float64        `json:"payout_rate"`
	DurationSec   int            `json:"duration_sec"`
	CreatedAt     time.Time      `json:"created_at"`
	SettlesAt     time.Time      `json:"settles_at"`
}

type ExecutionEngine struct {
	status string
}

func NewExecutionEngine() *ExecutionEngine {
	return &ExecutionEngine{status: "OPERATIONAL"}
}

func (e *ExecutionEngine) GetStatus() string {
	return e.status
}
