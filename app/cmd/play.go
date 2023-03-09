package cmd

import (
	"github.com/spf13/cobra"
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
}
