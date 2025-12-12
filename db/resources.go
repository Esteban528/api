package db

import (
	"estebandev_api/types"
	"log"
)

type Resource types.Resource

func createResource(r *Resource) error {
	_, err := db.Exec(`INSERT INTO resources 
		(title, description, link, image_url) 
		VALUES (?,?,?,?)`,
		r.Title, r.Description, r.Link, r.Image_URL)

	return err
}

func FindAllResources() []Resource {
	resources := []Resource{}

	rows, err := db.Query("SELECT * FROM resources")

	if err != nil {
		return resources
	}

	defer rows.Close()

	for rows.Next() {
		resource := Resource{}
		err := rows.Scan(
			&resource.ID,
			&resource.Title,
			&resource.Description,
			&resource.Link,
			&resource.Image_URL,
		)

		if err != nil {
			continue
		}
		resources = append(resources, resource)
	}
	return resources
}

func FindResource(id int) (Resource, error) {
	resource := Resource{}

	row := db.QueryRow("SELECT * FROM resources WHERE id = ?", id)
	err := row.Scan(
		&resource.ID,
		&resource.Title,
		&resource.Description,
		&resource.Link,
		&resource.Image_URL,
	)

	return resource, err
}

func (r *Resource) Save() error {
	if r.ID == 0 {
		return createResource(r)
	}

	_, err := db.Exec(`UPDATE resources 
		SET title=?, description=?, link=?, image_url=? 
		WHERE id=?`,
		r.Title, r.Description, r.Link, r.Image_URL, r.ID)

	if err != nil {
		log.Println("DB error updating", r, err)
	}
	return err
}

func (r *Resource) Delete() error {
	_, err := db.Exec("DELETE FROM resources WHERE id = ?", r.ID)

	if err != nil {
		log.Println("DB error deleting", r, err)
	}
	return err
}
