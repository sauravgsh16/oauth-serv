CREATE TABLE clients (
	id text PRIMARY KEY,
	created_at BIGINT,
	deleted_at BIGINT,
	updated_at BIGINT,
	client_id varchar(250) UNIQUE NOT NULL,
	secret varchar(250) NOT NULL,
	redirect_uri varchar(250)
)