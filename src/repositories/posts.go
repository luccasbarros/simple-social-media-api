package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(posts models.Post) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO posts (title, content, author_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(posts.Title, posts.Content, posts.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository Posts) GetById(postID uint64) (models.Post, error) {
	row, erro := repository.db.Query(`
	select p.*, u.nick from
	posts p inner join users u
	on u.id = p.author_id where p.id = ?`, postID)
	if erro != nil {
		return models.Post{}, erro
	}

	defer row.Close()

	var post models.Post

	if row.Next() {
		if erro = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); erro != nil {
			return models.Post{}, erro
		}
	}

	return post, nil
}

func (repository Posts) GetPersonalPosts(userID uint64) ([]models.Post, error) {
	rows, erro := repository.db.Query(`
	select distinct p.*, u.nick from posts p
	inner join users u on u.id = p.author_id
	inner join followers f on p.author_id = f.user_id
	where u.id = ? or f.follower_id = ?
	order by 1 desc`, userID, userID)
	if erro != nil {
		return nil, erro
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if erro = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil

}