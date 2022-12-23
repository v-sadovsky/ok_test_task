package mock

import (
	"errors"
	"github.com/v-sadovsky/ok_test_task/pkg/models"
)

const (
	NonWorkingTime       = "08:30:00"
	Shop1WorkingTimeOnly = "10:00:00"
	Shop2WorkingTimeOnly = "19:00:00"
	BothShopsWorkingTime = "12:00:00"
)

var (
	ErrUnexpectedInput = errors.New("mock db: unexpected test input")
)

var item1 = &models.Item{
	ID:          1,
	Name:        "Phone",
	Description: "<H1>Samsung S11 Edge </H1><div>Best choice on the market</div>",
	Price:       1257.00,
}

var item2 = &models.Item{
	ID:          2,
	Name:        "Notebook",
	Description: "<H1>Asus ZenBook </H1><div>Stylish modern notebook for productive work</div>",
	Price:       4548.00,
}

var item3 = &models.Item{
	ID:          3,
	Name:        "Monitor",
	Description: "<H1>HP 2W925AA </H1><div>Best choice for working and gaming</div>",
	Price:       1257.00,
}

var item4 = &models.Item{
	ID:          4,
	Name:        "Computer",
	Description: "<H1>MultiGame 5C104FD </H1><div>Reliable high performance package</div>",
	Price:       4548.00,
}

var offers1 = []*models.Item{item1, item2}
var offers2 = []*models.Item{item3, item4}

var schedule1 = &models.Schedule{
	Open:  "09:00:00",
	Close: "18:00:00",
}

var schedule2 = &models.Schedule{
	Open:  "11:30:00",
	Close: "20:30:00",
}

var mockShop1 = &models.Shop{
	ID:          1,
	Name:        "TTN",
	URL:         "ttn.by",
	WorkingTime: *schedule1,
	Offers:      offers1,
}

var mockShop2 = &models.Shop{
	ID:          2,
	Name:        "5 Element",
	URL:         "5element.by",
	WorkingTime: *schedule2,
	Offers:      offers2,
}

type StoresModel struct{}

func (m *StoresModel) Get(tm string) (map[int]*models.Shop, error) {
	shops := make(map[int]*models.Shop)
	var err error

	switch tm {
	case NonWorkingTime:
		return nil, models.ErrNoRecord
	case Shop2WorkingTimeOnly:
		shops[2] = mockShop2
	case Shop1WorkingTimeOnly:
		shops[1] = mockShop1
	case BothShopsWorkingTime:
		shops[1] = mockShop1
		shops[2] = mockShop2
	default:
		err = ErrUnexpectedInput
	}

	return shops, err
}
