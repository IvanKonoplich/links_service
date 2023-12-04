package in_memory

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// тест на простую работоспособность
func TestSaveLink1(t *testing.T) {
	m := NewInMemoryStorage()

	result, err := m.SaveLink(context.Background(), "https://job.ozon.ru/internships/")

	assert.Equal(t, 10, len(result))
	assert.Nil(t, err)
}

// тест на два одинаковых значения
func TestSaveLink2(t *testing.T) {
	m := NewInMemoryStorage()

	result1, err1 := m.SaveLink(context.Background(), "https://job.ozon.ru/internships/")

	assert.Equal(t, 10, len(result1))
	assert.Nil(t, err1)

	result2, err2 := m.SaveLink(context.Background(), "https://job.ozon.ru/internships/")

	assert.Equal(t, 10, len(result2))
	assert.Nil(t, err2)

	assert.NotEqualValues(t, result1, result2)
}

// тест на значение меньше 10 символов
func TestSaveLink3(t *testing.T) {
	m := NewInMemoryStorage()

	result, err := m.SaveLink(context.Background(), "http://")

	assert.Equal(t, 10, len(result))
	assert.Nil(t, err)
}

// тест на пустое значение
func TestSaveLink4(t *testing.T) {
	m := NewInMemoryStorage()

	_, err := m.SaveLink(context.Background(), "")

	assert.NotNil(t, err)
}

// тест на простую работоспособность
func TestGetLink1(t *testing.T) {
	m := NewInMemoryStorage()

	inputLink := "https://job.ozon.ru/internships/"
	shortLink, _ := m.SaveLink(context.Background(), inputLink)

	result, err := m.GetLink(context.Background(), shortLink)

	assert.Nil(t, err)
	assert.Equal(t, inputLink, result)
}

// тест на два одинаковых значения
func TestGetLink2(t *testing.T) {
	m := NewInMemoryStorage()

	inputLink := "https://job.ozon.ru/internships/"
	shortLink1, _ := m.SaveLink(context.Background(), inputLink)
	result1, err := m.GetLink(context.Background(), shortLink1)
	assert.Nil(t, err)

	shortLink2, _ := m.SaveLink(context.Background(), inputLink)
	result2, err := m.GetLink(context.Background(), shortLink2)
	assert.Nil(t, err)

	assert.NotEqualValues(t, shortLink1, shortLink2)
	assert.EqualValues(t, result1, result2)
}

// тест на пустое значение
func TestGetLink3(t *testing.T) {
	m := NewInMemoryStorage()

	shortLink := ""
	_, err := m.GetLink(context.Background(), shortLink)

	assert.NotNil(t, err)
}

// тест на несуществующее значение
func TestGetLink4(t *testing.T) {
	m := NewInMemoryStorage()

	shortLink := "https://job.ozon.ru/internships/"
	_, err := m.GetLink(context.Background(), shortLink)

	assert.NotNil(t, err)
}
