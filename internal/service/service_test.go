package service

import (
	"context"
	"github.com/IvanKonoplich/shortened_links_service/internal/service/mocks"
	"github.com/IvanKonoplich/shortened_links_service/pkg/shortened_links"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveLink1(t *testing.T) {
	mockStorage := mocks.NewStorage(t)
	mockStorage.Mock.On("SaveLink", context.Background(), "https://job.ozon.ru/internships/").Return("1234567890", nil)

	testService := New(mockStorage)
	result, err := testService.SaveLink(context.Background(), &shortened_links.OriginalLink{Message: "https://job.ozon.ru/internships/"})

	assert.Equal(t, 10, len(result.GetMessage()))
	assert.Nil(t, err)
}

func TestGetLink1(t *testing.T) {
	mockStorage := mocks.NewStorage(t)
	mockStorage.Mock.On("GetLink", context.Background(), "1234567890").Return("https://job.ozon.ru/internships/", nil)

	testService := New(mockStorage)
	result, err := testService.GetLink(context.Background(), &shortened_links.ShortenedLink{Message: "1234567890"})

	assert.Equal(t, "https://job.ozon.ru/internships/", result.GetMessage())
	assert.Nil(t, err)
}
