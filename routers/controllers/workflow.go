package controllers

import (
	"github.com/cloudreve/Cloudreve/v4/pkg/hashid"
	"github.com/cloudreve/Cloudreve/v4/pkg/queue"
	"github.com/cloudreve/Cloudreve/v4/pkg/serializer"
	"github.com/cloudreve/Cloudreve/v4/service/explorer"
	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	service := ParametersFromContext[*explorer.ListTaskService](c, explorer.ListTaskParamCtx{})
	resp, err := service.ListTasks(c)
	if err != nil {
		c.JSON(200, serializer.Err(c, err))
		c.Abort()
		return
	}

	if resp != nil {
		c.JSON(200, serializer.Response{
			Data: resp,
		})
	}
}

func GetTaskPhaseProgress(c *gin.Context) {
	taskId := hashid.FromContext(c)
	resp, err := explorer.TaskPhaseProgress(c, taskId)
	if err != nil {
		c.JSON(200, serializer.Err(c, err))
		c.Abort()
		return
	}

	if resp != nil {
		c.JSON(200, serializer.Response{
			Data: resp,
		})
	} else {
		c.JSON(200, serializer.Response{Data: queue.Progresses{}})
	}
}

func SetDownloadTaskTarget(c *gin.Context) {
	taskId := hashid.FromContext(c)
	service := ParametersFromContext[*explorer.SetDownloadFilesService](c, explorer.SetDownloadFilesParamCtx{})
	err := service.SetDownloadFiles(c, taskId)
	if err != nil {
		c.JSON(200, serializer.Err(c, err))
		c.Abort()
		return
	}

	c.JSON(200, serializer.Response{})
}

func CancelDownloadTask(c *gin.Context) {
	taskId := hashid.FromContext(c)
	err := explorer.CancelDownloadTask(c, taskId)
	if err != nil {
		c.JSON(200, serializer.Err(c, err))
		c.Abort()
		return
	}

	c.JSON(200, serializer.Response{})
}
