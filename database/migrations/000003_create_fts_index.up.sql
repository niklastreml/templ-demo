CREATE INDEX idx_project_name ON project USING GIN (to_tsvector('english', name));
