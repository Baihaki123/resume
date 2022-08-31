package cv

import (
	"strings"
	"time"
)

//==========================================================//
//================START Resume Formatter====================//
//==========================================================//
type ResumeFormatter struct {
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	PhoneNumber   string        `json:"phone_number"`
	LinkedinURL   string        `json:"linkedin_url"`
	PortofolioURL string        `json:"portofolio_url"`
	Occupations   []Occupations `json:"occupations"`
	Educations    []Educations  `json:"educations"`
	Achievements  []string      `json:"achievements"`
}

type Occupations struct {
	CompanyName           string    `json:"company_name"`
	OccupationPosition    string    `json:"occupation_position"`
	OccupationStart       time.Time `json:"occupation_start"`
	OccupationEnd         time.Time `json:"occupation_end"`
	OccupationStatus      string    `json:"occupation_status"`
	OccupationAchievement []string  `json:"occupation_achievement"`
}

type Educations struct {
	EducationName    string    `json:"education_name"`
	EducationDegree  string    `json:"education_degree"`
	EducationFaculty string    `json:"education_faculty"`
	EducationCity    string    `json:"education_city"`
	EducationStart   time.Time `json:"education_start"`
	EducationEnd     time.Time `json:"education_end"`
	EducationScore   float32   `json:"education_score"`
}

func FormatResume(resume Resume) ResumeFormatter {
	resumeFormatter := ResumeFormatter{}
	resumeFormatter.Name = resume.Name
	resumeFormatter.Email = resume.Email
	resumeFormatter.PhoneNumber = resume.PhoneNumber
	resumeFormatter.LinkedinURL = resume.LinkedinURL
	resumeFormatter.PortofolioURL = resume.PortofolioURL

	var occupationAchievements []string
	for _, occupationAchievement := range strings.Split(resume.OccupationAchievement, ",") {
		occupationAchievements = append(occupationAchievements, strings.TrimSpace(occupationAchievement))
	}

	occupations := []Occupations{}

	resumeFormatter.Occupations = append(occupations, Occupations{
		CompanyName:           resume.CompanyName,
		OccupationPosition:    resume.OccupationPosition,
		OccupationStart:       resume.OccupationStart,
		OccupationEnd:         resume.OccupationEnd,
		OccupationStatus:      resume.OccupationStatus,
		OccupationAchievement: occupationAchievements,
	})

	educations := []Educations{}

	resumeFormatter.Educations = append(educations, Educations{
		EducationName:    resume.EducationName,
		EducationDegree:  resume.EducationDegree,
		EducationFaculty: resume.EducationFaculty,
		EducationCity:    resume.EducationCity,
		EducationStart:   resume.EducationStart,
		EducationEnd:     resume.OccupationEnd,
		EducationScore:   resume.EducationScore,
	})

	var achievements []string
	for _, achievement := range strings.Split(resume.Achievements, ",") {
		achievements = append(achievements, strings.TrimSpace(achievement))
	}

	resumeFormatter.Achievements = achievements

	return resumeFormatter
}

func FormatResumes(resumes []Resume) []ResumeFormatter {
	resumesFormatter := []ResumeFormatter{}

	for _, resume := range resumes {
		resumeFormatter := FormatResume(resume)
		resumesFormatter = append(resumesFormatter, resumeFormatter)
	}

	return resumesFormatter
}

//==========================================================//
//================END Resume Formatter======================//
//==========================================================//
