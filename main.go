package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/streadway/amqp"
	"goSendEmail/routers"
	"goSendEmail/utils/pools"
	"log"
	"net"
	"time"
)

var (
	mqUrl  = "amqp://guest:guest@110.40.141.168:5672/" //根据实际情况填写mq配置连接
	mqPool pools.Pool
)

func main() {
	routers.Init()
	web.Run()

	//rabbitmq()
	////拿到一个连接
	//mq, _ := mqPool.Get()
	////实例化对象
	//mqConn := mq.(*amqp.Connection)
	////将连接放回连接池中
	//defer mqPool.Put(mq)
	////开始操作rabbitmq...
	//channel, err := mqConn.Channel()
	//if err != nil {
	//	return
	//}
	//send(channel)
	//do something....
}

func send(channel *amqp.Channel) {
	q, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	body := "Hello World!"
	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//rabbitmq rabbitmq连接池
func rabbitmq() {
	//factory 创建连接的方法
	factory := func() (interface{}, error) { return amqp.Dial(mqUrl) }
	//close 关闭连接的方法
	close := func(v interface{}) error { return v.(net.Conn).Close() }
	//创建一个连接池： 初始化2，最大连接5，空闲连接数是4
	poolConfig := &pools.Config{
		InitialCap: 2,
		MaxIdle:    5,
		MaxCap:     5,
		Factory:    factory,
		Close:      close,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	mqPool, _ = pools.NewChannelPool(poolConfig)
	//从连接池中取得一个连接
	//v, err := p.Get()
	//do something
	//conn :=v.(*amqp.Connection)
	//将连接放回连接池中
	//p.Put(v)
	//释放连接池中的所有连接
	//p.Release()
	//查看当前连接中的数量
	current := mqPool.Len()
	fmt.Println("len=", current)
	return
}
