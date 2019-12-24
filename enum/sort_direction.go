package enum

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func (this SortDirection) String() string {
	switch this {
	case SortDirectionAsc:
		return "asc"
	case SortDirectionDesc:
		return "desc"
	default:
		return "unknown"
	}
}

func (this SortDirection) Valid() bool {
	if this.String() == "unknown" {
		return false
	}
	return true
}
