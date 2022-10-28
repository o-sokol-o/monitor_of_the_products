package restclient

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/o-sokol-o/as/pkg/domain"
	"github.com/o-sokol-o/as/pkg/restclient"
)

func RestClientGet(url string) ([]domain.Product, error) {
	restClient, err := restclient.NewClient(time.Second * 10)
	if err != nil {
		return nil, err
	}

	// Запрашиваем CSV у внешнего сервиса по REST
	in_csv, err := restClient.Get(url)
	if err != nil {
		return nil, err
	}

	// Получаем список продуктов
	reader := csv.NewReader(strings.NewReader(string(in_csv)))
	reader.Comma = ';'
	products := make(domain.Products, 0)
	for {
		lines, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(lines) != 2 {
			continue
		}

		cost, err := strconv.Atoi(lines[1])
		if err != nil {
			cost = 0
		}
		if lines[0] != "" && cost != 0 {
			products = append(products, domain.Product{
				Name:        lines[0],
				Cost:        int32(cost),
				UpdateAt:    time.Now(),
				ChangeCount: 0,
			})
		}
	}

	return products, nil
}
