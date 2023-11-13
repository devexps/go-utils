package pagination

const (
	DefaultPage     = 1  // Default number of pages
	DefaultPageSize = 10 // Default number of lines per page
)

// GetPageOffset calculates offset
func GetPageOffset(pageNum, pageSize int32) int {
	return int((pageNum - 1) * pageSize)
}
