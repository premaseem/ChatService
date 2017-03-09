package main


import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/julienschmidt/httprouter"
)

const (
	tickMark = "\u2714"
	ballotX  = "\u2716"
)

func TestBasicChat(t *testing.T) {

	router := httprouter.New()
	router.POST("/",Index)
	req,err := http.NewRequest("POST","/",nil)
	if err != nil {
		t.Error("Applicaton root is not working properly",ballotX )
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w,req)
	fmt.Print(w.Code)
	fmt.Print(http.StatusOK)



}