package utils

import (
	"strconv"
)

type QueryParams struct {
	Filters   map[string]string
	Page      int64
	Limit     int64
	Field     string
	Value     string
	SortBy    string
	StartedAt string
	EndedAt   string
}

func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		Filters:   make(map[string]string),
		Page:      1,
		Limit:     10,
		Field:     "",
		Value:     "",
		SortBy:    "",
		StartedAt: "",
		EndedAt:   "",
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `page` param")
			}
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}
		if key == "field" {
			params.Field = value[0]
			continue
		}
		if key == "value" {
			params.Value = value[0]
			continue
		}
		if key == "sort_by" {
			params.SortBy = value[0]
			continue
		}
		if key == "started_at" {
			params.StartedAt = value[0]
			continue
		}
		if key == "ended_at" {
			params.EndedAt = value[0]
			continue
		}

		params.Filters[key] = value[0]
	}

	return &params, errStr
}
