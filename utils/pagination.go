package utils

import "math"

func CountOffset(page int, perPage int) int {
	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	return (page - 1) * perPage
}

func CountTotalPage(count, perPage int) int {
	return int(math.Ceil(float64(count) / float64(perPage)))
}
