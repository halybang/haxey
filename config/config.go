package config

// Import here anonymously all hexya addons you need in your application
import (
	_ "github.com/hexya-erp/hexya-base/base"
	_ "github.com/hexya-erp/hexya-base/web"
	_ "github.com/hexya-erp/hexya-base/webKanban"

	_ "github.com/hexya-erp/hexya-addons/account"
	_ "github.com/hexya-erp/hexya-addons/analytic"
	_ "github.com/hexya-erp/hexya-addons/procurement"
	_ "github.com/hexya-erp/hexya-addons/product"

	_ "github.com/hexya-erp/hexya-addons/sale"
	_ "github.com/hexya-erp/hexya-addons/saleTeams"
)
