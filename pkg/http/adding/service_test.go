package adding

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	exitVal := m.Run()
	os.Exit(exitVal)
}

func initEngine(t *testing.T) (*MockDataAccessor, *gin.Engine) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockDataAccessor(ctrl)
	engine := gin.New()
	api := engine.Group("/")
	NewService(m, api)
	return m, engine
}

func TestCreateCustomer(t *testing.T) {

	m, engine := initEngine(t)

	expected := Customer{
		Name: "Customer 1",
	}

	m.EXPECT().
		CreateCustomer(gomock.Any(), gomock.Any()).
		Return(expected, nil)

	// make object for submission
	p := struct {
		Name string
	}{
		Name: expected.Name,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(p)

	req, _ := http.NewRequest(http.MethodPost, "/sites", &buf)
	addDefaultHeaders(req)

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected to get status %d but instead got %d: %s\n", http.StatusCreated, w.Code, w.Body)
	}
	var res Customer
	json.Unmarshal(w.Body.Bytes(), &res)
	if res.Name != expected.Name {
		t.Fatal("names do not match")
	}
}

// set the proper header for performing the request
func addDefaultHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", getToken())
}

// generate test token
func getToken() string {
	mySigningKey := []byte("dev01")
	claims := auth.CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 1500,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	return "bearer " + ss
}
