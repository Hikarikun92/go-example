package post

import (
	"go-example/comment"
	"go-example/user"
	"time"
)

type Repository interface {
	FindByUserId(userId int) []*Post
	FindById(id int) *Post
}

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) FindByUserId(userId int) []*Post {
	author := &user.User{Id: userId, Username: "Username"}

	return []*Post{
		{
			Id:            1,
			Title:         "Example post no. 1",
			Body:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse placerat.",
			PublishedDate: time.Date(2021, time.January, 1, 12, 3, 18, 0, time.UTC),
			User:          author,
			Comments: []*comment.Comment{
				{
					Id:            1,
					Title:         "Example comment 1",
					Body:          "Praesent sapien leo, viverra sed.",
					PublishedDate: time.Date(2021, time.January, 1, 18, 42, 32, 0, time.UTC),
					User:          &user.User{Id: 2, Username: "John Doe"},
				},
				{
					Id:            2,
					Title:         "Great article",
					Body:          "Nice example!",
					PublishedDate: time.Date(2021, time.February, 28, 7, 38, 12, 0, time.UTC),
					User:          &user.User{Id: 3, Username: "Mary Doe"},
				},
			},
		},
		{
			Id:            2,
			Title:         "Another example post",
			Body:          "Integer malesuada lorem non nunc.",
			PublishedDate: time.Date(2021, time.March, 15, 17, 53, 7, 0, time.UTC),
			User:          author,
			Comments:      []*comment.Comment{},
		},
	}
}

func (r *repositoryImpl) FindById(id int) *Post {
	author := &user.User{Id: 1, Username: "Username"}

	switch id {
	case 1:
		return &Post{
			Id:            1,
			Title:         "Example post no. 1",
			Body:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse placerat.",
			PublishedDate: time.Date(2021, time.January, 1, 12, 3, 18, 0, time.UTC),
			User:          author,
			Comments: []*comment.Comment{
				{
					Id:            1,
					Title:         "Example comment 1",
					Body:          "Praesent sapien leo, viverra sed.",
					PublishedDate: time.Date(2021, time.January, 1, 18, 42, 32, 0, time.UTC),
					User:          &user.User{Id: 2, Username: "John Doe"},
				},
				{
					Id:            2,
					Title:         "Great article",
					Body:          "Nice example!",
					PublishedDate: time.Date(2021, time.February, 28, 7, 38, 12, 0, time.UTC),
					User:          &user.User{Id: 3, Username: "Mary Doe"},
				},
			},
		}
	case 2:
		return &Post{
			Id:            2,
			Title:         "Another example post",
			Body:          "Integer malesuada lorem non nunc.",
			PublishedDate: time.Date(2021, time.March, 15, 17, 53, 7, 0, time.UTC),
			User:          author,
			Comments:      []*comment.Comment{},
		}
	default:
		return nil
	}
}
