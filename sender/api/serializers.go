package api

import (
	"github.com/gin-gonic/gin"
	"sender/models"
	"strconv"
	"strings"
)

func CreateNotificationSerializer(c *gin.Context) (*models.Notification, bool) {
	// Convert bytes to <Notification> structure
	newNotification := &models.Notification{}
	if err := c.ShouldBindJSON(newNotification); err != nil {
		return newNotification, false
	}

	// E-mails validate
	if !emailsValidator(newNotification.To) {
		return newNotification, false
	}
	return newNotification, true
}

type GetNotificationsRequestData struct {
	to     []string
	limit  int64
	offset int64
}

func GetNotificationsSerializer(c *gin.Context) (*GetNotificationsRequestData, bool) {
	data := &GetNotificationsRequestData{}

	// Validate <to> param
	var to []string
	toParam := c.DefaultQuery("to", "")
	if len(toParam) != 0 {
		to = strings.Split(toParam, ",")
	}
	if !emailsValidator(to) {
		return data, false
	}

	// Validate <limit> and <offset> param
	limit, err := strconv.ParseInt(
		c.DefaultQuery("limit", "10"),
		10,
		64,
	)
	offset, err := strconv.ParseInt(
		c.DefaultQuery("offset", "0"),
		10,
		64,
	)
	if err != nil {
		return data, false
	}

	// Validate <limit> and <offset> values
	if (limit > 100 || limit < 0) || (offset < 0) {
		return data, false
	}

	// Data packing
	data.to = to
	data.limit = limit
	data.offset = offset

	return data, true
}
