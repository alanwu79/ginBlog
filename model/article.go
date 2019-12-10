package model

import "time"

type ArticleInfo struct {
	Id           int64
	CategoryId   int64
	Summary      string
	Title        string
	ViewCount    uint32
	CreatTime    time.Time
	CommentCount uint32
	Username     string
}

type ArticleDetail struct {
	ArticleInfo
	Content string
	Category
}
