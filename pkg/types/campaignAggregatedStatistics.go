package types

type CampaignAggregatedStatistics struct {
	Users     int64 `json:"Users,omitempty"`
	Delivered int64 `json:"Delivered,omitempty"`
	Errors    int64 `json:"Errors,omitempty"`
	Pending   int64 `json:"Pending,omitempty"`
	TimedOut  int64 `json:"TimedOut,omitempty"`
}
