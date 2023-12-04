CREATE TABLE IF NOT EXISTS links(
    original_link  varchar not null,
    short_link  varchar not null unique);

CREATE INDEX shorturl_link_hash_index ON links USING hash(short_link);
