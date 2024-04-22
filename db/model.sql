-- Users table must be created first since other tables depend on it.
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    hashed_password TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Texts table depends on the Users table.
CREATE TABLE IF NOT EXISTS texts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    text_type VARCHAR(50) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Chapters table depends on the Texts table.
CREATE TABLE IF NOT EXISTS chapters (
    id SERIAL PRIMARY KEY,
    text_id INT REFERENCES texts(id),
    chapter_title VARCHAR(255),
    chapter_number INT NOT NULL
);

-- Chapter Details table depends on the Chapters table.
CREATE TABLE IF NOT EXISTS chapter_details (
    id SERIAL PRIMARY KEY,
    chapter_id INT REFERENCES chapters(id),
    detail_type VARCHAR(100) NOT NULL,
    detail_description TEXT NOT NULL
);

-- Feedback table depends on Chapters and Users table.
CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    chapter_id INT REFERENCES chapters(id),
    user_id INT REFERENCES users(id),
    content TEXT NOT NULL,
    rating INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Trait Types table (independent of others).
CREATE TABLE IF NOT EXISTS trait_types (
    id SERIAL PRIMARY KEY,
    trait_type VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    trigger_text TEXT
);

-- Trait Keys table depends on Trait Types.
CREATE TABLE IF NOT EXISTS trait_keys (
    id SERIAL PRIMARY KEY,
    trait_type_id INT REFERENCES trait_types(id),
    trait_key VARCHAR(100) NOT NULL,
    description TEXT NOT NULL
);

-- Text Traits table depends on Texts and Trait Keys.
CREATE TABLE IF NOT EXISTS text_traits (
    id SERIAL PRIMARY KEY,
    text_id INT REFERENCES texts(id),
    trait_key_id INT REFERENCES trait_keys(id),
    trait_value TEXT NOT NULL
);

-- User Actions table depends on Users.
CREATE TABLE IF NOT EXISTS user_actions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    action_type VARCHAR(50) NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    entity_id INT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE standard_prompts (
  id SERIAL PRIMARY KEY,
  standard_name VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  prompt TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  version INTEGER DEFAULT 1
);

CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX idx_standard_prompts_standard_name ON standard_prompts (standard_name);

CREATE INDEX IF NOT EXISTS texts_content_gin_idx ON texts USING GIN (content gin_trgm_ops);
CREATE INDEX IF NOT EXISTS users_username_gin_idx ON users USING GIN (username gin_trgm_ops);
CREATE INDEX IF NOT EXISTS users_email_gin_idx ON users USING GIN (email gin_trgm_ops);
CREATE INDEX IF NOT EXISTS chapters_title_gin_idx ON chapters USING GIN (chapter_title gin_trgm_ops);
CREATE INDEX IF NOT EXISTS feedback_content_gin_idx ON feedback USING GIN (content gin_trgm_ops);

INSERT INTO trait_types (trait_type, description, trigger_text) VALUES
    ('Narrative Devices', 'Techniques used to structure and present the story', 'Incorporate the following narrative devices into your chapter: <<TEXT>>'),
    ('Reader Engagement', 'Elements that draw the reader into the story', 'Engage the reader by including: <<TEXT>>'),
    ('Stylistic Elements', 'Techniques used to enhance the writing style', 'Enhance your writing style with the following elements: <<TEXT>>'),
    ('Inter-textual References', 'References to external works, events, or culture', 'Include the following inter-textual references in your chapter: <<TEXT>>'),
    ('Authorial Intrusion', 'Instances where the author directly addresses the reader', 'Consider using the following authorial intrusion techniques: <<TEXT>>'),
    ('Critique and Reviews', 'Notable points from critiques, reviews, or fan reactions', 'Address the following critique points or fan reactions in your chapter: <<TEXT>>'),
    ('Character Development', 'Elements that contribute to character growth and depth', 'Develop your characters by exploring: <<TEXT>>'),
    ('Emotional Resonance', 'Techniques used to evoke emotions in the reader', 'Create emotional resonance with the following elements: <<TEXT>>'),
    ('Cliffhangers and Suspense', 'Elements that create tension and anticipation', 'Build suspense and anticipation with: <<TEXT>>')

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'analyze_paragraph',
  'Analyze Paragraph',
  'Analyze the following text and give me the result on this json format, do not mention anything especifique of the story like character names etc... make it generic: {
    "tone": "Analyze the overall tone and mood of the given text. Consider the emotional undertones, atmosphere, and the feelings evoked. Describe the tone using adjectives and elaborate on how it is conveyed through the choice of words, imagery, and literary devices employed.",

    "wording": "Examine the wording and language used in the text. Identify any notable or recurring stylistic elements, such as descriptive language, figurative expressions (metaphors, similes), or unique word choices. Comment on how the wording contributes to the overall tone and meaning of the text.",

    "style": "Provide an analysis of the writing style employed in the text. Identify the key characteristics that define the author''s style, such as introspection, descriptive writing, focus on internal states, or any other notable elements. Discuss how the style enhances the themes, emotions, or ideas conveyed in the text."
  } 
  
  My text is: 

  <<Text>>',
  1
);
INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'initial_brainstorming_1',
  'Initial Brainstorming 1',
  'Supply <<NUMBER>> high-concept ideas for a bestselling <<GENRE>> story, featuring a unique twist, captivating characters, and intense emotional stakes.',
  1
);
INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'initial_brainstorming_2',
  'Initial Brainstorming 2',
  'Provide <<NUMBER>> detailed character concepts for a bestselling <<GENRE>> narrative, each with distinct strengths and weaknesses, and a journey marked by conflict. Include a brief outline of their character development.',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'outlining_prompts_1',
  'Outlining Prompt 1',
  'Develop a narrative synopsis for a <<GENRE>> book using the core idea provided: <<PITCH>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'setting_prompts_1',
  'Setting Prompt 1',
  'Identify a range of possible settings for a novel outlined in <<SYNOPSIS>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'character_prompts_1',
  'Character Prompt 1',
  'Assemble a roster of potential characters for a story described by
    <<SUMMARY>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'writing_prompts_1',
  'Writing Prompt 1',
  'Compose <<NUMBER>> words from a chapter section, incorporating these specifics:
    Literary Genre: <<GENRE>>
    Emotional Tone: <<TONE>>
    Narrative Perspective: <<POV>>
    Backdrop: <<SETTING>>
    Principal Figures: <<CHARACTERS>>
    Synopsis: <<SUMMARY>>
    Central Tension: <<CONFLICT>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'editing_prompts_1',
  'Editing Prompt 1',
  'Craft an enhanced hook and introductory paragraph reflecting the hallmark style of a top-selling <<GENRE>> author, based on the scene provided: <<SCENE>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'title_description_prompts_1',
  'Title/Description Prompt 1',
  'Compile a selection of possible titles for a novel that revolves around the theme: <<THEME>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'fine_tuning_writting_1',
  'Fine Tuning Writting 1',
  'System: As the <<GENRE>> author <<AUTHOR>>, upon receiving a scene beat, compose <<WORDS>> words inspired by the prompt. Begin amidst the unfolding action. Adhere strictly to the given beat instructions. Refrain from resolving the scene or incorporating foreshadowing. User: <<STORYBEAT>> Assistant: <<PROSE>>',
  1
);

INSERT INTO standard_prompts (standard_name, title, prompt, version)
VALUES (
  'fine_tuning_copy_editing_1',
  'Fine Tuning Copy Editing 1',
  'System: You are a seasoned copy editor. When presented with prose, deliver a copy-edited version that maintains the distinctive style of <<AUTHOR>>. User: <<PROSE>> Assistant: <<EDITEDPROSE>>'
',
  1
);