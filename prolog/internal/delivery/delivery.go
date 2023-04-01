package delivery

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

var db *sql.DB

type Identity struct {
	ID int64
}

type SaleOrder struct {
	Identity
}

type Stop struct {
	Identity
	baseMapper      BaseMapper
	tripID          string
	deliveryOrderID int64
	deliveryOrder   *DeliveryOrder
	deliveryItems   []*DeliveryItem
}

func newStop(tripID string, deliveryOrderID int64, deliveryItems []*DeliveryItem) *Stop {
	return &Stop{
		tripID:          tripID,
		deliveryOrderID: deliveryOrderID,
		deliveryItems:   deliveryItems,
		baseMapper: &BaseMapperImpl{
			insertSql: func(o DomainObject) (Statement, error) {
				stop := o.(*Stop)
				// load delivery order
				do := &DeliveryOrder{}
				_, err := do.FindByID(deliveryOrderID)
				if err != nil {
					return Statement{}, fmt.Errorf("Cannot find delivery order by id = %v", deliveryOrderID)
				}
				stop.deliveryOrder = do

				itemsJson, err := json.Marshal(stop.deliveryItems)
				if err != nil {
					return Statement{}, err
				}
				return Statement{
					sql:  "INSERT INTO stop(trip_id, delivery_order_id, delivery_items)",
					args: []any{stop.tripID, stop.deliveryOrderID, itemsJson},
				}, nil
			},
		},
	}
}

func (s *Stop) Create() (*Stop, error) {
	if s.deliveryOrderID == 0 || len(s.deliveryItems) == 0 {
		return nil, fmt.Errorf("invalid stop: cannot create: %v", s)
	}
	_, err := s.baseMapper.Insert(s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Stop) GetID() int64 {
	return s.ID
}

func (s *Stop) SetID(id int64) {
	s.ID = id
}
