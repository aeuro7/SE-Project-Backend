package router

import (
	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/B1gdawg0/se-project-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRouter(app *fiber.App, rqHandler *rest.Handler) {
	user := app.Group("/users")
	auth := app.Group("/auth")
	table := app.Group("/tables")
	order := app.Group("/orders")
	menu := app.Group("/menu")
	orderLine := app.Group("/order-lines")
	admin := app.Group("/admin")
	igLine := app.Group("/ig-lines", middleware.CheckJWT)
	musicLine := app.Group("/music-lines", middleware.CheckJWT)
	discount := app.Group("/discounts")

	auth.Post("/login", rqHandler.Auth.Login)
	auth.Post("/register", rqHandler.Auth.Register)

	user.Use(middleware.CheckJWT)
	

	user.Get("", rqHandler.User.GetUsers)
	user.Get("/id=:id", rqHandler.User.GetUserByID)
	user.Get("/email=:email", rqHandler.User.GetUserByEmail)
	user.Get("/phone=:phone", rqHandler.User.GetCustomerByPhone)
	user.Post("", rqHandler.User.CreateUser)

	table.Get("", rqHandler.Table.GetTables)
	table.Get("/id=:id", rqHandler.Table.GetTableByID)
	table.Post("", rqHandler.Table.CreateTable)
	table.Put("id=:id", rqHandler.Table.UpdateTableByID)
	table.Delete("/id=:id", rqHandler.Table.DeleteTableByID)

	order.Get("", rqHandler.Order.GetAllOrder)
	order.Get("/id=:id", rqHandler.Order.GetOrderByID)
	order.Get("/t_id=:id", rqHandler.Order.GetOrderByTableID)
	order.Post("", rqHandler.Order.CreateOrderByID)
	order.Post("/w-olines", rqHandler.Order.CreateOrderWithOrderLines)

	menu.Get("", rqHandler.Menu.GetAllMenu)
	menu.Get("/id=:id", rqHandler.Menu.GetMenuByID)
	menu.Post("", rqHandler.Menu.CreateMenu)
	menu.Put("/id=:id", rqHandler.Menu.UpdateMenuByID)
	menu.Delete("/id=:id", rqHandler.Menu.DeleteMenuByID)
  
	orderLine.Get("",rqHandler.OrderLine.GetOrderLines)
	orderLine.Get("/id=:id",rqHandler.OrderLine.GetOrderLineByID)
	orderLine.Get("/o_id=:id",rqHandler.OrderLine.GetOrderLinesByOrderID)
	orderLine.Post("", rqHandler.OrderLine.CreateOrderLine)

	admin.Get("/users", rqHandler.User.GetUsers)
	admin.Get("/users/id=:id", rqHandler.User.GetUserByID)
	admin.Get("/users/email=:email", rqHandler.User.GetUserByEmail)

	admin.Get("/tables", rqHandler.Table.GetTables)
	admin.Get("/tables/id=:id", rqHandler.Table.GetTableByID)

	admin.Get("/orders", rqHandler.Order.GetAllOrder)
	admin.Get("/orders/id=:id", rqHandler.Order.GetOrderByID)

	admin.Get("/order-lines", rqHandler.OrderLine.GetOrderLines)
	admin.Get("/order-lines/id=:id", rqHandler.OrderLine.GetOrderLineByID)

	igLine.Post("", rqHandler.IgLine.CreateIgLine)
	igLine.Get("", rqHandler.IgLine.FindAllIgLine)

	musicLine.Post("", rqHandler.Musicline.CreateMusicLine)
	musicLine.Get("", rqHandler.Musicline.FindAllMusicLine)

	discount.Get("", rqHandler.Discount.GetAllDiscount)
	discount.Get("/id=:id", rqHandler.Discount.GetDiscountByID)
	discount.Post("", rqHandler.Discount.CreateDiscount)
}
