package mysql

import (
	"database/sql"
	"github.com/v-sadovsky/ok_test_task/pkg/models"
)

type StoresModel struct {
	DB *sql.DB
}

func (m *StoresModel) Get(tm string) (map[int]*models.Shop, error) {
	stmt := `SELECT stores.*, products.id, products.name, products.description, products.price  FROM stores 
    JOIN products ON store_id=stores.id WHERE ? BETWEEN stores.open_time AND stores.close_time 
                                        ORDER BY stores.id`
	rows, err := m.DB.Query(stmt, tm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shops := make(map[int]*models.Shop)
	for rows.Next() {
		s := &models.Shop{}
		o := &models.Item{}
		err = rows.Scan(&s.ID, &s.Name, &s.URL, &s.WorkingTime.Open, &s.WorkingTime.Close, &o.ID, &o.Name, &o.Description, &o.Price)
		if err != nil {
			return nil, err
		}

		if shop, ok := shops[s.ID]; !ok {
			s.Offers = append(s.Offers, o)
			shops[s.ID] = s
		} else {
			shop.Offers = append(shop.Offers, o)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(shops) == 0 {
		return nil, models.ErrNoRecord
	}

	return shops, nil
}
