CREATE UNIQUE INDEX IF NOT EXISTS unique_shared ON shared_entries (entry_id, shared_by, shared_with);
