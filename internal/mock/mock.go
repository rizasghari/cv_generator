package mock

import (
	"github.com/rizasgahri/cv_builder/internal/models"
)

func GetMockResume() models.Resume {
	return models.Resume{
		Contact: models.Contact{
			FirstName: "Reza",
			LastName:  "Asghari",
			Phone:     "+905368375804",
			Email:     "developer.rizasghari@gmail.com",
			Location: models.Location{
				Country: "Türkiye",
				State:   "Istanbul",
				City:    " Kağıthane",
			},
		},
		Skills: &models.Skill{
			Name:  "Golang",
			Level: 4,
		},
		Summary: models.Summary{
			Description: "Seasoned Software Engineering Manager with extensive experience in leading software development teams and managing complex projects from conception to deployment. Expert in mobile and web application development with a proven track record of integrating Agile frameworks to enhance product quality and operational efficiency. Skilled in fostering cross-functional collaboration and directly engaging with high-level stakeholders to ensure project success and customer satisfaction. Demonstrates a strong command of containerization technologies, software architecture, and cloud infrastructures, coupled with a dedication to mentoring teams and driving continuous innovation and improvement.",
		},
		SocialMedia: &models.SocialMedia{
			Github:   "https://github.com/rizasgahri",
			LinkedIn: "https://linkedin.com/ln/rizaasgahri",
		},
		Languages: &models.Language{
			Name:  "Turkish",
			Level: "Native",
		},
		Experiences: &models.Experience{
			Employer:  "Google",
			Role:      "Software Engineering Manager",
			StartDate: "time.Date(2019, time.August, 1, 0, 0, 0, 0, time.Local)",
			EndDate:   "time.Date(2019, time.August, 1, 0, 0, 0, 0, time.Local)",
			Present:   true,
			Location: &models.Location{
				Country: "Türkiye",
				State:   "Istanbul",
				City:    "Maslak",
			},
			Description: "This is a descripion about my job in google as a software enginner",
		},
		Educations: &models.Education{
			School:         "UUT",
			Field:          "IT Engineering",
			GraduationDate: "2024-09-03",
			Location: models.Location{
				Country: "Iran",
				State:   "West Azerbaijan",
				City:    "Urmia",
			},
			Description: "This is a descripion about my education at UUT university",
		},
	}
}
