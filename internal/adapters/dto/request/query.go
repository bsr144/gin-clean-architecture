package request

type Query struct {
	Search string `form:"search"`
	Page   int    `form:"page"`
	Size   int    `form:"size"`
}
