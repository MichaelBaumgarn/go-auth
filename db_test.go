package main

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var env Env

func init() {
	dsn := "host=localhost dbname=vocab_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("%v", err)
		panic("db connection failed")
	}

	env = Env{users: User{db: db}, grammar: Grammar{db: db}, userGrammar: UserGrammar{db: db}}
}

func TestMain(t *testing.T) {
	testEmail := "testEmail"
	newUser := env.users.Create(User{Email: testEmail, Password: "flkjsal"})
	fmt.Printf("check newUser %v", newUser.ID)

	user := env.users.GetByID(fmt.Sprint(newUser.ID))

	if testEmail != user.Email {
		t.Errorf("email should be %v, is %v", testEmail, user.Email)
	}

	env.users.db.Delete(user, user.ID)
}

func TestLogin(t *testing.T) {
	testEmail := "uniqueEmail123"
	newUser := env.users.Create(User{Email: testEmail, Password: "flkjsal"})

	if _token, err := env.users.Authenticate(newUser); err != nil || len(_token) < 10 {
		t.Errorf("no token was generated %v %v ", err, _token)
	}

	token, err := GenerateToken(newUser)
	if err != nil {
		t.Errorf("alaram %v %v ", err, token)
	}
	fmt.Printf("token %v", token)

	if claim, err := ParseToken(token, TokenSecret); err != nil {
		t.Errorf("no token was generated %v %v ", err, claim)
	} else {
		fmt.Printf("show claim %v", claim)

	}

	env.users.db.Delete(newUser, newUser.ID)
}

func TestGrammar(t *testing.T) {
	type Result struct {
		Word     string
		Email    string
		Complete string
	}
	testEmail := "uniqueEmail123"
	newUser := env.users.Create(User{Email: testEmail, Password: "flkjsal"})
	fmt.Printf("show %v", newUser)

	newGrammar := env.grammar.Create(Grammar{
		Language: "english",
		Index:    0,
		Word:     "ran",
		Complete: "somewhat",
	})
	newUserGrammar := env.userGrammar.Create(UserGrammar{
		UserId:    int(newUser.ID),
		GrammarId: int(newGrammar.ID),
	})
	fmt.Printf("newUserGrammar %v\n ", newUserGrammar)

	var result Result
	userGrammarQuery := `
		join users on users.user_id = user_grammar.user_id 
		join grammar on grammar.grammar_id = user_grammar.grammar_id`
	env.userGrammar.db.Model(UserGrammar{}).Select("word, complete, email").Joins(userGrammarQuery).Scan(&result)
	fmt.Printf("final result  %v", result)
	if result.Email != newUser.Email || result.Word != newGrammar.Word {

		t.Errorf("final result %v  ", result)
	}

	env.users.db.Delete(newUser, newUser.ID)
	env.grammar.db.Delete(newGrammar, newGrammar.ID)
	env.userGrammar.db.Delete(newUserGrammar, newUserGrammar.ID)
}
