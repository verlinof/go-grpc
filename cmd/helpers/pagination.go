package helpers

import (
	"math"

	paginationPb "github.com/verlinof/go-grpc/pb/pagination"
	"gorm.io/gorm"
)

func Pagination(sql *gorm.DB, page int64, pagination *paginationPb.Pagination) (int64, int64) {
	var total int64
	var limit int64 = 1
	var offset int64

	if page == 0 {
		page = 1
	}

	// Pagination
	sql.Count(&total)
	offset = (page - 1) * limit

	pagination.Total = uint64(total)
	pagination.PerPage = uint32(limit)
	pagination.CurrentPage = uint32(page)
	pagination.LastPage = uint32(math.Ceil(float64(total) / float64(limit)))

	return offset, limit
}
