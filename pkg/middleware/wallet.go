package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Up2110H/e-wallet/pkg/requests"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (m *Middleware) WalletMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var x requests.X
		if err := c.ShouldBindHeader(&x); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := m.validate.Struct(&x); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := m.checkXDigest(c, x.UserId, x.Digest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.Next()
	}
}

func (m *Middleware) checkXDigest(c *gin.Context, userId int, digest string) error {
	user, err := m.store.User().GetById(userId)
	if err != nil {
		return err
	}

	key := user.Key
	body, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	if !validMAC(body, []byte(key), digest) {
		return errors.New("incorrect X-Digest")
	}

	return nil
}

func validMAC(message, key []byte, messageMAC string) bool {
	expectedMAC := MACEncrypt(message, key)
	fmt.Println(base64.StdEncoding.EncodeToString(expectedMAC))
	msgMac, err := base64.StdEncoding.DecodeString(messageMAC)
	if err != nil {
		return false
	}
	return hmac.Equal(msgMac, expectedMAC)
}

func MACEncrypt(message, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
