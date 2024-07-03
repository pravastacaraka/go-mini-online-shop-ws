package domain

type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Error  string        `json:"error,omitempty"`
}

type PageMetadata struct {
	Page      uint8  `json:"page"`
	Size      uint8  `json:"size"`
	TotalItem uint16 `json:"total_item"`
	TotalPage uint16 `json:"total_page"`
}
