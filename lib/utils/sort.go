package utils

import "strings"

func ValidateSort(
	sortBy string,
	sort string,
	allowedColumns map[string]string,
	defaultSortBy string,
	defaultSort string,
) (string, string) {
	sortBy = strings.ToLower(strings.TrimSpace(sortBy))
	sort = strings.ToLower(strings.TrimSpace(sort))

	if _, exists := allowedColumns[sortBy]; !exists {
		sortBy = defaultSortBy
	}

	if sort != "asc" && sort != "desc" {
		sort = defaultSort
	}

	return sortBy, sort
}
