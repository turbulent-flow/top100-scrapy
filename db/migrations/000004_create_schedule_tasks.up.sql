CREATE TABLE schedule_tasks (
    id bigserial PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL,
    recorded_key varchar(55) NOT NULL,
    recorded_value varchar(55) NOT NULL,
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp
);

CREATE UNIQUE INDEX schedule_tasks_name_recorded_key_idx ON schedule_tasks (name, recorded_key);
