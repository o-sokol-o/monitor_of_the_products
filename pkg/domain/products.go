package domain

import (
	"time"
)

type LogItem struct {
	RequestType string    `bson:"request_type"`
	Time        time.Time `bson:"timestamp"`
	Input       string    `bson:"input"`
}

type GrpcRequestFetch struct {
	Url string `bson:"url"`
}

const (
	Collection_Products string = "products"
	Collection_Logs     string = "logs"

	Sort_Name         int32 = 1
	Sort_Cost         int32 = 2
	Sort_Update_At    int32 = 3
	Sort_Change_Count int32 = 4

	SortAsc  int32 = 1 // по возрастанию
	SortNone int32 = 0
	SortDesc int32 = -1 // по убыванию
)

func GetSortFieldName(num_field int32) string {
	switch num_field {
	case 2:
		return "cost"
	case 3:
		return "update_at"
	case 4:
		return "change_count"
	default:
		return "name"
	}
}

type Paging struct {
	NumPage        int32 `bson:"num_page"`
	ProductPerPage int32 `bson:"product_per_page"`
}
type Sorting struct {
	SortField int32 `bson:"field"`
	Direction int32 `bson:"direction"`
}
type GrpcRequestList struct {
	Paging  Paging
	Sorting Sorting
}

type Product struct {
	Name        string    `bson:"name" json:"name"`
	Cost        int32     `bson:"cost" json:"cost"`
	UpdateAt    time.Time `bson:"update_at" json:"update_at"`
	ChangeCount int32     `bson:"change_count" json:"change_count"`
}

type Products []Product
