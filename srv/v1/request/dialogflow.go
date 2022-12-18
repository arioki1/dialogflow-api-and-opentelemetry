package request

import (
	"encoding/json"
	"strings"
)

type DialogflowRequest struct {
	ResponseId  string `json:"responseId"`
	QueryResult struct {
		QueryText                string     `json:"queryText"`
		Action                   string     `json:"action"`
		Parameters               Parameters `json:"parameters,omitempty"`
		AllRequiredParamsPresent bool       `json:"allRequiredParamsPresent"`
		FulfillmentText          string     `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		} `json:"fulfillmentMessages"`
		OutputContexts []OutputContexts `json:"outputContexts"`
		Intent         struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"intent"`
		IntentDetectionConfidence json.Number `json:"intentDetectionConfidence"`
		LanguageCode              string      `json:"languageCode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Source  string `json:"source"`
		Payload struct {
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	Session string `json:"session"`
}

type Parameters struct {
	ID *string `json:"id,omitempty"`
}

type OutputContexts struct {
	Name          string      `json:"name"`
	LifespanCount int         `json:"lifespanCount,omitempty"`
	Parameters    interface{} `json:"parameters"`
}

func (d *DialogflowRequest) GetSessionId() string {
	s := strings.Split(d.Session, "/")
	if len(s) > 0 {
		return s[len(s)-1]
	} else {
		return ""
	}
}

func (d *DialogflowRequest) GetContext(name string) *OutputContexts {
	for _, c := range d.QueryResult.OutputContexts {
		s := strings.Split(c.Name, "/")
		if len(s) > 0 {
			if s[len(s)-1] == name {
				return &c
			}
		}
	}
	return nil
}