package models

// Train model
type Train struct {
	TrainNo   string `json:"trainNo"`
	TrainName string `json:"trainName"`
	SEQ       string `json:"seq"`
	Code      string `json:"code"`
	StName    string `json:"stName"`
	ATime     string `json:"atime"`
	DTime     string `json:"dtime"`
	Distance  string `json:"distance"`
	SS        string `json:"ss"`
	SSname    string `json:"ssname"`
	Ds        string `json:"ds"`
	DsName    string `json:"dsName"`
}
