package cv

type Service interface {
	GetResume() ([]Resume, error)
	GetResumeByID(input GetResumeDetailInput) (Resume, error)
	CreateResume(input ResumeInput) (Resume, error)
	UpdateResume(inputID GetResumeDetailInput, inputData ResumeInput) (Resume, error)
	DeleteResume(inputID GetResumeDetailInput) (Resume, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetResume() ([]Resume, error) {
	resume, err := s.repository.GetAll()
	if err != nil {
		return resume, err
	}

	return resume, nil
}

func (s *service) GetResumeByID(input GetResumeDetailInput) (Resume, error) {
	resume, err := s.repository.GetByID(input.ID)
	if err != nil {
		return resume, err
	}

	return resume, nil
}

func (s *service) CreateResume(input ResumeInput) (Resume, error) {
	resume := Resume{}
	resume.Name = input.Name
	resume.Email = input.Email
	resume.PhoneNumber = input.PhoneNumber
	resume.LinkedinURL = input.LinkedinURL
	resume.PortofolioURL = input.PortofolioURL
	resume.CompanyName = input.CompanyName
	resume.OccupationPosition = input.OccupationPosition
	resume.OccupationStart = input.OccupationStart
	resume.OccupationEnd = input.OccupationEnd
	resume.OccupationStatus = input.OccupationStatus
	resume.OccupationAchievement = input.OccupationAchievement
	resume.EducationName = input.EducationName
	resume.EducationDegree = input.EducationDegree
	resume.EducationFaculty = input.EducationFaculty
	resume.EducationCity = input.EducationCity
	resume.EducationStart = input.EducationStart
	resume.EducationEnd = input.EducationEnd
	resume.EducationScore = input.EducationScore
	resume.Achievements = input.Achievements

	newResume, err := s.repository.Save(resume)
	if err != nil {
		return newResume, err
	}

	return newResume, nil
}

func (s *service) UpdateResume(inputID GetResumeDetailInput, inputData ResumeInput) (Resume, error) {
	resume, err := s.repository.GetByID(inputID.ID)
	if err != nil {
		return resume, err
	}

	resume.Name = inputData.Name
	resume.Email = inputData.Email
	resume.PhoneNumber = inputData.PhoneNumber
	resume.LinkedinURL = inputData.LinkedinURL
	resume.PortofolioURL = inputData.PortofolioURL
	resume.CompanyName = inputData.CompanyName
	resume.OccupationPosition = inputData.OccupationPosition
	resume.OccupationStart = inputData.OccupationStart
	resume.OccupationEnd = inputData.OccupationEnd
	resume.OccupationStatus = inputData.OccupationStatus
	resume.OccupationAchievement = inputData.OccupationAchievement
	resume.EducationName = inputData.EducationName
	resume.EducationDegree = inputData.EducationDegree
	resume.EducationFaculty = inputData.EducationFaculty
	resume.EducationCity = inputData.EducationCity
	resume.EducationStart = inputData.EducationStart
	resume.EducationEnd = inputData.EducationEnd
	resume.EducationScore = inputData.EducationScore
	resume.Achievements = inputData.Achievements

	updatedResume, err := s.repository.Update(resume)
	if err != nil {
		return updatedResume, err
	}

	return updatedResume, nil
}

func (s *service) DeleteResume(inputID GetResumeDetailInput) (Resume, error) {
	resume, err := s.repository.GetByID(inputID.ID)
	if err != nil {
		return resume, err
	}

	err = s.repository.Delete(resume)
	if err != nil {
		return resume, err
	}

	return resume, nil
}
