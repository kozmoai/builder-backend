package kozmomarketplacesdk

import "github.com/kozmoai/builder-backend/src/model"

type AppForMarketplace struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewAppForMarketplace(app *model.App) *AppForMarketplace {
	appConfig := app.ExportConfig()
	return &AppForMarketplace{
		Name:        app.Name,
		Description: appConfig.Description,
	}
}
