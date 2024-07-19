package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"todo/database/models"
)

func getAllTasksHandler(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not fetch tasks, try again later."},
		)
		return
	}

	context.JSON(http.StatusOK, tasks)
}

func getTaskHandler(context *gin.Context) {
	taskID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the task id"})
	}
	task, err := models.GetTask(taskID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the task"})
		return
	}
	context.JSON(http.StatusOK, task)
}

func createTaskHandler(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse requested data."})
		return
	}

	fmt.Printf("Task before save: %+v\n", task)

	err = task.Save()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not create task, please try again later."},
		)
		return
	}

	fmt.Printf("Task after save: %+v\n", task)

	context.JSON(http.StatusCreated, gin.H{"message": "task created successfuly.", "task": task})
}

func updateTaskHandler(context *gin.Context) {
	taskID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the task id"})
	}

	_, err = models.GetTask(taskID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the task"})
		return
	}

	var updatedTask models.Task
	err = context.ShouldBindJSON(&updatedTask)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse task id"})
		return
	}

	updatedTask.ID = taskID
	err = updatedTask.Update()
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not update task, try again later."},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "task updated successfuly"})
}

func deleteTaskHandler(context *gin.Context) {
	taskID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the task id"})
	}

	task, err := models.GetTask(taskID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the task"})
		return
	}

	err = task.Delete()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "could not delete task, try again later."},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "task deleted successfuly."})
}
