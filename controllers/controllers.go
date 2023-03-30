package controllers

import (
	"io"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"github.com/lucassauro/go-api-boilerplate/models"
	"github.com/lucassauro/go-api-boilerplate/structs"
)

func Echo(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Writer.Header().Add("allow", "GET, POST")
		return
	}

	if c.Request.Method == http.MethodPost {
		ct := c.Request.Header.Values("Content-Type")
		if len(ct) != 0 {
			if ct[0] != "application/json" && ct[0] != "text/plain" {
				c.String(http.StatusUnsupportedMediaType, "Unsupported media type")
				return
			}
		}

		str, err := io.ReadAll(c.Request.Body)
		if err != nil {
			var e structs.Erro
			e.Err = err
			c.JSON(http.StatusInternalServerError, e)
			return
		}

		if len(str) != 0 {
			var e structs.Echo
			e.Echo = string(str)
			json.Unmarshal(str, &e)

			accept := c.Request.Header.Values("Accept")
			switch accept[0] {
			case "*/*", "text/plain":
				c.String(http.StatusOK, e.Echo)
				return
			case "application/json":
				c.JSON(http.StatusOK, e)
				return
			case "application/xml":
				c.XML(http.StatusOK, e)
			default:
				c.String(http.StatusNotAcceptable, "Cannot honor Accept media type")
				return
			}
		}
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func TestDB(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Writer.Header().Add("allow", "GET")
		return
	}

	s := models.TestDB()

	// https://stackoverflow.com/questions/28447297/how-to-check-for-an-empty-struct
	if s == (structs.DBItem{}) {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, s)
}


func TestsDB(c *gin.Context) {
	if c.Request.Method == http.MethodOptions {
		c.Writer.Header().Add("allow", "GET")
		return
	}

	s := models.TestsDB()
	
	if len(s) < 1 {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, s)
}