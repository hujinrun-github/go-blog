package model

import (
	"fmt"
)

// Category 记录不同类别的文章信息
type Category struct {
	CategoryID  int    // 类别id
	Description string // 类别描述
	Articles    []*Content
}

var categories map[string]*Category

// SetArticleIntoCategory 用于将文章设置进类别中
func SetArticleIntoCategory(key string, content *Content) error {
	if content == nil {
		return fmt.Errorf("content should not be error")
	}

	if _, ok := categories[key]; !ok {
		categories[key] = &Category{}
		categories[key].Articles = []*Content{content}
	}

	return nil
}
