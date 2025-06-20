package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"myapp/config"
	"myapp/middleware"
	"myapp/model"
)

func withAuthContext(ctx context.Context, admin *middleware.User) context.Context {
	return context.WithValue(ctx, middleware.CtxKey, admin)
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error reading .env file: %q", err)
	}

	config.ConnectDB()

	code := m.Run()
	os.Exit(code)
}

func TestPostCreate(t *testing.T) {
	db := config.GetDB()
	require.NotNil(t, db)

	s := &Service{DB: db}
	ctx := withAuthContext(context.Background(), &middleware.User{ID: 1})

	tests := []struct {
		name    string
		input   *model.NewPost
		wantErr bool
	}{
		{
			name:    "valid post",
			input:   &model.NewPost{Title: "Valid Post", Content: "This is a valid post content"},
			wantErr: false,
		},
		{
			name:    "missing title",
			input:   &model.NewPost{Title: "", Content: "Missing title"},
			wantErr: true,
		},
		{
			name:    "missing content",
			input:   &model.NewPost{Title: "Missing content", Content: ""},
			wantErr: true,
		},
		{
			name:    "empty post",
			input:   &model.NewPost{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := s.PostCreate(ctx, *tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error for input %+v, but got nil", tt.input)
				}
				if post != nil {
					t.Errorf("expected nil post for input %+v, but got: %+v", tt.input, post)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for input %+v: %v", tt.input, err)
				}
				if post == nil {
					t.Errorf("expected post for input %+v, got nil", tt.input)
					return
				}
				if post.Title != tt.input.Title {
					t.Errorf("expected title %q, got %q", tt.input.Title, post.Title)
				}
				if post.Content != tt.input.Content {
					t.Errorf("expected content %q, got %q", tt.input.Content, post.Content)
				}
			}
		})
	}
}

func TestPostUpdate(t *testing.T) {
	db := config.GetDB()
	require.NotNil(t, db)

	s := &Service{DB: db}
	ctx := withAuthContext(context.Background(), &middleware.User{ID: 1})

	tests := []struct {
		name    string
		input   *model.UpdatePost
		wantErr bool
	}{
		{
			name: "valid update",
			input: &model.UpdatePost{
				ID:      16,
				Title:   "Updated Post",
				Content: "This is an updated post content",
			},
			wantErr: false,
		},
		{
			name:    "missing ID",
			input:   &model.UpdatePost{Title: "Missing ID", Content: "Missing ID"},
			wantErr: true,
		},
		{
			name:    "missing title",
			input:   &model.UpdatePost{ID: 1, Title: "", Content: "Missing title"},
			wantErr: true,
		},
		{
			name:    "missing content",
			input:   &model.UpdatePost{ID: 1, Title: "Missing Content", Content: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := s.PostUpdate(ctx, *tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error for input %+v, but got nil", tt.input)
				}
				if post != nil {
					t.Errorf("expected nil post for input %+v, but got: %+v", tt.input, post)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for input %+v: %v", tt.input, err)
				}
				if post == nil {
					t.Errorf("expected post for input %+v, got nil", tt.input)
					return
				}
				if post.Title != tt.input.Title {
					t.Errorf("expected title %q, got %q", tt.input.Title, post.Title)
				}
				if post.Content != tt.input.Content {
					t.Errorf("expected content %q, got %q", tt.input.Content, post.Content)
				}
			}
		})
	}
}

func TestPostDelete(t *testing.T) {
	db := config.GetDB()
	require.NotNil(t, db)

	s := &Service{DB: db}
	ctx := withAuthContext(context.Background(), &middleware.User{ID: 1})

	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "Valid delete",
			id:      17,
			wantErr: false,
		},
		{
			name:    "Missing Post ID (Post doesn't exist)",
			id:      9999,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := s.PostDeleteByID(ctx, tt.id)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error for ID %d, but got nil", tt.id)
				}
				if msg != "" {
					t.Errorf("expected empty message for ID %d, but got: %s", tt.id, msg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for ID %d: %v", tt.id, err)
				}
				if msg != "Success" {
					t.Errorf("expected success message for ID %d, got: %s", tt.id, msg)
				}
			}
		})
	}
}
