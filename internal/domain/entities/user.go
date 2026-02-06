package entities

type User struct {
	Base `gorm:"embedded"`

	Email string  `gorm:"uniqueIndex;not null" json:"email"`
	Phone *string `gorm:"uniqueIndex" json:"phone,omitempty"`

	Name           *string `json:"name,omitempty"`
	ProfilePicture *string `json:"profile_picture,omitempty"`

	Password *string `json:"-"`

	Reminders []*Reminder `json:"reminders,omitempty" gorm:"constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "Users"
}
