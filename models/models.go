package models

// Config model
type Config struct {
	ConnString string `mapstructure:"CONNECTION_STRING"`
	DbName     string `mapstructure:"DATABASE_NAME"`
	Collection string `mapstructure:"COLLECTION"`
}

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
	TimeDiff  int
}
