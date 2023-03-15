package config

import "gohub/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{

			// Default Number of Per page
			"perpage": 10,

			// URL page parameter field
			// If this value is modified, the request verification rules need to be modified at the same time
			"url_query_page": "page",

			// The parameter used to distinguish the order in the URL (use id or other)
			// If this value is modified, the request verification rules need to be modified at the same time
			"url_query_sort": "sort",

			// The parameter used to distinguish the sort in the URL (use ASC or DESC)
			// If this value is modified, the request verification rules need to be modified at the same time
			"url_query_order": "order",

			// URL page parameter per_page field
			// If this value is modified, the request verification rules need to be modified at the same time
			"url_query_per_page": "per_page",
		}
	})
}
