package repository

import (
	"fmt"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/post-service/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	Database *gorm.DB
}

func (repo *PostRepository) CreatePost(post *model.Post) error {
	result := repo.Database.Create(post)
	fmt.Print(result)
	return nil
}

func (repo *PostRepository) FindAllPosts() []model.Post {
	var posts []model.Post
	repo.Database.Select("*").Find(&posts)
	return posts
}

func (repo *PostRepository) FindByID(ID uuid.UUID) *model.Post {
	post := &model.Post{}
	if repo.Database.First(&post, "id = ?", ID).RowsAffected == 0 {
		return nil
	}
	return post
}

func (repo *PostRepository) FindAllPostsForUser(userId uuid.UUID) []model.Post {
	var posts []model.Post
	repo.Database.Select("*").Where("user_id = ? ", userId).Find(&posts)
	return posts
}
