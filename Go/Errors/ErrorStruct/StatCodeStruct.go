package errorstruct

import "fmt"

type StatCode struct {
	Code           int
	Message        string
	Reason         string `json:"reason"`
	ReasonsMessage string `json:"message"`
}

func (err *StatCode) Error() string {
	if err.Reason != "" {
		return fmt.Sprintf("[%d %s] %s: %s", err.Code, err.Message, err.Reason, err.ReasonsMessage)
	}
	return fmt.Sprintf("[%d %s]", err.Code, err.Message)
}
