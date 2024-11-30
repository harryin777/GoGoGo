package main

import (
	"fmt"
)

/* 对于一个原始的结构体，我们想在不更改原始结构体或者原始功能的基础上对这个结构体的行为进行拓展，这种设计模式就是装饰器模式。
核心在于，新的行为所抽象成的结构体，必须要包含对原结构体的引用，才可以达到效果
*/

// 定义一个简单的 Client 接口
type Client interface {
	DoSomething() string
}

// 默认客户端实现
type DefaultClient struct{}

func (d *DefaultClient) DoSomething() string {
	return "Doing something in DefaultClient"
}

// Wrapper 用于包装 Client
type Wrapper func(Client) Client

// 分片结构
type shard struct {
	key    string
	Client Client
}

func (s *shard) DoSomething() string {
	return fmt.Sprintf("Sharded request based on key %s: %s", s.key, s.Client.DoSomething())
}

// NewClientWrapper 创建一个包装器，它根据一个分片键将请求路由到不同的客户端
func NewClientWrapper(key string) Wrapper {
	return func(c Client) Client {
		return &shard{
			key:    key,
			Client: c,
		}
	}
}

//func main() {
//	// 创建默认客户端
//	var cl Client = &DefaultClient{}
//
//	// 包装客户端，添加分片逻辑
//	cl = NewClientWrapper("ShardKey")(cl)
//
//	// 调用包装后的客户端方法
//	fmt.Println(cl.DoSomething())
//}
