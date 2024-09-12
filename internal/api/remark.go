package api

type Remark struct {
	Type    string `json:"type"`
	Summary string `json:"summary"`
	Code    string `json:"code"`
	Text    string `json:"text"`
}
