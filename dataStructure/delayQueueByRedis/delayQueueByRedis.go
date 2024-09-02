package delayQueueByRedis

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
		DB:   0, // 使用默认数据库
	})

	// 添加延时任务到队列中
	go func() {
		for i := 0; i < 10; i++ {
			message := fmt.Sprintf("task-%d", i)
			delay := time.Duration(5+i*2) * time.Second // 延时 5s, 7s, 9s...
			err := addDelayedTask(rdb, "delayedQueue", message, delay)
			if err != nil {
				log.Fatalf("Failed to add delayed task: %v", err)
			}
			fmt.Printf("Scheduled: %s for execution in %v\n", message, delay)
		}
	}()

	// 消费延时任务
	go func() {
		for {
			task, err := fetchReadyTask(rdb, "delayedQueue")
			if err != nil {
				log.Printf("Failed to fetch task: %v", err)
				continue
			}
			if task != "" {
				fmt.Printf("Executed: %s\n", task)
			}
			time.Sleep(1 * time.Second) // 控制轮询频率
		}
	}()

	// 让主程序运行一段时间，以观察调度和执行
	time.Sleep(60 * time.Second)
}

// 添加延时任务
func addDelayedTask(rdb *redis.Client, queue, task string, delay time.Duration) error {
	executionTime := time.Now().Add(delay).Unix()
	_, err := rdb.ZAdd(ctx, queue, &redis.Z{
		Score:  float64(executionTime),
		Member: task,
	}).Result()
	return err
}

// 获取准备好执行的任务
func fetchReadyTask(rdb *redis.Client, queue string) (string, error) {
	now := time.Now().Unix()
	res, err := rdb.ZRangeByScoreWithScores(ctx, queue, &redis.ZRangeBy{
		Min:   "0",
		Max:   fmt.Sprintf("%d", now),
		Count: 1,
	}).Result()
	if err != nil {
		return "", err
	}

	if len(res) > 0 {
		task := res[0].Member.(string)
		_, err = rdb.ZRem(ctx, queue, task).Result()
		if err != nil {
			return "", err
		}
		return task, nil
	}

	return "", nil
}
