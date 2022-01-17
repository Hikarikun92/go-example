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
	rows, err := r.db.Query("select p.id, p.title, p.body, p.published_date, author.id, author.username, c.id, "+
		"c.title, c.body, c.published_date, commenter.id, commenter.username from post p join user author on "+
		"author.id = p.user_id left join comment c on c.post_id = p.id left join user commenter on c.user_id = commenter.id "+
		"where p.id = ? order by c.published_date", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Read all the rows first so that we can perform checks on them later
	type record struct {
		id                   int
		title                string
		body                 string
		publishedDate        time.Time
		authorId             int
		authorUsername       string
		commentId            sql.NullInt32
		commentTitle         sql.NullString
		commentBody          sql.NullString
		commentPublishedDate sql.NullTime
		commenterId          sql.NullInt32
		commenterUsername    sql.NullString
	}
	var records []*record

	for rows.Next() {
		r := record{}
		err := rows.Scan(&r.id, &r.title, &r.body, &r.publishedDate, &r.authorId, &r.authorUsername, &r.commentId,
			&r.commentTitle, &r.commentBody, &r.commentPublishedDate, &r.commenterId, &r.commenterUsername)

		if err != nil {
			return nil, err
		}
		records = append(records, &r)
	}

	//If there were no rows, this post doesn't exist
	if len(records) == 0 {
		return nil, nil
	}

	//Create the post and its author from the first row
	first := records[0]

	//Cache to avoid creating user entities of already converted user rows
	userCache := make(map[int]*user.User)

	author := &user.User{Id: first.authorId, Username: first.authorUsername}
	userCache[author.Id] = author

	//Go through the other rows for the post's comments
	var comments []*comment.Comment
	for _, rec := range records {
		if rec.commentId.Valid { //If commentId is not valid, it means there were no comments
			commenter, ok := userCache[int(rec.commenterId.Int32)]
			if !ok {
				commenter = &user.User{Id: int(rec.commenterId.Int32), Username: rec.commenterUsername.String}
				userCache[commenter.Id] = commenter
			}

			c := &comment.Comment{
				Id:            int(rec.commentId.Int32),
				Title:         rec.commentTitle.String,
				Body:          rec.commentBody.String,
				PublishedDate: rec.commentPublishedDate.Time,
				User:          commenter,
			}
			comments = append(comments, c)
		}
	}

	return &Post{
		Id:            first.id,
		Title:         first.title,
		Body:          first.body,
		PublishedDate: first.publishedDate,
		User:          author,
		Comments:      comments,
	}, nil
}
