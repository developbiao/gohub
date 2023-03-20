package link

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/cache"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"
	"time"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {
	//Set cache key
	cacheKey := "links:all"
	// Set cache time
	expireTime := 120 * time.Minute
	// Get data
	cache.GetObject(cacheKey, &links)

	// If not hit cache query from database
	if helpers.Empty(links) {
		// Query database
		links = All()
		if helpers.Empty(links) {
			return
		}
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
