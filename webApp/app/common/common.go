package common

import (
	"net/http"
	"time"
)

type EventSubData struct {
	Subscription SubscriptionData `json:"subscription"`
	Event        EventData        `json:"event"`
}

type ChallangeData struct {
	Challenge    string           `json:"challenge"`
	Subscription SubscriptionData `json:"subscription"`
}

type SubscribeData struct {
	Type      string        `json:"type"`
	Version   string        `json:"version"`
	Condition ConditionData `json:"condition"`
	Transport TransportData `json:"transport"`
}

type SubscribeResponseData struct {
	Data         []SubRespData `json:"data"`
	Total        int           `json:"total"`
	TotalCost    int           `json:"total_cost"`
	MaxTotalCost int           `json:"max_total_cost"`
}

type EventData struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type SubscriptionData struct {
	ID        string        `json:"id"`
	Status    string        `json:"status"`
	Type      string        `json:"type"`
	Version   string        `json:"version"`
	Cost      int           `json:"cost"`
	Condition ConditionData `json:"condition"`
	Transport TransportData `json:"transport"`
	CreatedAt time.Time     `json:"created_at"`
}

type ConditionData struct {
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

type TransportData struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type SubRespData struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	Version   string `json:"version"`
	Cost      int    `json:"cost"`
	Condition struct {
		BroadcasterUserID string `json:"broadcaster_user_id"`
	} `json:"condition"`
	Transport struct {
		Method   string `json:"method"`
		Callback string `json:"callback"`
	} `json:"transport"`
	CreatedAt time.Time `json:"created_at"`
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func ContainsInt(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
