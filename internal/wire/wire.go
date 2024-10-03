package wire

import (
	"github.com/B1gdawg0/se-project-backend/config"
	"github.com/B1gdawg0/se-project-backend/internal/adapters/repositories"
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest/handlers"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure"
	"github.com/B1gdawg0/se-project-backend/internal/infrastructure/entities"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/auth"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/table"
	"github.com/B1gdawg0/se-project-backend/internal/usecases/user"
)

func InitHandler() *rest.Handler{
	config := config.ProvideConfig()
	db := infrastructure.ProvidePostGresDB(*config)

	if err := db.AutoMigrate(
		&entities.User{},
        &entities.MusicLine{},
        &entities.IGLine{},
        &entities.Table{},
        &entities.Order{},
        &entities.OrderLine{},
        &entities.Menu{},
		); 
		err != nil{ panic("Can't Automigrate Database: Auto shutdown...")}

	userRepo := repositories.ProvideUserRepository(db)
	userService := user.ProvideUserService(userRepo)
	userHandler := handlers.ProvideUserRestHandler(userService)

	authService := auth.ProvideAuthService(userRepo)
	authHandler := handlers.ProvideAuthRestHandler(authService)

	tableRepo := repositories.ProvideTableRepository(db)
	tableService := table.ProvideTableService(tableRepo)
	tableHandler := handlers.ProvideTableRestHandler(tableService)

	handler := rest.ProvideHandler(userHandler, authHandler, tableHandler)

	return handler
}