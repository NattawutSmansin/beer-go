package helper

import (
	"beer/enum"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseJson(c *gin.Context, statusCode enum.ResponseCode, data interface{}, messages ...string) {
	var response gin.H

	switch statusCode {
	case enum.Success, enum.Created, enum.Accepted, enum.Deleted:
		response = gin.H{
			"status": true,
			"code":   statusCode,
			"data":   data,
		}
	case enum.Fail:
		response = gin.H{
			"status": false,
			"code":   statusCode,
			"error": gin.H{
				"message": messages,
			},
		}
	case enum.NotFound:
		response = gin.H{
			"status": false,
			"code":   statusCode,
			"error": gin.H{
				"message": "Not found",
			},
		}
	case enum.Validate:
		response = gin.H{
			"status": false,
			"code":   statusCode,
			"error": gin.H{
				"message": "Validation error",
				"field":   data,
			},
		}
	case enum.Unauthorized:
		response = gin.H{
			"status": false,
			"code":   statusCode,
			"error": gin.H{
				"message": "Unauthorized to access the resource",
			},
		}
	case enum.Error, enum.Forbidden:
		response = gin.H{
			"status": false,
			"code":   statusCode,
			"error": gin.H{
				"message": messages,
			},
		}
	case enum.ManyRequest:
		response = gin.H{
			"status": false,
			"code":   statusCode,
			"error": gin.H{
				"message": "Too Many Requests",
			},
		}
	default:
		response = gin.H{
			"status": false,
			"code":   http.StatusInternalServerError,
			"error": gin.H{
				"message": "Internal Server Error",
			},
		}
	}

	c.JSON(int(statusCode), response)
}

func ResponseJsonPaginate(c *gin.Context, statusCode enum.ResponseCode, data interface{}, page,limit,total int) {
	response := gin.H{
		"status": true,
		"code":   statusCode,
		"page":   page,
		"limit":  limit,
		"total":  total,
		"data":   data,
	}

	c.JSON(int(statusCode), response)
}
