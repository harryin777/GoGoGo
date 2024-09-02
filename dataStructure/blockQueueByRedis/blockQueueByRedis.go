package blockQueueByRedis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})

	// 启动生产者
	go producer(rdb, "mainQueue")

	// 启动消费者
	go consumerWithBackup(rdb, "mainQueue", "backupQueue")

	// 让主程序运行一段时间，以观察生产和消费
	time.Sleep(20 * time.Second)
}

// producer 将消息推送到主队列中
func producer(rdb *redis.Client, queueName string) {
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("message-%d", i)
		err := rdb.RPush(ctx, queueName, message).Err()
		if err != nil {
			log.Fatalf("Failed to push message: %v", err)
		}
		fmt.Printf("Produced: %s\n", message)
		time.Sleep(1 * time.Second) // 模拟生产延迟
	}
}

// consumerWithBackup 从队列中阻塞地获取消息，并备份到另一个队列
func consumerWithBackup(rdb *redis.Client, mainQueue, backupQueue string) {
	luaScript := `
	local msg = redis.call('BLPOP', KEYS[1], ARGV[1])
	if msg then
		redis.call('RPUSH', KEYS[2], msg[2])
		return msg[2]
	end
	return nil
	`

	for {
		// 使用 Lua 脚本从主队列中获取消息，并将其备份到备份队列
		res, err := rdb.Eval(ctx, luaScript, []string{mainQueue, backupQueue}, 0).Result()
		if err != nil {
			log.Printf("Failed to execute Lua script: %v", err)
			continue
		}

		if res != nil {
			fmt.Printf("Consumed: %s\n", res)
			// 在这里处理消息（消费）
			// 如果消费成功，可以选择从备份队列中删除对应的消息, 其实每次启动可以先看这个backupqueue里有没有元素
			// 如果有可以到下游的db或者接口验证下消费成功没有，没有消费成功就消费，消费成功只是没有移除就移除
		}
	}
}
