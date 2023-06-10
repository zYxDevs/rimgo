package pages

import (
	"codeberg.org/rimgo/rimgo/api"
	"codeberg.org/rimgo/rimgo/utils"
)

var ApiClient *api.Client

func InitializeApiClient() {
    ApiClient = api.NewClient(utils.Config.ImgurId)
}