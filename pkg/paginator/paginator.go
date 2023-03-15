package paginator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"strings"
)

// Paging paging data
type Paging struct {
	CurrentPage int    // Current page
	PerPage     int    // Number of Per page
	TotalPage   int    // Total page
	TotalCount  int64  // Total count
	NextPageURL string // Next page url
	PrevPageURL string // Previous url
}

// Paginator Paginator operator
type Paginator struct {
	BaseURL    string // Contact URL
	PerPage    int    // Number of per page
	Page       int    // Current Page
	Offset     int    // Database read offset value
	TotalCount int64  // Total count
	TotalPage  int    // Total page = TotalCount / PerPage
	Sort       string // Sort rule
	Order      string // Order ASC or DESC

	query *gorm.DB     // db query handle
	ctx   *gin.Context // gin context
}

// Paginate
// c —— gin.context Get paging URL parameters
// db —— GORM query handle, for query data total count
// baseURL —— paging link
// data —— model array, address reference  get data
// PerPage —— Number of Per page, priority get from url parameter otherwise use perPage value
// Usage:
//         query := database.DB.Model(Topic{}).Where("category_id = ?", cid)
//      var topics []Topic
//         paging := paginator.Paginate(
//             c,
//             query,
//             &topics,
//             app.APIURL(database.TableName(&Topic{})),
//             perPage,
//
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {
	// Initialization Paginator
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseURL)

	// Query database
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).Error
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

// initProperties initialization paginator properties
func (p *Paginator) initProperties(perPage int, baseURL string) {
	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)

	// Order parameters, have been verified in the controller and can be used with confidence
	p.Order = p.ctx.DefaultQuery(config.Get("paging_url_query_order"), "ASC")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging_url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

// formatBaseURL compatible with URLs with and without `?`
func (p *Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

// getTotalCount get number of total
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// getTotalPage get number of total page
func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return nums
}

// Get current page
func (p *Paginator) getCurrentPage() int {
	// Priority get user request page
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page <= 0 {
		// Default is 1
		page = 1
	}

	// TotalPage equal 0, means that the data is not paginated enough
	if p.TotalPage == 0 {
		return 0
	}

	// Request page greater than total page
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

// GetPerPage get number of per page
func (p *Paginator) getPerPage(perPage int) int {
	// First use the request `per_page` parameter
	queryPerPage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	// Not pass parameter using default perpage
	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}
	return perPage
}

// getNextPageURL get next page url address
func (p *Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

// getPrevPageUrl get previous  page url
func (p *Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}

func (p *Paginator) getPageLink(page int) string {

	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)

}
