CREATE TABLE events (
    id          number PRIMARY KEY,
    user_id     number NOT NULL,
    action      text NOT NULL,
    metadata    text NOT NULL,
    time_event  timestamp NOT NULL
);

CREATE TABLE user_event_stats (
    id	        number PRIMARY KEY,
    user_id     number NOT NULL,
    start_time	timestamp NOT NULL,
    end_time	timestamp NOT NULL
    event_count number NOT NULL,
)