CREATE TABLE IF NOT EXISTS chat(
    chatname VARCHAR(100) NOT NULL,
    creator VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS messages(
    message_id SERIAL PRIMARY KEY,
    chatname VARCHAR(100) NOT NULL,
    creator VARCHAR(100) NOT NULL,
    mess_text TEXT NOT NULL
);