package routes

import (
	controllers "notes_app/controllers"

	gin "github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Welcome)
	router.GET("/notes", controllers.GetAllNotes)
	router.POST("/note", controllers.CreateNote)
	router.GET("/note/:noteId", controllers.GetSingleNote)
	router.PUT("/note/:noteId", controllers.EditNote)
	router.DELETE("/note/:noteId", controllers.DeleteNote)
}
