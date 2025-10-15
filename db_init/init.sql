CREATE TABLE events (
    id          SERIAL PRIMARY KEY,
    user_id     int NOT NULL,
    action      text NOT NULL,
    metadata    text NOT NULL,
    time_event  timestamp NOT NULL
);

CREATE TABLE user_event_stats (
    id	        SERIAL PRIMARY KEY,
    user_id     int NOT NULL,
    start_time	timestamp NOT NULL,
    end_time	timestamp NOT NULL,
    event_count int NOT NULL
)