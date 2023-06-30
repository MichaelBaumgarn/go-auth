package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/gin-gonic/gin"
)

const TokenSecret = `-----BEGIN RSA PRIVATE KEY-----
 MIICXgIBAAKBgQCfRbf5Z8vqWtP++zinsIsLgfWjaqfyHJx8MZuGnZcphrRBffrT
 y57anScrLLV7C2fS8eQG923nanhpfDR4DYLlYsPY3OE7j7FuRUBm320JXH/Z2gyr
 9i9unqgmY7OcT/KauKtqXi2RmktE9eN2vGwVb2x4ekcBGhfMkbzQh3n9CwIDAQAB
 AoGBAIdlHSBHvHFdBRqdNWxYB9ugftCMunh3Gg/5m9yz2uzSNin2jmO9lS0Cq2lo
 eF5BM+F+//lsHWG8gOArVe84jSxXWSSAJnm/AZIjr3BwwotbpMHAsrgYhQ86KR17
 k5wN5Snp5ScgR/7DRghoDmcnjBbqC2n9BBoLPgKtMxpybJYBAkEA7E1g2gv1WcZI
 ns0uhHO8m36KK8H3wgaEDuZvE6icmW2AYdZw8nDPpNkSijmyukBMGWvB81/wfdcz
 wRXIHyS4YwJBAKyMi03rzLqSZumkF7YNdmAHywYg+iThaxRseFtvBp31VWF/K68o
 /CK2hI1gw1HHgs0Ylronze4LTcWr2Z9xBTkCQQDK69vrA2/raxo7vKlgtt7VjQHs
 d8JiTSQkg4AJmrb4Do+79OHDYFsADiUcrFWuGb7/6YiBjfbAqghYnHLhQ7BPAkEA
 j4bwsUY9K8xv0lTubD7Sgpq45EY2DMt8+KYpj1TRGj9iYBEaUz8yS+WqaLdegP4F
 7XhQmFRX1XSNoqmKAjhO8QJAeXoILM1hmgxXjwGy9wAldyS6QwABxiJHHVYRKxdG
 Y5gW3mgEMMk8dvRLr0vL3xKdAeaPbF8MNVzGlo09pyPgeQ==
 -----END RSA PRIVATE KEY-----`

// type Config struct {
// 	db *gorm.DB
// }

func (e Env) GetUser(c *gin.Context) {
	users := e.users.GetAll()

	c.IndentedJSON(http.StatusOK, users)
}

func (e Env) PostUser(c *gin.Context) {
	var incomingUser User
	if err := c.BindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: err.Error()})
		return
	}
	e.users.Create(incomingUser)

	fmt.Printf("%v", &incomingUser)

	c.IndentedJSON(http.StatusOK, &incomingUser)
}

func (e Env) Login(c *gin.Context) {
	var incomingUser User
	if err := c.BindJSON(&incomingUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Error: err.Error()})
		return
	}

	var token string
	var err error
	if token, err = e.users.Authenticate(incomingUser); err != nil {
		c.JSON(http.StatusUnauthorized, Response{Code: 400, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Data: fmt.Sprintf("Bearer %s", token)})
}

func (e Env) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	// user := UserModel{db:}.ModelGetUserByID(id, config)
	user := e.users.GetByID(id)

	c.IndentedJSON(http.StatusOK, user)
}

func GenerateToken(user User) (string, error) {
	// key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	now := time.Now()
	expiry := time.Now().Add(time.Hour * 24 * 4)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{ID: user.ID, ExpieresAt: expiry.Unix(), IssuedAt: now.Unix()})
	// crypto.Signer()
	tokenString, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		panic(err)
	}

	return tokenString, err
}
