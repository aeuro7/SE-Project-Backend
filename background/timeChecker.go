package background

import (
	"fmt"
	"log"
	"time"

	"github.com/B1gdawg0/se-project-backend/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func DailyMonitor(app *fiber.App, handler *rest.Handler){
	for{
		fmt.Println("Monitor Daily Check!")
		now := time.Now()

		if now.Hour() >= 2 && now.Hour() <=5{
			c := app.AcquireCtx(&fasthttp.RequestCtx{})
			if err := handler.Table.ClearTablesDaily(c); err!=nil{
				log.Fatal("can't clear database")
			}
		}

		time.Sleep(time.Hour)
	}
}