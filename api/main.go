package main

import (
	"yuelng.com/explorer/api/handlers"
	"yuelng.com/explorer/api/migration"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	//"base/rpc"
	//pb "base/protos/helloworld"
	"yuelng.com/explorer/api/services"
	"yuelng.com/explorer/api/models"
)

const (
	address = "localhost:50051"
)

func init() {
	migration.CreateDatabase()

	services.InitDB()
	models.InitSchema()
	models.Seed()

	services.InitGrpc()
	services.InitMem()
}

func main() {
	p := fmt.Println

	// flag 解析命令行参数
	// flag.String("e", "default", "help message")
	// 使用 flag.String()  Bool(), Int()
	// 最后调用flag.Parse() 对命令行参数进行解析
	enviroment := flag.String("e", "development", "")
	flag.Parse()
	p(*enviroment)

	r := gin.Default()
	handlers.Register(r)
	r.Run(":8080")

	//rpc.StartServiceConns("","")
	//defer conn.Close()
	//c := pb.NewGreeterClient(conn)
}
