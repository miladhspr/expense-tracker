package expense

type Expense struct {
	ID        int64   `json:"id"`
	Amount    float64 `json:"amount"`
	Desc      string  `json:"desc"`
	CreatedAt string  `json:"created_at"`
}
