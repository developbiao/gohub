package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/cache"
	"gohub/pkg/console"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cahe-key",
	Run:   runCacheForget,
}

// forget command option
var cacheKey string

func init() {
	// Registration Sub clear command
	CmdCache.AddCommand(CmdCacheClear)

	// Set cache forget command option
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	CmdCache.AddCommand(CmdCacheForget)
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared!")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] has been deleted!", cacheKey))
}
