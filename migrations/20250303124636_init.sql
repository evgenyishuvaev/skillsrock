-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.tasks (
	id serial4 NOT NULL,
	title varchar NULL,
	description text NULL,
	status text DEFAULT 'new'::text NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT tasks_pkey PRIMARY KEY (id),
	CONSTRAINT tasks_status_check CHECK ((status = ANY (ARRAY['new'::text, 'in_progress'::text, 'done'::text])))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
