package delivery

import (
	"fmt"
)

type DeliveryOrder struct {
	Identity
	baseMapper    BaseMapper
	saleOrderID   string
	deliveryItems []*DeliveryItem
}

func newDeliveryOrder(saleOrderId string, deliveryItems []*DeliveryItem) *DeliveryOrder {
	return &DeliveryOrder{
		saleOrderID:   saleOrderId,
		deliveryItems: deliveryItems,
		baseMapper: &BaseMapperImpl{
			insertSql: func(o DomainObject) (Statement, error) {
				do := o.(*DeliveryOrder)
				return Statement{
					sql:  "INSERT INTO delivery_order (sale_order_id) VALUES (?)",
					args: []any{do.saleOrderID},
				}, nil
			},
			findByIDSql: func(id int64) (Statement, error) {
				return Statement{
					sql:  "SELECT * FROM delivery_order WHERE id = ?",
					args: []any{id},
				}, nil
			},
			findAll: func() (Statement, error) {
				return Statement{
					sql: "SELECT * FROM delivery_order",
				}, nil
			},
			loadRow: func(row Row) (DomainObject, error) {
				do := &DeliveryOrder{}
				if err := row.Scan(&do.ID, &do.saleOrderID); err != nil {
					return nil, err
				}
				return do, nil
			},
		},
	}
}

func (do *DeliveryOrder) GetID() int64 {
	return do.ID
}

func (do *DeliveryOrder) SetID(id int64) {
	do.ID = id
}

func (do *DeliveryOrder) Create() (*DeliveryOrder, error) {
	if do.saleOrderID == "" || len(do.deliveryItems) == 0 {
		return nil, fmt.Errorf("invalid delivery order")
	}
	_, err := do.baseMapper.Insert(do)
	if err != nil {
		return nil, err
	}
	err = do.InsertItems()
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (do *DeliveryOrder) FindByID(id int64) (*DeliveryOrder, error) {
	rs, err := do.baseMapper.FindByID(id)
	if err != nil {
		return nil, err
	}
	do = rs.(*DeliveryOrder)
	return do, nil
}

func (do *DeliveryOrder) FindAll() ([]*DeliveryOrder, error) {
	domainObjs, err := do.baseMapper.FindAll()
	if err != nil {
		return nil, err
	}
	var deliveryOrders []*DeliveryOrder
	for _, o := range domainObjs {
		deliveryOrders = append(deliveryOrders, o.(*DeliveryOrder))
	}
	return deliveryOrders, nil
}

func (do *DeliveryOrder) InsertItems() error {
	for _, item := range do.deliveryItems {
		item.deliveryOrder = do
		_, err := item.baseMapper.Insert(item)
		if err != nil {
			return fmt.Errorf("insertItems: %v", err)
		}
	}
	return nil
}

type DeliveryItem struct {
	Identity
	baseMapper    BaseMapper
	productID     string
	quantity      uint64
	deliveryOrder *DeliveryOrder
}

func newDeliveryItem(productID string, quantity uint64) *DeliveryItem {
	return &DeliveryItem{
		productID: productID,
		quantity:  quantity,
		baseMapper: &BaseMapperImpl{
			insertSql: func(o DomainObject) (Statement, error) {
				item := o.(*DeliveryItem)
				if item.deliveryOrder == nil {
					return Statement{}, fmt.Errorf("Invalid delivery item: missing delivery order")
				}
				return Statement{
					sql:  "INSERT INTO delivery_item (delivery_order_id, product_id, quantity) VALUES (?, ?, ?)",
					args: []any{item.deliveryOrder.ID, item.productID, item.quantity},
				}, nil
			},
		},
	}
}

func (di *DeliveryItem) GetID() int64 {
	return di.ID
}

func (di *DeliveryItem) SetID(id int64) {
	di.ID = id
}
