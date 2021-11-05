package controllers

import (
	"log"
	"net/http"
	"time"

	. "notes_app/models"

	gin "github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

// db instance to make db changes
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

// Create User Table
func CreateNotesTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&Note{}, opts)

	if createError != nil {
		log.Printf("Error while creating Notes table, Reason: %v\n", createError)
		return createError
	}

	log.Printf("Notes table created")
	return nil
}

// base route
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
}

func GetAllNotes(c *gin.Context) {
	var notes []Note

	err := dbConnect.Model(&notes).Select()
	if err != nil {
		log.Printf("Error while getting all notes, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Sucessfully found all notes",
		"data":    notes,
	})
}

func CreateNote(c *gin.Context) {

	// create a Note variable and bind it to the JSON request body
	var note Note
	c.BindJSON(&note)

	title := note.Title
	body := note.Body
	noteType := note.Type
	completed := false
	id := guuid.New().String()

	if noteType != "HOME" && noteType != "PERSONAL" && noteType != "WORK" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please select a valid type for the note",
		})
		return
	}

	insertError := dbConnect.Insert(&Note{
		ID:        id,
		Title:     title,
		Body:      body,
		Type:      noteType,
		Completed: completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if insertError != nil {
		log.Printf("Error while inserting new note into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Note created Successfully",
	})
}

func GetSingleNote(c *gin.Context) {
	noteId := c.Param("noteId")
	note := &Note{ID: noteId}

	err := dbConnect.Select(note)

	if err != nil {
		log.Printf("Error while getting a single note, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Found Note with ID " + noteId,
		"data":    note,
	})
}

func EditNote(c *gin.Context) {
	noteId := c.Param("noteId")
	var note Note
	c.BindJSON(&note)
	completed := note.Completed

	_, err := dbConnect.Model(&Note{}).Set("completed = ?", completed).Where("id = ?", noteId).Update()

	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Note Edited Successfully",
	})
}

func DeleteNote(c *gin.Context) {
	noteId := c.Param("noteId")
	todo := &Note{ID: noteId}

	err := dbConnect.Delete(todo)
	if err != nil {
		log.Printf("Error while deleting a single note, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Note deleted successfully",
	})
}
