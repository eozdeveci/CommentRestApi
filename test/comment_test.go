package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
func TestGetComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment/1")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/comment")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUpdateComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Put(BASE_URL + "/api/comment/1")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		Delete(BASE_URL + "/api/comment/2")

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
