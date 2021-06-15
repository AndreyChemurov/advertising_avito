CREATE TABLE IF NOT EXISTS advertisement (
    id BIGSERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(200) NOT NULL CHECK (name <> ''),
	description VARCHAR(1000) NOT NULL CHECK (description <> ''),
	price NUMERIC(16, 2) CHECK (price > 0.0) NOT NULL,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS photos (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	adv_id INT NOT NULL,
	link TEXT NOT NULL CHECK (link <> ''),
	FOREIGN KEY (adv_id) REFERENCES advertisement(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS adv_id_idx ON advertisement (id);
CREATE INDEX IF NOT EXISTS adv_price_idx ON advertisement (price);
CREATE INDEX IF NOT EXISTS adv_date_idx ON advertisement (created_at);
CREATE INDEX IF NOT EXISTS photos_adv_id_idx ON photos (adv_id);