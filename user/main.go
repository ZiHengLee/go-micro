package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"go-micro/user/domain/repository"
	service2 "go-micro/user/domain/service"
	"go-micro/user/handler"
	user "go-micro/user/proto/user"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql", "root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//禁止副表民称
	db.SingularTable(true)

	//只执行一次，数据表初始化
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()
	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	//注册handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}
	// Register handler

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
