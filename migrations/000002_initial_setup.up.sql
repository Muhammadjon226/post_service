CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL Primary Key,
    assignee VARCHAR(128) NOT NULL,
    title VARCHAR(128) NOT NULL,
    summary VARCHAR(256) NOT NULL,
    deadline timestamp NOT NULL,
    status VARCHAR(64) NOT NULL,
    created_at timestamp default current_timestamp,
    updated_at timestamp ,
    deleted_at timestamp   
);
