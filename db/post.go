package db

import (
	"log"
	"time"
)

type Post struct {
	ID          int
	Author      string
	Title       string
	Description string
	Content     string
	Date        string
}

func createPost(p *Post) error {
	_, err := db.Exec(`INSERT INTO posts 
						(author, title, description, content, created_at) VALUES (?,?,?,?,?)`,
		p.Author, p.Title, p.Description, p.Content, time.Now())

	if err == nil {
		log.Printf("Post created: %+v\n", p)
	}

	return err
}

func FindPost(id int) (Post, error) {
	post := Post{}

	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)
	err := row.Scan(&post.ID, &post.Author, &post.Title, &post.Description, &post.Content, &post.Date)

	if err != nil {
		log.Println("Database error: ", err)
	}

	return post, err
}

func FindAllPost() []Post {
	posts := []Post{}

	rows, _ := db.Query("SELECT * FROM posts")

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Author, &post.Title,
			&post.Description, &post.Content, &post.Date)

		if err != nil {
			continue
		}
		posts = append(posts, post)
	}

	log.Println(posts)
	return posts
}

func (p *Post) Save() error {
	if p.ID == 0 {
		return createPost(p)
	}

	_, err := db.Exec(`UPDATE posts SET author=?, title=?, description=?, content=?, created_at=? 
	WHERE id=?`, p.Author, p.Title, p.Description, p.Content, p.Date, p.ID)

	if err == nil {
		log.Printf("Post updated: %+v\n", p)
	}

	return err
}

func (p *Post) Delete() error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", p.ID)
	if err != nil {
		log.Println("DB error deleting", p, err)
	}
	return err
}
