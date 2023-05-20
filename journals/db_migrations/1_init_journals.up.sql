CREATE TABLE IF NOT EXISTS journals
(
  id        INTEGER PRIMARY KEY  NOT NULL,
  title     TEXT                 NOT NULL,
  date      DATE                 NOT NULL,
  content   TEXT
)
