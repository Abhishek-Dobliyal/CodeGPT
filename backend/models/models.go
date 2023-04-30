package models

type Request struct {
	CodeText        string  `json:"code"`
	Action          string  `json:"action"`
	ConvertTo       *string `json:"convertTo"`
	AddTest         *bool   `json:"addTest"`
	MultipleFunc    *bool   `json:"multipleFunc"`
	CommentEachLine *bool   `json:"commentEachLine"`
	ConciseComment  *bool   `json:"conciseComment"`
	IndentSize      int     `json:"indentSize"`
	BestPractice    *bool   `json:"bestPractice"`
}

type Response struct {
	Message    string  `json:"message"`
	StatusCode int     `josn:"statusCode"`
	Data       *string `json:"data"`
}

type ChatGPTApiResponse struct {
	Result *string
}
