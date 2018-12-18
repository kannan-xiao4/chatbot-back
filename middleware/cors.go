package middleware

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

var reg = regexp.MustCompile("")

func CORS(c *gin.Context) {

	if c.Request.Method == "OPTIONS" {
		// for preflight
		//origin := c.Request.Header.Get("Origin")

		//r := reg.Copy()
		//if r.MatchString(origin) {
		if true {
			headers := c.Request.Header.Get("Access-Control-Request-Headers")

			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
			//c.Writer.Header().Set("Access-Control-Allow-Headers", "")

			c.Data(200, "text/plain", []byte{})
			c.Abort()
		} else {
			c.Data(403, "text/plain", []byte{})
			c.Abort()
		}
	} else {
		// for actual response
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Expose-Headers", "")
		c.Next()
	}

	return
}
