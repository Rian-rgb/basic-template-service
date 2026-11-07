package utils

type PaginationParams struct {
	Page   int
	Limit  int
	Offset int
}

func ValidatePagination(page, limit int) PaginationParams {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if limit > 20 {
		limit = 20
	}

	return PaginationParams{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}
