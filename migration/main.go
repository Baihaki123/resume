package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Resume struct {
	ID                    int       `db:"id"`
	Name                  string    `db:"name"`
	Email                 string    `db:"email"`
	PhoneNumber           string    `db:"phone_number"`
	LinkedInURL           string    `db:"linkedin_url"`
	PortofolioURL         string    `db:"portofolio_url"`
	CompanyName           string    `db:"company_name"`
	OccupationPosition    string    `db:"occupation_position"`
	OccupationStart       time.Time `db:"occupation_start"`
	OccupationEnd         time.Time `db:"occupation_end"`
	OccupationStatus      string    `db:"occupation_status"`
	OccupationAchievement string    `db:"occupation_achievement"`
	EducationName         string    `db:"education_name"`
	EducationDegree       string    `db:"education_degree"`
	EducationFaculty      string    `db:"education_faculty"`
	EducationCity         string    `db:"education_city"`
	EducationStart        time.Time `db:"education_start"`
	EducationEnd          time.Time `db:"education_end"`
	EducationScore        float32   `db:"education_score"`
	Achievement           string    `db:"achievements"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../resume.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS resumes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(50),
		email VARCHAR(50),
		phone_number VARCHAR(128),
		linkedin_url VARCHAR(128),
		portofolio_url VARCHAR(128),
		company_name VARCHAR(50),
		occupation_position VARCHAR(50),
		occupation_start DATETIME,
		occupation_end DATETIME,
		occupation_status VARCHAR(20),
		occupation_achievement VARCHAR(128),
		education_name VARCHAR(50),
		education_degree VARCHAR(50),
		education_faculty VARCHAR(50),
		education_city VARCHAR(128),
		education_start DATETIME,
		education_end DATETIME,
		education_score FLOAT(3,2),
		achievements VARCHAR(128)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Rollback(db *sql.DB) {
	sqlStmt := `DROP TABLE resumes;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

// Jalankan main untuk melakukan migrasi database
func main() {
	_, err := Migrate()
	if err != nil {
		panic(err)
	}

	// Use Rollback() // to rollback the changes
	// Rollback(db)
}
