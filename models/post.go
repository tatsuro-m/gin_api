package models

import (
	"post_crud_golang/database"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"size:50;not null"`
	Content string `gorm:"size:2000;not null"`
}

func dbInit() (d *gorm.DB) {
	d = database.Init()
	defer d.Close()
	d.AutoMigrate(&Post{})
	return d
}

func insert(title string, content string) {
	d := dbInit()
	d.Create(&Post{
		Title:   title,
		Content: content,
	})
	defer d.Close()
}

func update(id int, title string, content string) {
	d := dbInit()
	var post Post
	d.First(&post, id)
	post.Title = title
	post.Content = content
	d.Save(&post)
	d.Close()
}

func delete(id int) {
	d := dbInit()
	var post Post
	d.First(&post, id)
	d.Delete(&post)
	d.Close()
}

func getAll() []Post {
	d := dbInit()
	var posts []Post
	d.Order("created_at desc").Find(&posts)
	d.Close()
	return posts
}

func getFirst(id int) Post {
	d := dbInit()
	var post Post
	d.First(&post, id)
	d.Close()
	return post
}
