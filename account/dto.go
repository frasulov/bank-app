package account

type CreateAccountInput struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance" validate:"required"`
	Currency string `json:"currency" validate:"oneof=AZN USD"`
}

type AccountOutput struct {
	Id       int64  `json:"id"`
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type ListAccountParam struct {
	PageId   int32 `query:"page_id"`
	PageSize int32 `query:"page_size"`
}

func (l *ListAccountParam) setDefaults() {
	if l.PageSize == 0 {
		l.PageSize = 5
	}
	if l.PageId == 0 {
		l.PageId = 1
	}
}
