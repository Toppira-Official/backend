package main

import (
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/shared/repositories",
		Mode: gen.WithDefaultQuery |
			gen.WithQueryInterface,
	})

	g.ApplyBasic(
		entities.User{},
		entities.Reminder{},
	)

	g.Execute()
}
