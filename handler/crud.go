package handler

import (
	"context"

	proto "grpcProject/proto/crud"
	"grpcProject/repository"

	"github.com/micro/micro/v3/service/errors"
	"github.com/sirupsen/logrus"
)

type Crud struct {
	items *repository.Items
}

func NewCrud(items *repository.Items) *Crud {
	return &Crud{
		items: items,
	}
}

// micro call  Crud.CreateItem '{"name":"test1", "id":"id1"}'

// Not working
// curl \
// -d "service=grpcproject" \
// -d "method=crud.GetItem" \
// -d "request={\"id\": \"John\"}" \
// http://localhost:8080/rpc

// curl http://localhost:8080/grpcproject/crud/getItem?id=joohn
func (crud *Crud) CreateItem(ctx context.Context, emp *proto.Employee, id *proto.ID) error {
	logrus.Infof("Handle create %v", emp)
	tempId := emp.GetId()
	crud.items.Put(tempId, emp)

	id.Id = tempId
	return nil
}

// micro call grpcproject Crud.GetItem '{"id":"id1"}'
func (crud *Crud) GetItem(ctx context.Context, id *proto.ID, emp *proto.Employee) error {
	logrus.Infof("Handle get %v", id)
	tempId := id.GetId()

	res := crud.items.Get(tempId)
	if res == nil {
		return errors.NotFound("handler.crud", "Item not found") // ??? id ???
	}

	*emp = *res
	return nil
}

func (crud *Crud) UpdateItem(ctx context.Context, emp *proto.Employee, id *proto.ID) error {
	logrus.Infof("Handle update %v", emp)
	// todo
	return nil
}

func (crud *Crud) DeleteItem(ctx context.Context, id *proto.ID, resId *proto.ID) error {
	logrus.Infof("Handle delete %v", id)
	tempId := id.GetId()
	crud.items.Detele(tempId)
	return nil
}
