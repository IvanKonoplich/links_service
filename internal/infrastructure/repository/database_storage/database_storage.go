package database_storage

import (
	"context"
	"fmt"
	"github.com/IvanKonoplich/shortened_links_service/internal/infrastructure/repository"
	"github.com/jmoiron/sqlx"
)

type DatabaseStorage struct {
	db *sqlx.DB
}

func NewDatabaseStorage(db *sqlx.DB) *DatabaseStorage {
	return &DatabaseStorage{db: db}
}

func (d *DatabaseStorage) SaveLink(ctx context.Context, originalLink string) (shortenedLink string, err error) {
	query := "INSERT INTO links(original_link, short_link) VALUES ($1, $2)"

	var shortLink string
	for {
		shortLink = repository.GenerateRandomLink()

		exist, err := d.CheckShortLink(ctx, shortLink)
		if err != nil {
			return "", fmt.Errorf("ошибка во время запроса к бд %s", err.Error())
		}

		if !exist {
			break
		}
	}

	_, err = d.db.ExecContext(ctx, query, originalLink, shortLink)
	if err != nil {
		return "", fmt.Errorf("ошибка во время запроса к бд %s", err.Error())
	}

	return shortLink, nil
}

func (d *DatabaseStorage) GetLink(ctx context.Context, shortenedLink string) (originalLink string, err error) {
	exist, err := d.CheckShortLink(ctx, shortenedLink)
	if err != nil {
		return "", fmt.Errorf("ошибка во время запроса к бд %s", err.Error())
	}

	if !exist {
		return "", fmt.Errorf("нет сохраненных данных для ссылки %s", shortenedLink)
	}

	query := "SELECT original_link FROM links WHERE short_link=$1"

	row := d.db.QueryRowContext(ctx, query, shortenedLink)
	if err := row.Scan(&originalLink); err != nil {
		return "", fmt.Errorf("ошибка во время запроса к бд %s", err.Error())
	}

	return originalLink, nil
}

func (d *DatabaseStorage) CheckShortLink(ctx context.Context, shortenedLink string) (bool, error) {
	var result bool
	query := "SELECT exists(SELECT original_link FROM links WHERE short_link=$1)"
	row := d.db.QueryRowContext(ctx, query, shortenedLink)
	if err := row.Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}
