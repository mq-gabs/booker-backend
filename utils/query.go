package utils

type QueryOptions struct {
	Page     int     `json:"page,omitempty"`
	PageSize int     `json:"page_size,omitempty"`
	Term     *string `json:"term,omitempty"`
}

func NewQuery() *QueryOptions {
	return &QueryOptions{
		PageSize: 10,
	}
}
