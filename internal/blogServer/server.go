package blogServer

import (
	"errors"
	"gorm.io/gorm"
)

func Get(id int) (*Blog, error) {
	var blogs Blog
	err := DB.First(&blogs, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &blogs, nil
}
func GetAll() ([]Blog, error) {
	var blogs []Blog
	err := DB.Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
func Update(id int, blog *Blog) error {
	var b Blog
	DB.First(&b, id)
	if blog.Title != "" {
		b.Title = blog.Title
	}
	if blog.Content != "" {
		b.Content = blog.Content
	}
	err := DB.Save(&b).Error
	if err != nil {
		return err
	}
	return nil
}
func Delete(id int) error {
	err := DB.Delete(&Blog{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func Insert(blog *Blog) error {
	err := DB.Create(blog).Error
	if err != nil {
		return err
	}
	return nil
}
