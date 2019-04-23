package main

import (
	"fmt"
	configUtils "github.com/sundogrd/gopkg/config"
	"github.com/sundogrd/gopkg/db"
	userGen "github.com/sundogrd/user-grpc/grpc_gen/user"
	userRepo "github.com/sundogrd/user-grpc/providers/repos/user/repo"
	"github.com/sundogrd/user-grpc/servers/user"
	userService "github.com/sundogrd/user-grpc/services/user/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func main() {
	config, err := configUtils.ReadConfigFromFile("./config", nil)
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", config.Get("grpcPort").(string))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	// init services and repos
	//authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	//ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	//
	//timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	//au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	//_articleHttpDeliver.NewArticleHttpHandler(e, au)

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           config.Get("db.options.user").(string),
		Password:       config.Get("db.options.password").(string),
		Host:           config.Get("db.options.host").(string),
		Port:           config.Get("db.options.port").(string),
		DBName:         config.Get("db.options.dbname").(string),
		ConnectTimeout: config.Get("db.options.connectTimeout").(string),
	})
	if err != nil {
		panic(err)
	}

	ur, err := userRepo.NewUserRepo(gormDB, 2*time.Second)
	if err != nil {
		panic(err)
	}
	us, err := userService.NewUserService(&ur, 2*time.Second)

	grpcServer := grpc.NewServer()
	userGen.RegisterUserServiceServer(grpcServer, &user.UserServiceServer{
		GormDB:      gormDB,
		UserRepo:    ur,
		UserService: us,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
