package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"techtrain-go-practice/db"
	"techtrain-go-practice/handler/router"
	"techtrain-go-practice/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupEnv() error {
	err := os.Setenv("BASIC_AUTH_USER_ID", "user")
	if err != nil {
		return err
	}
	err = os.Setenv("BASIC_AUTH_PASSWORD", "password")
	if err != nil {
		return err
	}
	return nil
}

func TestBasicAuth(t *testing.T) {
	setupEnv()

	t.Parallel()

	todos := []*model.TODO{
		{
			ID:      3,
			Subject: "todo subject 3",
		},
		{
			ID:      2,
			Subject: "todo subject 2",
		},
		{
			ID:      1,
			Subject: "todo subject 1",
		},
	}

	dbPath := "./temp_test.db"
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		t.Errorf("データベースの作成に失敗しました: %v", err)
		return
	}

	t.Cleanup(func() {
		if err := todoDB.Close(); err != nil {
			t.Errorf("データベースのクローズに失敗しました: %v", err)
			return
		}
		if err := os.Remove(dbPath); err != nil {
			t.Errorf("テスト用のDBファイルの削除に失敗しました: %v", err)
			return
		}
	})

	stmt, err := todoDB.Prepare(`INSERT INTO todos(subject, description) VALUES(?, ?)`)
	if err != nil {
		t.Errorf("データベースのステートメントの作成に失敗しました: %v", err)
		return
	}

	t.Cleanup(func() {
		if err := stmt.Close(); err != nil {
			t.Errorf("データベースのステートメントのクローズに失敗しました: %v", err)
			return
		}
	})

	for _, todo := range []*model.TODO{todos[2], todos[1], todos[0]} {
		if _, err = stmt.Exec(todo.Subject, todo.Description); err != nil {
			t.Errorf("データベースのステートメントの実行に失敗しました: %v", err)
			return
		}
	}

	testcases := map[string]struct {
		UserID         string
		Password       string
		Path           string
		WantStatusCode int
	}{
		"Do not need Basic Auth": {
			UserID:         "user",
			Password:       "password",
			Path:           "/healthz",
			WantStatusCode: http.StatusOK,
		},
		"Correct UserID and Password": {
			UserID:         "user",
			Password:       "password",
			Path:           "/todos",
			WantStatusCode: http.StatusOK,
		},
		"Incorrect UserID": {
			UserID:         "wrong_user",
			Password:       "password",
			Path:           "/todos",
			WantStatusCode: http.StatusUnauthorized,
		},
		"Incorrect Password": {
			UserID:         "user",
			Password:       "wrong_password",
			Path:           "/todos",
			WantStatusCode: http.StatusUnauthorized,
		},
		"Incorrect UserID and Password": {
			UserID:         "wrong_user",
			Password:       "wrong_password",
			Path:           "/todos",
			WantStatusCode: http.StatusUnauthorized,
		},
		"Empty UserID and Password": {
			UserID:         "",
			Password:       "",
			Path:           "/todos",
			WantStatusCode: http.StatusUnauthorized,
		},
		"Empty UserID": {
			UserID:         "",
			Password:       "password",
			Path:           "/todos",
			WantStatusCode: http.StatusUnauthorized,
		},
		"Empty Password": {
			UserID:         "user",
			Password:       "",
			Path:           "/todos",
			WantStatusCode: http.StatusUnauthorized,
		},
	}

	client := &http.Client{}
	router := router.NewRouter(todoDB)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", testServer.URL+tc.Path, nil)

			// Set Basic Auth for `/todos``
			if tc.Path == "/todos" {
				req.SetBasicAuth(tc.UserID, tc.Password)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Request failed: %v", err)
				return
			}

			assert.Equal(t, tc.WantStatusCode, resp.StatusCode)
		})
	}
}
