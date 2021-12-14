CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL Primary Key,
    assignee VARCHAR(128),
    title VARCHAR(128),
    summary VARCHAR(256),
    deadline timestamp default null,
    status VARCHAR(64)
);
