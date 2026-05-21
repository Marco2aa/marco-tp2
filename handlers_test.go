package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCat(t *testing.T) {
	catsDatabase = map[string]Cat{
		"550e8400-e29b-41d4-a716-446655440000": {Name: "Toto", Color: "Grey", BirthDate: "2023-04-16"},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/cats/550e8400-e29b-41d4-a716-446655440000", nil)
	req.SetPathValue("catId", "550e8400-e29b-41d4-a716-446655440000")

	code, body := getCat(req)
	if code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, code)
	}

	cat, ok := body.(Cat)
	if !ok {
		t.Fatalf("expected Cat response, got %T", body)
	}
	if cat.ID != "550e8400-e29b-41d4-a716-446655440000" || cat.Name != "Toto" || cat.Color != "Grey" {
		t.Fatalf("unexpected cat response: %#v", cat)
	}
}

func TestGetCatNotFound(t *testing.T) {
	catsDatabase = map[string]Cat{}

	req := httptest.NewRequest(http.MethodGet, "/api/cats/missing", nil)
	req.SetPathValue("catId", "missing")

	code, body := getCat(req)
	if code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, code)
	}

	if body != "Cat not found" {
		t.Fatalf("unexpected body: %#v", body)
	}
}

func TestDeleteCat(t *testing.T) {
	catsDatabase = map[string]Cat{
		"550e8400-e29b-41d4-a716-446655440000": {Name: "Toto"},
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/cats/550e8400-e29b-41d4-a716-446655440000", nil)
	req.SetPathValue("catId", "550e8400-e29b-41d4-a716-446655440000")

	code, _ := deleteCat(req)
	if code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, code)
	}

	if _, found := catsDatabase["550e8400-e29b-41d4-a716-446655440000"]; found {
		t.Fatal("expected cat to be deleted")
	}
}

func TestDeleteCatNotFound(t *testing.T) {
	catsDatabase = map[string]Cat{}

	req := httptest.NewRequest(http.MethodDelete, "/api/cats/missing", nil)
	req.SetPathValue("catId", "missing")

	code, body := deleteCat(req)
	if code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, code)
	}

	if body != "Cat not found" {
		t.Fatalf("unexpected body: %#v", body)
	}
}
