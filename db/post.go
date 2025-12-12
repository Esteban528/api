package db

import (
	"estebandev_api/types"
	"log"
	"time"
)

type Post types.Post

func createPost(p *Post) error {
	res, err := db.Exec(`INSERT INTO posts 
						(author, title, description, content, created_at) VALUES (?,?,?,?,?)`,
		p.Author, p.Title, p.Description, p.Content, time.Now())

	if err == nil {
		log.Printf("Post created: %+v\n", p)
	}

	id, _ := res.LastInsertId()
	p.ID = int(id)

	events.NotifyAll[Post]("create_post", *p)

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

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return posts
	}
	defer rows.Close()

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

	_, err := db.Exec(`UPDATE posts SET author=?, title=?, description=?, content=? 
	WHERE id=?`, p.Author, p.Title, p.Description, p.Content, p.ID)

	if err == nil {
		log.Printf("Post updated: %+v\n", p)
	}

	events.NotifyAll[Post]("update_post", *p)

	return err
}

func (p *Post) Delete() error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", p.ID)
	if err != nil {
		log.Println("DB error deleting", p, err)
	}
	return err
}
