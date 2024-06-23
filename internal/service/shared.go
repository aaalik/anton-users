package service

type Order int32

const (
	ORDER_UNSPECIFIED Order = iota
	ORDER_ASC
	ORDER_DESC
)

type Range struct {
	GTE int64 `json:"gte"`
	LTE int64 `json:"lte"`
}

type SortBy struct {
	Field string `json:"field,omitempty"`
	Order Order  `json:"order,omitempty"`
}

type Queries struct {
	Page        int32   `json:"page,omitempty"`
	Rows        int32   `json:"rows,omitempty"`
	Sort        *SortBy `json:"sort,omitempty"`
	Keyword     string  `json:"keyword,omitempty"`
	WithDeleted []bool  `json:"with_deleted,omitempty"`
}
