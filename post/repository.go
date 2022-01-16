package post

import (
	"database/sql"
	"go-example/comment"
	"go-example/user"
	"time"
)

type Repository interface {
	FindByUserId(userId int) ([]*Post, error)
	FindById(id int) (*Post, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindByUserId(userId int) ([]*Post, error) {
	rows, err := r.db.Query("select p.id, p.title, p.body, p.published_date from post p where p.user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post

	for rows.Next() {
		p := Post{}
		if err := rows.Scan(&p.Id, &p.Title, &p.Body, &p.PublishedDate); err != nil {
			return nil, err
		}
		posts = append(posts, &p)
	}

	return posts, nil
}

func (r *repositoryImpl) FindById(id int) (*Post, error) {

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
		}, nil
	case 2:
		return &Post{
			Id:            2,
			Title:         "Another example post",
			Body:          "Integer malesuada lorem non nunc.",
			PublishedDate: time.Date(2021, time.March, 15, 17, 53, 7, 0, time.UTC),
			User:          author,
			Comments:      []*comment.Comment{},
		}, nil
	default:
		return nil, nil
	}
}
