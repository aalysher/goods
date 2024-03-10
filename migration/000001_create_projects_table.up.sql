CREATE TABLE projects (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_projects_id ON projects(id);


