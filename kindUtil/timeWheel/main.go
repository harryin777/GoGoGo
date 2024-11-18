package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定时任务结构
type Task struct {
	delay int    // 延迟时间（毫秒）
	task  func() // 要执行的任务
	round int    // 剩余的轮数
}

// TimeWheel 时间轮结构
type TimeWheel struct {
	interval    time.Duration // 每个槽的时间间隔
	slots       [][]*Task     // 时间轮的槽
	currentPos  int           // 当前指针位置
	size        int           // 时间轮大小（槽数量）
	mutex       sync.Mutex    // 保护任务添加和执行的并发安全
	ticker      *time.Ticker  // 定时器
	stopChannel chan struct{} // 停止信号
}

// NewTimeWheel 初始化时间轮
func NewTimeWheel(interval time.Duration, size int) *TimeWheel {
	slots := make([][]*Task, size)
	for i := range slots {
		slots[i] = make([]*Task, 0)
	}
	return &TimeWheel{
		interval:    interval,
		slots:       slots,
		size:        size,
		currentPos:  0,
		ticker:      time.NewTicker(interval),
		stopChannel: make(chan struct{}),
	}
}

// AddTask 添加任务到时间轮
func (tw *TimeWheel) AddTask(delay time.Duration, task func()) {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	// 计算任务位置和轮数
	totalTicks := int(delay / tw.interval)
	slot := (tw.currentPos + totalTicks) % tw.size
	round := totalTicks / tw.size

	// 创建任务
	newTask := &Task{
		delay: int(delay.Milliseconds()),
		task:  task,
		round: round,
	}

	// 添加到对应槽
	tw.slots[slot] = append(tw.slots[slot], newTask)
}

// Start 启动时间轮
func (tw *TimeWheel) Start() {
	go func() {
		for {
			select {
			case <-tw.ticker.C:
				tw.tick()
			case <-tw.stopChannel:
				return
			}
		}
	}()
}

// Stop 停止时间轮
func (tw *TimeWheel) Stop() {
	close(tw.stopChannel)
	tw.ticker.Stop()
}

// tick 时间轮移动一格
func (tw *TimeWheel) tick() {
	tw.mutex.Lock()
	defer tw.mutex.Unlock()

	// 获取当前槽的任务
	slot := tw.slots[tw.currentPos]
	tw.slots[tw.currentPos] = make([]*Task, 0) // 清空当前槽

	// 执行到期任务或减少任务轮数
	for _, task := range slot {
		if task.round > 0 {
			task.round--
			// 如果还有轮数，放回当前槽
			tw.slots[tw.currentPos] = append(tw.slots[tw.currentPos], task)
		} else {
			// 执行任务
			go task.task()
		}
	}

	// 指针移动到下一个槽
	tw.currentPos = (tw.currentPos + 1) % tw.size
}

func main() {
	// 初始化时间轮：100ms 间隔，10 个槽
	tw := NewTimeWheel(100*time.Millisecond, 10)

	// 添加任务
	tw.AddTask(150*time.Millisecond, func() {
		fmt.Println("任务 1 执行于：", time.Now())
	})
	tw.AddTask(950*time.Millisecond, func() {
		fmt.Println("任务 2 执行于：", time.Now())
	})
	tw.AddTask(2050*time.Millisecond, func() {
		fmt.Println("任务 3 执行于：", time.Now())
	})

	// 启动时间轮
	tw.Start()

	// 主程序等待 3 秒
	time.Sleep(3 * time.Second)

	// 停止时间轮
	tw.Stop()
	fmt.Println("时间轮已停止")
}
