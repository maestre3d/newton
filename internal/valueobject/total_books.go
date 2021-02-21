package valueobject

import (
	"math"

	"github.com/maestre3d/newton/internal/domain"
)

const (
	totalBooksMinValue = 0
)

// ErrTotalBooksOutOfRange the given total books counter length is out of range, use uint64 valid value
var ErrTotalBooksOutOfRange = domain.NewOutOfRange("total_books", totalBooksMinValue, math.MaxInt64)

// TotalBooks Author's total books counter
type TotalBooks uint64

// NewTotalBooks creates and validates an Author total book counter
func NewTotalBooks(v int64) (TotalBooks, error) {
	total := TotalBooks(0) // avoid uint negative overflow at runtime
	if err := total.ensurePositive(v); err != nil {
		return 0, err
	}
	total = TotalBooks(v)
	return total, nil
}

func (b TotalBooks) ensurePositive(v int64) error {
	if v < totalBooksMinValue {
		return ErrTotalBooksOutOfRange
	}
	return nil
}

func (b TotalBooks) Value() uint64 {
	return uint64(b)
}
