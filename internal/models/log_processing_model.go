package models

type VpcNormalizedData struct {
	Version         int
	AccountID       string
	InterfaceID     string
	SourceIP        string
	DestinationIP   string
	DestinationPort int
	SourcePort      int
	Protocol        int
	Packets         int
	Bytes           int
	StartTime       int
	EndTime         int
	Action          string
	LogStatus       string
}
