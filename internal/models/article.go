package models

import "github.com/gin-gonic/gin"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

func NewArticle() Article {
	return Article{}
}

func (t *Article) Get(c *gin.Context) {

}

func (t *Article) List(c *gin.Context) {

}

func (t *Article) Create(c *gin.Context) {

}

func (t *Article) Update(c *gin.Context) {

}

func (t *Article) Delete(c *gin.Context) {

}
