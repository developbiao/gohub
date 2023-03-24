package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/app/provider/hello"
	"gohub/framework"
	"gohub/pkg/console"
	"gohub/pkg/redis"
	"time"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

func runPlay(cmd *cobra.Command, args []string) {
	// Set value to redis
	redis.Redis.Set("hello", "Hi from redis", 10*time.Second)
	// Get value from redis
	console.Success(redis.Redis.Get("hello"))

	// Create container
	container := framework.NewGoHubContainer()
	fmt.Printf("before is bind key: %v\n", container.IsBind(hello.KEY))
	// Bind provider
	container.Bind(&hello.HelloServiceProvider{})

	// Get service instance
	helloService := container.MustMake(hello.KEY).(hello.Service)

	// Call method on service instance
	foo := helloService.SayHello()
	fmt.Println("service:", foo)

	fmt.Printf("after is bind key: %v\n", container.IsBind(hello.KEY))
}
