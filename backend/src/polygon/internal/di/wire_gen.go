// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/dao"
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/server/grpc"
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/server/http"
	"github.com/wejudge/wejudge-polygon/src/polygon/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	redis, cleanup, err := dao.NewRedis()
	if err != nil {
		return nil, nil, err
	}
	memcache, cleanup2, err := dao.NewMC()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	db, cleanup3, err := dao.NewDB()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	daoDao, cleanup4, err := dao.New(redis, memcache, db)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceService, cleanup5, err := service.New(daoDao)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	engine, err := http.New(serviceService)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, err := grpc.New(serviceService)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup6, err := NewApp(serviceService, engine, server)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
