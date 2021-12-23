package models

import "github.com/gin-gonic/gin"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

func NewTag() Tag {
	return Tag{}
}

func (t *Tag) Get(c *gin.Context) {

}

func (t *Tag) List(c *gin.Context) {

}

func (t *Tag) Create(c *gin.Context) {

}

func (t *Tag) Update(c *gin.Context) {

}

func (t *Tag) Delete(c *gin.Context) {

}
