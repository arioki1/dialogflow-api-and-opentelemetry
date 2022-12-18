package response

type DialogFlowResponseGenerator struct {
	QueryResult *struct {
		Parameters map[string]interface{} `json:"parameters"`
	} `json:"queryResult,omitempty"`

	FulfillmentText *string `json:"fulfillmentText,omitempty"`

	FulfillmentMessages []*struct {
		Text struct {
			Text []string `json:"text,omitempty"`
		} `json:"text,omitempty"`
	} `json:"fulfillmentMessages,omitempty"`

	FollowupEventInput *struct {
		Name         string                  `json:"name,omitempty"`
		Parameters   *map[string]interface{} `json:"parameters,omitempty"`
		LanguageCode string                  `json:"languageCode,omitempty"`
	} `json:"followupEventInput,omitempty"`

	OutputContexts []*struct {
		Name          string      `json:"name,omitempty"`
		LifespanCount int         `json:"lifespanCount,omitempty"`
		Parameters    interface{} `json:"parameters,omitempty"`
	} `json:"outputContexts,omitempty"`
}

func (d *DialogFlowResponseGenerator) AddQueryResultParameters(id string, value interface{}) {
	if d.QueryResult == nil {
		d.QueryResult = &struct {
			Parameters map[string]interface{} `json:"parameters"`
		}{}
	}
	if d.QueryResult.Parameters == nil {
		d.QueryResult.Parameters = map[string]interface{}{
			id: value,
		}
	} else {
		d.QueryResult.Parameters[id] = value
	}
}
func (d *DialogFlowResponseGenerator) SetFulfillmentText(text string) {
	d.FulfillmentText = &text
}
func (d *DialogFlowResponseGenerator) SetFulfillmentMessages(text string) {
	d.FulfillmentMessages = []*struct {
		Text struct {
			Text []string `json:"text,omitempty"`
		} `json:"text,omitempty"`
	}{
		{
			Text: struct {
				Text []string `json:"text,omitempty"`
			}{
				Text: []string{text},
			},
		},
	}

}
func (d *DialogFlowResponseGenerator) SetFollowupEventInput(name string, parameters *map[string]interface{}, languageCode *string) {
	l := "en-US"
	if languageCode != nil {
		l = *languageCode
	}
	d.FollowupEventInput = &struct {
		Name         string                  `json:"name,omitempty"`
		Parameters   *map[string]interface{} `json:"parameters,omitempty"`
		LanguageCode string                  `json:"languageCode,omitempty"`
	}{
		Name:         name,
		Parameters:   parameters,
		LanguageCode: l,
	}
}
func (d *DialogFlowResponseGenerator) AddOutputContexts(sessionId string, name string, lifespanCount int, parameters *map[string]string) {
	outputContexts := struct {
		Name          string      `json:"name,omitempty"`
		LifespanCount int         `json:"lifespanCount,omitempty"`
		Parameters    interface{} `json:"parameters,omitempty"`
	}{
		Name:          sessionId + "/contexts/" + name,
		LifespanCount: lifespanCount,
		Parameters:    parameters,
	}
	d.OutputContexts = append(d.OutputContexts, &outputContexts)
}

/*
example to use
func main() {
	response := DialogFlowResponse{}
	response.SetFulfillmentText("123")
	response.SetFulfillmentMessages("123")
	params := &map[string]interface{}{"123": "123"}
	response.SetFollowupEventInput("123", params, nil)
	p := &map[string]string{"123": "123"}
	response.AddOutputContexts("123", "123", 123, p)
	response.AddOutputContexts("123", "123", 123, p)
	response.AddQueryResultParameters("123", "123")
	response.AddQueryResultParameters("12322", 123)
	marshal, _ := json.Marshal(response)
	fmt.Println(string(marshal))
}
*/
