package cv

import "time"

type GetResumeDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type ResumeInput struct {
	Name                  string    `json:"name" binding:"required"`
	Email                 string    `json:"email" binding:"required,email"`
	PhoneNumber           string    `json:"phone_number" binding:"required"`
	LinkedinURL           string    `json:"linkedin_url" binding:""`
	PortofolioURL         string    `json:"portofolio_url" binding:""`
	CompanyName           string    `json:"company_name"`
	OccupationPosition    string    `json:"occupation_position"`
	OccupationStart       time.Time `json:"occupation_start"`
	OccupationEnd         time.Time `json:"occupation_end"`
	OccupationStatus      string    `json:"occupation_status"`
	OccupationAchievement string    `json:"occupation_achievement"`
	EducationName         string    `json:"education_name"`
	EducationDegree       string    `json:"education_degree"`
	EducationFaculty      string    `json:"education_faculty"`
	EducationCity         string    `json:"education_city"`
	EducationStart        time.Time `json:"education_start"`
	EducationEnd          time.Time `json:"education_end"`
	EducationScore        float32   `json:"education_score"`
	Achievements          string    `json:"achievements" binding:"required"`
}
