package main

import (
	"awesomeProject/web_admin/model"
	_ "awesomeProject/web_admin/router" //只执行包中的init方法
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
	"go.etcd.io/etcd/clientv3"
	"time"

	"github.com/astaxie/beego"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:Admin@123@tcp(192.168.2.4:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}
	model.InitDb(database)
	return
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.2.4:2379",},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	model.InitEtcd(cli)
	return
}

func main() {
	err := initDb()
	if err != nil {
		logs.Warn("initDb failed,err:%v", err)
		return
	}
	err = initEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return
	}
	beego.Run()
}
