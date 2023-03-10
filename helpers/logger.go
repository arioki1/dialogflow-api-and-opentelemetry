package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"moul.io/http2curl"
	"net/http"
	"runtime"
)

type AppLog struct {
	Severity string                  `json:"severity"`
	Message  string                  `json:"message"`
	Ctx      *map[string]interface{} `json:"ctx,omitempty"`
}

func (l *AppLog) PrintLog() {
	_, fn, line, _ := runtime.Caller(2)
	location := fmt.Sprintf("%s:%d", fn, line)
	b, err := json.Marshal(l)
	if err != nil {
		return
	}

	log.Printf("%s %s", location, string(b))
}

func (l *AppLog) PrintFatalLog() {
	b, err := json.Marshal(l)
	if err != nil {
		return
	}

	log.Fatal(string(b))
}

func PrintFatalStringLog(msg string) {
	appLog := AppLog{
		Severity: "fatal",
		Message:  msg,
	}

	appLog.PrintFatalLog()
}

func PrintInfoStringLog(msg string) {
	appLog := AppLog{
		Severity: "info",
		Message:  msg,
	}

	appLog.PrintLog()
}

func PrintErrStringLog(msg string) {
	appLog := AppLog{
		Severity: "error",
		Message:  msg,
	}

	appLog.PrintLog()
}

func PrintHttpResponseLog(req *http.Request, resp *http.Response) {
	if req != nil {
		curl, _ := http2curl.GetCurlCommand(req)
		request := AppLog{
			Severity: "request",
			Message:  curl.String(),
		}
		request.PrintLog()
	}

	if resp != nil {
		result, _ := ioutil.ReadAll(resp.Body)

		response := AppLog{
			Severity: "response",
			Message:  string(result),
		}
		response.PrintLog()
	}
}
