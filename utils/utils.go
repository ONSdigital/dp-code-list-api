package utils

import (
	"net/url"
	"strconv"

	"github.com/ONSdigital/dp-code-list-api/models"
)

// GetPositiveIntQueryParameter obtains the positive int value of query var defined by the provided varKey
func GetPositiveIntQueryParameter(queryVars url.Values, varKey string, defaultValue int) (val int, err error) {
	strVal, found := queryVars[varKey]
	if !found {
		return defaultValue, nil
	}
	val, err = strconv.Atoi(strVal[0])
	if err != nil {
		return -1, err
	}
	if val < 0 {
		return 0, nil
	}
	return val, nil
}

func CodelistsSlice(full []models.CodeList, offset, limit int) (sliced []models.CodeList) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []models.CodeList{}
	}
	return full[offset:end]
}

func EditionsSlice(full []models.Edition, offset, limit int) (sliced []models.Edition) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []models.Edition{}
	}
	return full[offset:end]
}

func CodesSlice(full []models.Code, offset, limit int) (sliced []models.Code) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []models.Code{}
	}
	return full[offset:end]
}

func DatasetsSlice(full []models.Dataset, offset, limit int) (sliced []models.Dataset) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []models.Dataset{}
	}
	return full[offset:end]
}
