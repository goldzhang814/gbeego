package controllers
import(
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
    "github.com/spf13/viper"
    "errors"
)

type CommonsController struct {
	beego.Controller
}


func CreateToken(user *UserInfo)(tokenss string,err error){
	//自定义claim
  claim := jwt.MapClaims{
	  "id":       user.ID,
	  "username": user.Username,
	  "nbf":      time.Now().Unix(),
      "iat":      time.Now().Unix(),
      "exp":       time.Now().Unix()+20*24*60*60,
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
  tokenss,err  = token.SignedString([]byte(viper.GetString("token.secret"))) 
  return
}

func secret()jwt.Keyfunc{
    return func(token *jwt.Token) (interface{}, error) {
        return []byte(viper.GetString("token.secret")),nil
    }
}


func  ParseToken(tokenss string)(user *UserInfo,err error){
    user = &UserInfo{}
    token,err := jwt.Parse(tokenss,secret())
    if err != nil{
        return
    }
    claim,ok := token.Claims.(jwt.MapClaims)
    if !ok{
        err = errors.New("cannot convert claim to mapclaim")
        return
    }
    //验证token，如果token被修改过则为false
    if  !token.Valid{
        err = errors.New("token is invalid")
        return
    }

    user.ID =uint64( claim["id"].(float64))
    user.Username = claim["username"].(string)
    user.Nbf = int64(claim["nbf"].(float64))
    user.Iat = int64(claim["iat"].(float64))
    user.Exp = int64(claim["exp"].(float64))
    return
}