package delivery

import (
	"database/sql"
	"fmt"
)

type DomainObject interface {
	GetID() int64
	SetID(int64)
}

type Statement struct {
	sql  string
	args []any
}

type BaseMapper interface {
	Insert(DomainObject) (DomainObject, error)
	FindByID(int64) (DomainObject, error)
	FindAll() ([]DomainObject, error)
}

type Row interface {
	Scan(dest ...any) error
}

type BaseMapperImpl struct {
	// @param: Domain object
	// @return: [sql string, args]
	insertSql   func(DomainObject) (Statement, error)
	findByIDSql func(int64) (Statement, error)
	findAll     func() (Statement, error)
	loadRow     func(Row) (DomainObject, error)
}

func (b *BaseMapperImpl) Insert(o DomainObject) (DomainObject, error) {
	stat, err := b.insertSql(o)
	if err != nil {
		return nil, err
	}
	result, err := db.Exec(stat.sql, stat.args...)
	if err != nil {
		return o, fmt.Errorf("Insert error: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return o, fmt.Errorf("Get LastInsertID error: %v", err)
	}
	o.SetID(id)
	return o, nil
}

func (b *BaseMapperImpl) FindByID(id int64) (DomainObject, error) {
	stat, err := b.findByIDSql(id)
	if err != nil {
		return nil, err
	}
	row := db.QueryRow(stat.sql, stat.args...)
	domainObj, err := b.loadRow(row)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Cannot find with id %v", id)
	}
	if err != nil {
		return nil, fmt.Errorf("FindByID error, id=%v", id)
	}
	if domainObj == nil {
		return nil, fmt.Errorf("")
	}
	return domainObj, nil
}

func (b *BaseMapperImpl) FindAll() ([]DomainObject, error) {
	stat, err := b.findAll()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(stat.sql, stat.args...)
	if err != nil {
		return nil, fmt.Errorf("FindAll error")
	}
	defer rows.Close()

	var objs []DomainObject
	for rows.Next() {
		obj, err := b.loadRow(rows)
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}
