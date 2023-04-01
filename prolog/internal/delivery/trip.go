package delivery

import "fmt"

type Trip struct {
	Identity
	baseMapper     BaseMapper
	deliveryOrders []DeliveryOrder
	deliveryItems  []DeliveryItem
	stops          []Stop
}

func newTrip() *Trip {
	return &Trip{
		baseMapper: &BaseMapperImpl{
			insertSql: func(o DomainObject) (Statement, error) {
				return Statement{
					sql:  "INSERT INTO trip VALUES ()",
					args: []any{},
				}, nil
			},
		},
	}
}

func (tr *Trip) Create() (*Trip, error) {
	if tr.ID != 0 {
		return nil, fmt.Errorf("invalid trip: already attached entity, cannot create")
	}
	// create a new trip
	_, err := tr.baseMapper.Insert(tr)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (tr *Trip) AddStop(stop Stop) {
	stop.Create()
}

func (di *Trip) GetID() int64 {
	return di.ID
}

func (di *Trip) SetID(id int64) {
	di.ID = id
}
