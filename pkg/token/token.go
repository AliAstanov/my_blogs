package token

import (
	"time"

	"github.com/AliAstanov/my_blogs/models"
	"github.com/golang-jwt/jwt"
)

 type Claim struct {
	UserId string
	UserRole string
	jwt.StandardClaims
}

var secretjWTKey = []byte("secret_key")

func GenerateJWT(claim models.Claim)(string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	jwtClaim := Claim{
		UserId: claim.UserId,
		UserRole: claim.UserRole,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwtClaim)
	tokenString, err := token.SignedString(secretjWTKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParsJwt(tokenString string)(*Claim, error){
	var claim = &Claim{}

	token, err := jwt.ParseWithClaims(tokenString,claim,func (t *jwt.Token)(interface{},error)  {
		return secretjWTKey, nil
		
	})
	if err != nil{
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return claim, nil
}
 

// import (
// 	"time"

// 	log "github.com/AliAstanov/my_blogs/pkg/logger"
// 	"github.com/saidamir98/udevs_pkg/logger"

// 	//"github.com/dgrijalva/jwt-go"
// 	"github.com/golang-jwt/jwt"
// )

// type JwtHandler struct {
// 	Sub       string
// 	Exp       string
// 	Iat       string
// 	Aud       []string
// 	Role      string
// 	SignedKey string
// 	Log       log.Log
// 	Token     string
// 	Timeout   int
// }

// type Customclaims struct {
// 	*jwt.Token
// 	Sub  string   `json:"sub"`
// 	Exp  string   `json:"exp"`
// 	Iat  string   `json:"iat"`
// 	Aud  []string `json:"aud"`
// 	Role string   `json:"role"`
// }

// func (j *JwtHandler) GenerateToken() (string, error) {
// 	var (
// 		token  *jwt.Token
// 		claims jwt.MapClaims
// 	)

// 	token = jwt.New(jwt.SigningMethodES256)

// 	claims = token.Claims.(jwt.MapClaims)
// 	claims["sub"] = j.Sub
// 	claims["exp"] = time.Now().Add(time.Minute * time.Duration(j.Timeout)).Unix()
// 	claims["iat"] = time.Now().Unix()
// 	claims["aud"] = j.Aud
// 	claims["user_role"] = j.Role

// 	accesToken, err = token.SignedString([]byte(j.SignedKey))
// 	if err != nil {
// 		j.Log.Error("error on generating token", logger.Error(err))
// 		return "", err
// 	}

// 	return accesToken, nil
// }
