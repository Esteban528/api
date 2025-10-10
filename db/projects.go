package db

import "log"

type Project struct {
	ID          int
	Title       string
	Description string
	Visit_URL   string
	Source_URL  string
}

func createProject(p *Project) error {
	_, err := db.Exec(`INSERT INTO projects 
						(title, description, visit_url, source_url) VALUES (?,?,?,?)`,
		p.Title, p.Description, p.Visit_URL, p.Source_URL)

	return err
}

func FindAllProject() ([]Project) {
	projects := []Project{}

	rows,_ := db.Query("SELECT * FROM projects")

	for rows.Next() {
		project := Project{}
		err := rows.Scan(&project.ID, &project.Title, 
			&project.Description, &project.Visit_URL, &project.Source_URL)

		if err != nil {
			continue
		}
		projects = append(projects, project)
	}
	return projects
}

func FindProject(id int) (Project, error) {
	project := Project{}

	row := db.QueryRow("SELECT * FROM projects WHERE id = ?", id)
	err := row.Scan(&project.ID, &project.Title, 
		&project.Description, &project.Visit_URL, &project.Source_URL)

	return project, err
}

func (p *Project) Save() error {
	if p.ID == 0 {
		return createProject(p)
	}

	_, err := db.Exec(`UPDATE projects 
		SET title=?, description=?, visit_url=?, source_url=? 
		WHERE id=?`,
		p.Title, p.Description, p.Visit_URL, p.Source_URL, p.ID)

	if err != nil {
		log.Println("DB error deleting", p, err)
	}
	return err
}

func (p *Project) Delete() error {
	_, err := db.Exec("DELETE FROM projects WHERE id = ?", p.ID)

	if err != nil {
		log.Println("DB error deleting", p, err)
	}
	return err
}
