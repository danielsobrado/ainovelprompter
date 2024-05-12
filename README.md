# AI Novel Prompter

AI Novel Prompter can generate writing prompts for novels based on user-specified characteristics. 

## Features

- User registration and authentication
- Text creation and management
- Chapter creation and management
- Feedback submission and management
- Prompt generation based on traits
- Integration with a local ollama service
- Based on Berry template (https://codedthemes.gitbook.io/berry)
- Inspired on Jason Hamilton Youtube (https://www.youtube.com/@TheNerdyNovelist)

## Technologies Used

- Frontend:
  - React
  - TypeScript
  - Axios
  - React Router
  - React Toastify
- Backend:
  - Go
  - Gin Web Framework
  - GORM (Go ORM)
  - PostgreSQL

## Prerequisites

Before running the application, make sure you have the following installed:

- Node.js (v18 or higher)
- Go (v1.18 or higher)
- PostgreSQL
- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/danielsobrado/ainovelprompter.git
   ```
2. Navigate to the project directory:
   ```
   cd ainovelprompter
   ```
3. Set up the backend:

- Navigate to the `server` directory:

  ```
  cd server
  ```

- Install the Go dependencies:

  ```
  go mod download
  ```

- Update the `config.yaml` file with your database configuration.

- Run the database migrations:

  ```
  go run cmd/main.go migrate
  ```

- Start the backend server:

  ```
  go run cmd/main.go
  ```

4. Set up the frontend:

- Navigate to the `client` directory:

  ```
  cd ../client
  ```

- Install the frontend dependencies:

  ```
  npm install
  ```

- Start the frontend development server:
  ```
  npm start
  ```
5. Open your web browser and visit `http://localhost:3000` to access the application.

## Getting Started (Docker)

1. Clone the repository:
```
git clone https://github.com/danielsobrado/ainovelprompter.git
```

2. Navigate to the project directory:
```
cd ainovelprompter
```

3. Update the `docker-compose.yml` file with your database configuration.

4. Start the application using Docker Compose:
```
docker-compose up -d
```

5. Open your web browser and visit `http://localhost:3000` to access the application.

## Configuration

- Backend configuration can be modified in the `server/config.yaml` file.
- Frontend configuration can be modified in the `client/src/config.ts` file.

## Build

To build the frontend for production, run the following command in the `client` directory:
   ```
   npm run build
   ```
The production-ready files will be generated in the `client/build` directory.

## Installation and Management Guide for PostgreSQL on WSL

This small guide provides instructions on how to install PostgreSQL on the Windows Subsystem for Linux (WSL), along with steps to manage user permissions and troubleshoot common issues.

---

### Prerequisites

- Windows 10 or higher with WSL enabled. (Or just Ubuntu)
- Basic familiarity with Linux command line and SQL.

---

### Installation

1. **Open WSL Terminal**: Launch your WSL distribution (Ubuntu recommended).

2. **Update Packages**:
   ```bash
   sudo apt update
   ```

3. **Install PostgreSQL**:
   ```bash
   sudo apt install postgresql postgresql-contrib
   ```

4. **Check Installation**:
   ```bash
   psql --version
   ```

5. **Set PostgreSQL User Password**:
   ```bash
   sudo passwd postgres
   ```

---

### Database Operations

1. **Create Database**:
   ```bash
   createdb mydb
   ```

2. **Access Database**:
   ```bash
   psql mydb
   ```

3. **Import Tables from SQL File**:
   ```bash
   psql -U postgres -q mydb < /path/to/file.sql
   ```

4. **List Databases and Tables**:
   ```sql
   \l  # List databases
   \dt # List tables in the current database
   ```

5. **Switch Database**:
   ```sql
   \c dbname
   ```

---

### User Management

1. **Create New User**:
   ```sql
   CREATE USER your_db_user WITH PASSWORD 'your_db_password';
   ```

2. **Grant Privileges**:
   ```sql
   ALTER USER your_db_user CREATEDB;
   ```

---

### Troubleshooting

1. **Role Does Not Exist Error**:
   Switch to the 'postgres' user:
   ```bash
   sudo -i -u postgres
   createdb your_db_name
   ```

2. **Permission Denied to Create Extension**:
   Login as 'postgres' and execute:
   ```sql
   CREATE EXTENSION IF NOT EXISTS pg_trgm;
   ```

3. **Unknown User Error**:
   Ensure you are using a recognized system user or correctly refer to a PostgreSQL user within the SQL environment, not via `sudo`.

---

## Generating Custom Training Data to Fine-Tune a Language Model (Manual Steps)

To [generate custom training data](https://github.com/danielsobrado/ainovelprompter/tree/main/finetune_data_example1) for fine-tuning a language model to emulate the writing style of George MacDonald, the process begins by obtaining the full text of one of his novels, "The Princess and the Goblin," from Project Gutenberg. The text is then broken down into individual story beats or key moments using a prompt that instructs the AI to generate a JSON object for each beat, capturing the author, emotional tone, type of writing, and the actual text excerpt.

Next, GPT-4 is used to rewrite each of these story beats in its own words, generating a parallel set of JSON data with unique identifiers linking each rewritten beat to its original counterpart. To simplify the data and make it more useful for training, the wide variety of emotional tones is mapped to a smaller set of core tones using a Python function. The two JSON files (original and rewritten beats) are then used to generate training prompts, where the model is asked to rephrase the GPT-4 generated text in the style of the original author. Finally, these prompts and their target outputs are formatted into JSONL and JSON files, ready to be used for fine-tuning the language model to capture MacDonald's distinctive writing style.

---

## Fine-Tuning lessons

* Dataset Quality Matters: 95% of outcomes depend on dataset quality. A clean dataset is essential since even a little bad data can hurt the model.

* Manual Data Review: Cleaning and evaluating the dataset can greatly improve the model. This is a time-consuming but necessary step because no amount of parameter adjusting can fix a defective dataset.

* Training parameters should not improve but prevent model degradation. In robust datasets, the goal should be to avoid negative repercussions while directing the model. There is no optimal learning rate.

* Model Scale and Hardware Limitations: Larger models (33b parameters) may enable better fine-tuning but require at least 48GB VRAM, making them impractical for majority of home setups.

* Gradient Accumulation and Batch Size: Gradient accumulation helps reduce overfitting by enhancing generalisation across different datasets, but it may lower quality after a few batches.

* The size of the dataset is more important for fine-tuning a base model than a well-tuned model. Overloading a well-tuned model with excessive data might degrade its previous fine-tuning.

* An ideal learning rate schedule starts with a warmup phase, holds steady for an epoch, and then gradually decreases using a cosine schedule.

* Model Rank and Generalisation: The amount of trainable parameters affects the model's detail and generalisation. Lower-rank models generalise better but lose detail.

* LoRA's Applicability: Parameter-Efficient Fine-Tuning (PEFT) is applicable to large language models (LLMs) and systems like Stable Diffusion (SD), demonstrating its versatility.

---

## Evaluation Metrics

When fine-tuning a language model for paraphrasing in an author's style, it's important to evaluate the quality and effectiveness of the generated paraphrases. 

The following evaluation metrics can be used to assess the model's performance:

1. **BLEU (Bilingual Evaluation Understudy):**
   - BLEU measures the n-gram overlap between the generated paraphrase and the reference text, providing a score between 0 and 1.
   - To calculate BLEU scores, you can use the `sacrebleu` library in Python.
   - Example usage: `from sacrebleu import corpus_bleu; bleu_score = corpus_bleu(generated_paraphrases, [original_paragraphs])`

2. **ROUGE (Recall-Oriented Understudy for Gisting Evaluation):**
   - ROUGE measures the overlap of n-grams between the generated paraphrase and the reference text, focusing on recall.
   - To calculate ROUGE scores, you can use the `rouge` library in Python.
   - Example usage: `from rouge import Rouge; rouge = Rouge(); scores = rouge.get_scores(generated_paraphrases, original_paragraphs)`

3. **Perplexity:**
   - Perplexity quantifies the uncertainty or confusion of the model when generating text.
   - To calculate perplexity, you can use the fine-tuned language model itself.
   - Example usage: `perplexity = model.perplexity(generated_paraphrases)`

4. **Stylometric Measures:**
   - Stylometric measures capture the writing style characteristics of the target author.
   - To extract stylometric features, you can use the `stylometry` library in Python.
   - Example usage: `from stylometry import extract_features; features = extract_features(generated_paraphrases)`

### Integration with Axolotl

To integrate these evaluation metrics into your Axolotl pipeline, follow these steps:

1. Prepare your training data by creating a dataset of paragraphs from the target author's works and splitting it into training and validation sets.

2. Fine-tune your language model using the training set, following the approach discussed earlier.

3. Generate paraphrases for the paragraphs in the validation set using the fine-tuned model.

4. Implement the evaluation metrics using the respective libraries (`sacrebleu`, `rouge`, `stylometry`) and calculate the scores for each generated paraphrase.

5. Perform human evaluation by collecting ratings and feedback from human evaluators.

6. Analyze the evaluation results to assess the quality and style of the generated paraphrases and make informed decisions to improve your fine-tuning process.

Here's an example of how you can integrate these metrics into your pipeline:

```python
from sacrebleu import corpus_bleu
from rouge import Rouge
from stylometry import extract_features

# Fine-tune the model using the training set
fine_tuned_model = train_model(training_data)

# Generate paraphrases for the validation set
generated_paraphrases = generate_paraphrases(fine_tuned_model, validation_data)

# Calculate evaluation metrics
bleu_score = corpus_bleu(generated_paraphrases, [original_paragraphs])
rouge = Rouge()
rouge_scores = rouge.get_scores(generated_paraphrases, original_paragraphs)
perplexity = fine_tuned_model.perplexity(generated_paraphrases)
stylometric_features = extract_features(generated_paraphrases)

# Perform human evaluation
human_scores = collect_human_evaluations(generated_paraphrases)

# Analyze and interpret the results
analyze_results(bleu_score, rouge_scores, perplexity, stylometric_features, human_scores)
```

Remember to install the necessary libraries (sacrebleu, rouge, stylometry) and adapt the code to fit your implementation in Axolotl or similar.

---

## AI Writing Model Comparison

In this [experiment](https://github.com/danielsobrado/ainovelprompter/tree/main/compare), I explored the capabilities and differences between various AI models in generating a 1500-word text based on a detailed prompt. I tested models from https://chat.lmsys.org/, ChatGPT4, Claude 3 Opus, and some local models in LM Studio. Each model generated the text three times to observe variability in their outputs. I also created a separate prompt for evaluating the writing of the first iteration from each model and asked ChatGPT 4 and Claude Opus 3 to provide feedback.

Through this process, I observed that some models exhibit higher variability between executions, while others tend to use similar wording. There were also significant differences in the number of words generated and the amount of dialogue, descriptions, and paragraphs produced by each model. The evaluation feedback revealed that ChatGPT suggests a more "refined" prose, while Claude recommends less purple prose. Based on these findings, I compiled a list of takeaways to incorporate into the next prompt, focusing on precision, varied sentence structures, strong verbs, unique twists on fantasy motifs, consistent tone, distinct narrator voice, and engaging pacing. Another technique to consider is asking for feedback and then rewriting the text based on that feedback. 

I'm open to collaborating with others to further fine-tune prompts for each model and explore their capabilities in creative writing tasks.

## Contributing

All comments are welcome. Open an issue or send a pull request if you find any bugs or have recommendations for improvement.

## License

This project is licensed under: Attribution-NonCommercial-NoDerivatives (BY-NC-ND) license See: https://creativecommons.org/licenses/by-nc-nd/4.0/deed.en
