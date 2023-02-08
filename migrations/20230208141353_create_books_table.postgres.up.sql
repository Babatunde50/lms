CREATE TABLE IF NOT EXISTS books (
    isbn VARCHAR(15) PRIMARY KEY NOT NULL,
    title VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    publication_date TIMESTAMP NOT NULL,
    publisher VARCHAR(50) NOT NULL,
    genre VARCHAR(50) NOT NULL,
    language VARCHAR(50) NOT NULL,
    description VARCHAR(250) NOT NULL,
    cover_image VARCHAR(200),
    available_copies INTEGER NOT NULL,
    total_copies INTEGER NOT NULL,
    format VARCHAR(25) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)