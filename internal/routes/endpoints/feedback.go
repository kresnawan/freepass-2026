package endpoints

import (
	"canteen/internal/handlers/owner"
	"canteen/internal/middleware"

	"github.com/gin-gonic/gin"
)

func FeedbackEndpointsGroup(rg *gin.RouterGroup) {
	FeedbackEndpoint := rg.Group("/feedback")
	{
		OwnerField := FeedbackEndpoint.Group("")
		OwnerField.Use(middleware.OwnerAuth())
		OwnerField.Use(middleware.CheckOwnerOwnership())
		{
			OwnerField.DELETE("/:fid", owner.DeleteFeedbackById)

		}
	}

}
