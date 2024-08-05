package models

import "time"

type Contact struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Location  Location `json:"location"`
}

type Location struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}

type SocialMedia struct {
	Github   string `json:"github"`
	LinkedIn string `json:"linkedin"`
}

type Summary struct {
	Description string `json:"description"`
}

type Language struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type Skill struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       uint8  `json:"level"`
}

type Experience struct {
	Employer    string     `json:"employer"`
	Role        string     `json:"role"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Present     bool       `json:"is_present"`
	Remote      bool       `json:"is_remote"`
	Location    *Location  `json:"location"`
	Description string     `json:"description"`
}

type Education struct {
	School         string    `json:"school"`
	Field          string    `json:"field"`
	GraduationDate time.Time `json:"graduation_date"`
	Location       Location  `json:"location"`
	Description    string    `json:"description"`
}

type Resume struct {
	Contact     Contact       `json:"contact"`
	Summary     Summary       `json:"summary"`
	SocialMedia *SocialMedia  `json:"social_media"`
	Languages   *[]Language   `json:"languages"`
	Skills      *[]Skill      `json:"skills"`
	Experiences *[]Experience `json:"experiences"`
	Educations  *[]Education  `json:"educations"`
}
