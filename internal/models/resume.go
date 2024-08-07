package models

type Contact struct {
	FirstName string   `json:"first_name" form:"first_name"`
	LastName  string   `json:"last_name" form:"last_name"`
	Email     string   `json:"email" form:"email"`
	Phone     string   `json:"phone" form:"phone"`
	Location  Location `json:"location" form:"location"`
}

type Location struct {
	Country string `json:"country" form:"country"`
	State   string `json:"state" form:"state"`
	City    string `json:"city" form:"city"`
}

type SocialMedia struct {
	Github   string `json:"github" form:"github"`
	LinkedIn string `json:"linkedin" form:"linkedin"`
}

type Summary struct {
	Description string `json:"description" form:"description"`
}

type Language struct {
	Name  string `json:"name" form:"language_name"`
	Level string `json:"level" form:"language_level"`
}

type Skill struct {
	Name        string `json:"skill_name" form:"skill_name"`
	Description string `json:"skill_description" form:"skill_description"`
	Level       uint8  `json:"skill_level" form:"skill_level"`
}

type Experience struct {
	Employer    string    `json:"employer" form:"experience_employer"`
	Role        string    `json:"role" form:"experience_role"`
	StartDate   string    `json:"start_date" form:"experience_start_date"`
	EndDate     string    `json:"end_date" form:"experience_end_date"`
	Present     bool      `json:"is_present" form:"experience_is_present"`
	Remote      bool      `json:"is_remote" form:"experience_is_remote"`
	Location    *Location `json:"location" form:"experience_location"`
	Description string    `json:"description" form:"experience_description"`
}

type Education struct {
	School         string   `json:"school" form:"education_school"`
	Field          string   `json:"field" form:"education_field"`
	GraduationDate string   `json:"graduation_date" form:"education_graduation_date"`
	Location       Location `json:"location" form:"education_location"`
	Description    string   `json:"description" form:"education_description"`
}

type Resume struct {
	Contact     Contact      `json:"contact"`
	Summary     Summary      `json:"summary"`
	SocialMedia *SocialMedia `json:"social_media"`
	Languages   *Language    `json:"languages"`
	Skills      *Skill       `json:"skills"`
	Experiences *Experience  `json:"experiences"`
	Educations  *Education   `json:"educations"`
}
