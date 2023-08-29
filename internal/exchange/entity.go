package exchange

import "time"

type Quotation struct {
	UUID       string
	Code       string
	CodeIn     string
	Name       string
	High       float64
	Low        float64
	VarBid     float64
	PctChange  float64
	Bid        float64
	Ask        float64
	Timestamp  string
	CreateDate time.Time
	CreatedAt  time.Time
}

type Exchange struct {
	CurrentValue float64 `json:"current_value,omitempty"`
}
