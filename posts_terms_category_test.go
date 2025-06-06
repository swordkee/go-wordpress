package wordpress_test

import (
	"context"
	"github.com/swordkee/go-wordpress"
	"net/http"
	"testing"
)

func cleanUpPostsTermsCategory(t *testing.T, postID int, id int) {

	wp, ctx := initTestClient()
	// terms does not support trashing
	deletedTerm, resp, err := wp.Posts.Entity(postID).Terms().Category().Delete(ctx, id, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new term: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedTerm.ID != id {
		t.Errorf("Deleted term ID should be the same as newly created term: %v != %v", deletedTerm.ID, id)
	}
}

func getAnyOnePostsTermsCategory(t *testing.T, ctx context.Context, wp *wordpress.Client, postID int) *wordpress.PostsTerm {

	terms, resp, _ := wp.Posts.Entity(postID).Terms().Category().List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(terms) < 1 {
		t.Fatalf("Should not return empty terms")
	}

	id := terms[0].ID

	term, resp, _ := wp.Posts.Entity(postID).Terms().Category().Get(ctx, id, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return term
}

func TestPostsTermsCategory_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Post object manually to retrieve PostsTermsService
	// A proper API call would inject the right PostsTermsService, Client and other goodies into a post,
	// allowing user to call post.Terms()
	invalidPost := wordpress.Post{}
	invalidTerms := invalidPost.Terms()
	if invalidTerms != nil {
		t.Errorf("Expected meta to be nil, %v", invalidTerms)
	}
}

func TestPostsTermsCategoryList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp, ctx := initTestClient()
	post := getAnyOnePost(t, ctx, wp)
	postID := post.ID

	terms, resp, err := wp.Posts.Entity(postID).Terms().Category().List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if terms == nil {
		t.Errorf("Should not return nil terms")
	}
}

func TestPostsTermsCategoryGet(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp, ctx := initTestClient()
	post := getAnyOnePost(t, ctx, wp)
	postID := post.ID
	tt := getAnyOnePostsTermsCategory(t, ctx, wp, postID)

	term, resp, err := wp.Posts.Entity(postID).Terms().Category().Get(ctx, tt.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

}

func TestPostsTermsCategoryCreate_Existing(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp, ctx := initTestClient()
	post := getAnyOnePost(t, ctx, wp)
	tt := getAnyOneTermsCategory(t, ctx, wp)
	postID := post.ID
	termID := tt.ID

	term, resp, err := wp.Posts.Entity(postID).Terms().Category().Create(ctx, termID)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

	// clean up
	cleanUpPostsTermsCategory(t, postID, term.ID)
}

func TestPostsTermsCategoryDelete(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp, ctx := initTestClient()
	post := getAnyOnePost(t, ctx, wp)
	tt := getAnyOneTermsCategory(t, ctx, wp)
	postID := post.ID
	termID := tt.ID

	// create category
	newTerm, resp, _ := wp.Posts.Entity(postID).Terms().Category().Create(ctx, termID)
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// delete category
	// Note: Terms does not support trashing; `force=true` is required
	deletedTerm, resp, err := wp.Posts.Entity(postID).Terms().Category().Delete(ctx, newTerm.ID, "force=true")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedTerm == nil {
		t.Errorf("Should not return nil deletedTerm")
	}
}
