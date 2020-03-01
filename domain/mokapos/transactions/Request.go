package transactions

type Request struct {
	OutletId     int64
	PerPage      int64
	Since        int64
	Until        int64
	TimeFilter   string
	IncludePromo bool
	ReorderType  string
}
