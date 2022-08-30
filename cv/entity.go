package cv

import "time"

type Resume struct {
	ID                    int
	Name                  string
	Email                 string
	PhoneNumber           string
	LinkedinURL           string
	PortofolioURL         string
	CompanyName           string
	OccupationPosition    string
	OccupationStart       time.Time
	OccupationEnd         time.Time
	OccupationStatus      string
	OccupationAchievement string
	EducationName         string
	EducationDegree       string
	EducationFaculty      string
	EducationCity         string
	EducationStart        time.Time
	EducationEnd          time.Time
	EducationScore        float32
	Achievements          string
}
