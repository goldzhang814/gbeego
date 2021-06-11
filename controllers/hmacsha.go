package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"strconv"
	"time"
	//"net/http"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	"encoding/hex"
)

type HmacshaController struct {
	//beego.Controller
	CommonsController
}

func (c *HmacshaController) FilterOauth() {
	jsoninfo := c.GetString("timestamp")
	timestamp, err := strconv.ParseInt(jsoninfo, 10, 64)
	if err != nil {
		c.Abort("failed0")
	}
	timestampT := time.Unix(timestamp, 0)
	nowT := time.Now()
	if nowT.Sub(timestampT).Minutes() > 10 {
		fmt.Println(timestampT, "Duration is too long")
		c.Abort("failed1")
	}

	size := c.GetString("size")
	if size == "" {
		c.Abort("failed2")
	}

	sign := c.GetString("sign")
	if sign == "" {
		c.Abort("failed3")
	}

	if !CheckMAC(size+timestampT.String(), sign, "seckey") {
		fmt.Println("Verify error")
		c.Abort("failed4")
	}

	c.Abort("ok")
}

func CheckMAC(message, sign, key string) bool {
	mac := hmac.New(sha1.New, []byte(key))

	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil)) == sign
}

func (c *HmacshaController) GetMAC() {
	fmt.Println("come getMac")

	key := "seckey"

	jsoninfo := c.GetString("timestamp")
	timestamp, err := strconv.ParseInt(jsoninfo, 10, 64)
	if err != nil {
		c.Abort("failed0")
	}
	timestampT := time.Unix(timestamp, 0)
	size := "9999"
	message := size + timestampT.String()

	mac := hmac.New(sha1.New, []byte(key))

	mac.Write([]byte(message))
	//c.Ctx.WriteString(hex.EncodeToString(mac.Sum(nil)))
	c.Abort(hex.EncodeToString(mac.Sum(nil)))
	//return hex.EncodeToString(mac.Sum(nil))
}
