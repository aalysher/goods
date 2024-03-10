CREATE TABLE goods (
                       id SERIAL PRIMARY KEY,
                       project_id INT NOT NULL,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       priority INT,
                       removed BOOLEAN NOT NULL DEFAULT FALSE,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE INDEX idx_goods_id ON goods(id);
CREATE INDEX idx_goods_project_id ON goods(project_id);
CREATE INDEX idx_goods_name ON goods(name);
