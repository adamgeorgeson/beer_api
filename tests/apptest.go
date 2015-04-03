package tests

import "github.com/revel/revel/testing"

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestThatBeerListPageWorks() {
	t.Get("/beers")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	t.AssertContains("La Trappe Quadrupel Oak Aged")
}

func (t *AppTest) TestThatBeerShowPageWorks() {
	t.Get("/beers/2")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	t.AssertContains("BrewDog")
}

func (t *AppTest) TestInvalidBeerID() {
	t.Get("/beers/999")
	t.AssertNotFound()
	t.AssertContentType("text/html; charset=utf-8")
	t.AssertContains("Could not find beer")
}

func (t *AppTest) After() {
	println("Tear down")
}
