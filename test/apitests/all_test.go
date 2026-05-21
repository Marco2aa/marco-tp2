package apitests

import (
	"fmt"
	"net/http"
	"testing"
)

func createTestCat(t *testing.T) string {
	t.Helper()

	code := 0
	catID := ""
	err := call("POST", "/cats", &CatModel{
		Name:      "Toto",
		Color:     "Grey",
		BirthDate: "2023-04-16",
	}, &code, &catID)
	if err != nil {
		t.Fatalf("POST /cats request error: %v", err)
	}

	fmt.Println("POST /cats ->", code, catID)

	if code != http.StatusCreated {
		t.Fatalf("POST /cats should return 201, got %d", code)
	}
	if catID == "" {
		t.Fatal("POST /cats should return the created cat ID")
	}

	return catID
}

func deleteTestCat(t *testing.T, catID string) {
	t.Helper()

	code := 0
	err := call("DELETE", "/cats/"+catID, nil, &code, nil)
	if err != nil {
		t.Fatalf("DELETE /cats/%s request error: %v", catID, err)
	}

	fmt.Println("DELETE /cats/"+catID+" ->", code)

	if code != http.StatusNoContent {
		t.Fatalf("DELETE /cats/%s should return 204, got %d", catID, code)
	}
}

func contains(items []string, expected string) bool {
	for _, item := range items {
		if item == expected {
			return true
		}
	}
	return false
}

func TestGetCats(t *testing.T) {
	catID := createTestCat(t)
	defer deleteTestCat(t, catID)

	code := 0
	result := []string{}
	err := call("GET", "/cats", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats ->", code, result)

	if code != http.StatusOK {
		t.Error("We should get code 200, got", code)
	}

	if !contains(result, catID) {
		t.Error("Listing the IDs, expected to find", catID, "got", result)
	}
}

func TestCreateCat(t *testing.T) {
	catID := createTestCat(t)
	defer deleteTestCat(t, catID)

	if catID == "" {
		t.Fatal("Cat ID should not be empty")
	}
}

func TestGetCat(t *testing.T) {
	catID := createTestCat(t)
	defer deleteTestCat(t, catID)

	code := 0
	result := CatModel{}
	err := call("GET", "/cats/"+catID, nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats/"+catID+" ->", code, result)

	if code != http.StatusOK {
		t.Fatalf("GET /cats/%s should return 200, got %d", catID, code)
	}
	if result.ID != catID {
		t.Fatalf("GET /cats/%s should return cat ID %s, got %s", catID, catID, result.ID)
	}
	if result.Name != "Toto" || result.Color != "Grey" || result.BirthDate != "2023-04-16" {
		t.Fatalf("Unexpected cat details: %#v", result)
	}
}

func TestGetCatNotFound(t *testing.T) {
	code := 0
	result := ""
	err := call("GET", "/cats/missing-cat", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats/missing-cat ->", code, result)

	if code != http.StatusNotFound {
		t.Fatalf("GET /cats/missing-cat should return 404, got %d", code)
	}
	if result != "Cat not found" {
		t.Fatalf("Unexpected response body: %q", result)
	}
}

func TestDeleteCat(t *testing.T) {
	catID := createTestCat(t)

	deleteTestCat(t, catID)

	code := 0
	result := ""
	err := call("GET", "/cats/"+catID, nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("GET /cats/"+catID+" after delete ->", code, result)

	if code != http.StatusNotFound {
		t.Fatalf("GET /cats/%s after delete should return 404, got %d", catID, code)
	}
}

func TestDeleteCatNotFound(t *testing.T) {
	code := 0
	result := ""
	err := call("DELETE", "/cats/missing-cat", nil, &code, &result)
	if err != nil {
		t.Error("Request error", err)
	}

	fmt.Println("DELETE /cats/missing-cat ->", code, result)

	if code != http.StatusNotFound {
		t.Fatalf("DELETE /cats/missing-cat should return 404, got %d", code)
	}
	if result != "Cat not found" {
		t.Fatalf("Unexpected response body: %q", result)
	}
}
