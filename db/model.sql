CREATE TABLE IF NOT EXISTS trait_types (
    trait_type_id SERIAL PRIMARY KEY,
    trait_type VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    trigger_text TEXT
);

CREATE TABLE IF NOT EXISTS trait_keys (
    trait_key_id SERIAL PRIMARY KEY,
    trait_type_id INT REFERENCES trait_types(trait_type_id),
    trait_key VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS text_traits (
    trait_id SERIAL PRIMARY KEY,
    text_id INT REFERENCES texts(text_id),
    trait_key_id INT REFERENCES trait_keys(trait_key_id),
    trait_value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    hashed_password TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS texts (
    text_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    text_type VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS chapters (
    chapter_id SERIAL PRIMARY KEY,
    text_id INT REFERENCES texts(text_id),
    chapter_title VARCHAR(255),
    chapter_number INT NOT NULL
);

CREATE TABLE IF NOT EXISTS chapter_details (
    detail_id SERIAL PRIMARY KEY,
    chapter_id INT REFERENCES chapters(chapter_id),
    detail_type VARCHAR(100) NOT NULL,
    detail_description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS feedback (
    feedback_id SERIAL PRIMARY KEY,
    chapter_id INT REFERENCES chapters(chapter_id),
    user_id INT REFERENCES users(user_id) NULL,
    content TEXT NOT NULL,
    rating INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_actions (
    action_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    action_type VARCHAR(50) NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    entity_id INT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Add this line to enable the full-text search extension
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Add a GIN index for full-text search on the 'content' column of the 'texts' table
CREATE INDEX IF NOT EXISTS texts_content_gin_idx ON texts USING GIN (content gin_trgm_ops);

-- Add a GIN index for full-text search on the 'username' and 'email' columns of the 'users' table
CREATE INDEX IF NOT EXISTS users_username_gin_idx ON users USING GIN (username gin_trgm_ops);
CREATE INDEX IF NOT EXISTS users_email_gin_idx ON users USING GIN (email gin_trgm_ops);

-- Add a GIN index for full-text search on the 'chapter_title' column of the 'chapters' table
CREATE INDEX IF NOT EXISTS chapters_title_gin_idx ON chapters USING GIN (chapter_title gin_trgm_ops);

-- Add a GIN index for full-text search on the 'content' column of the 'feedback' table
CREATE INDEX IF NOT EXISTS feedback_content_gin_idx ON feedback USING GIN (content gin_trgm_ops);

INSERT INTO trait_types (trait_type, description, trigger_text)
VALUES
    ('Narrative Devices', 'Techniques used to structure and present the story', 'Incorporate the following narrative devices into your chapter:'),
    ('Reader Engagement', 'Elements that draw the reader into the story', 'Engage the reader by including:'),
    ('Stylistic Elements', 'Techniques used to enhance the writing style', 'Enhance your writing style with the following elements:'),
    ('Inter-textual References', 'References to external works, events, or culture', 'Include the following inter-textual references in your chapter:'),
    ('Authorial Intrusion', 'Instances where the author directly addresses the reader', 'Consider using the following authorial intrusion techniques:'),
    ('Critique and Reviews', 'Notable points from critiques, reviews, or fan reactions', 'Address the following critique points or fan reactions in your chapter:'),
    ('Character Development', 'Elements that contribute to character growth and depth', 'Develop your characters by exploring:'),
    ('Emotional Resonance', 'Techniques used to evoke emotions in the reader', 'Create emotional resonance with the following elements:'),
    ('Cliffhangers and Suspense', 'Elements that create tension and anticipation', 'Build suspense and anticipation with:'),
    ('Philosophical and Intellectual Depth', 'Incorporation of thought-provoking ideas and themes', 'Enhance the philosophical and intellectual depth of your chapter by including:'),
    ('Narrative Experimentation', 'Unconventional storytelling techniques and structures', 'Experiment with your narrative by using:'),
    ('Visual Elements', 'Notable visual aspects of the text or book design', 'Consider the following visual elements in your chapter:');