# Generating Custom Training Data to Fine-Tune a Language Model

This is the process of generating custom training data to fine-tune a large language model to generate text in the distinctive style of 19th century author George MacDonald. The source material is MacDonald's novel "The Princess and the Goblin", which is available copyright-free through Project Gutenberg.

## Process Overview

1. Obtain the source text
2. Break the text into story beats
3. Generate GPT-4 versions of the story beats
4. Map tones to a simplified set
5. Generate the training prompts
6. Format the data for the model

## Detailed Steps

### 1. Obtain the source text

Download a plain text version of MacDonald's novel "The Princess and the Goblin" from the Gutenberg library.

### 2. Break the text into story beats

Use a prompt to have the AI assistant break the full text of the novel down into an array of individual story beats or key moments. Each entry should capture a key story point, which may span multiple paragraphs.

The prompt should instruct the AI to generate a JSON object for each beat, with the following fields:
- `author` (always "George MacDonald")
- `tone` (a one-word description of the emotional tone, e.g. "mysterious", "tense", "joyful")
- `type` (the type of writing, e.g. "dialog", "description")
- `text` (the actual excerpt from the source material)

The output of this step is saved in a file called `analyzed_George_MacDonald.json`.

### 3. Generate GPT-4 versions of the story beats

Take the JSON array of story beats and use it to generate a new prompt. For each original story beat, ask GPT-4 to "rewrite the following text in your own words, that is clearly recognized as GPT-4 generated text by the descriptive wording and tone."

The output of this step is a JSON array of the same length, with each entry containing:
- `id` (a numeric identifier linking it to the original beat)
- `gptText` (GPT-4's rewritten version of that story beat)

This output is saved in `gptStyle_George_MacDonald.json`.

### 4. Map tones to a simplified set

The original prompts generate a wide variety of tones to describe the emotional tenor of each beat. To make the data more useful for training, map this large set of tones to a more limited set of 10 core tones (e.g. Dramatic, Thoughtful, Mysterious, etc.)

Use the following Python function to perform this mapping:

```python
def rewrite_json_with_grouped_tones(data, output_file_path):
    tone_mapping = {
        "Serene": "Calm", "Charming": "Positive", "Mysterious": "Mysterious", 
        "Intriguing": "Thoughtful", "Reflective": "Thoughtful", "Dreary": "Gloomy",
        ...
    }
    
    for index, entry in enumerate(data, start=1):
        original_tone = entry.get('tone')
        entry['tone'] = tone_mapping.get(original_tone, original_tone) 
        entry['id'] = index  

    with open(output_file_path, 'w') as file:
        json.dump(data, file, indent=4)
```

The output of this step is saved in `modified_data_George_MacDonald.json`.

### 5. Generate the training prompts

With the two JSON files (the original story beat breakdowns and the GPT-4 rewritten versions), generate prompts to train the model. For each beat, the prompt should provide the GPT-4 rewrite and instruct the model to "Rephrase the following text in the style of George MacDonald." The target output is the original text.

Use this Python function to generate the prompts:

```python
def generate_prompts_for_rewriting(input_file_path, output_file_path):
    with open(input_file_path, 'r') as json_file, open(output_file_path, 'w') as text_file:
        data = json.load(json_file)
        for index, entry in enumerate(data, start=1):
            if 'text' in entry:
                text = entry['text']  
                prompt = (f"Rewrite the following text in JSON format with your own words, "
                          f"that is clearly recognized as GPT4 generated text by the descriptive wording and tone "
                          f"and return the result in a JSON format with a field called 'id' that has number {index} "
                          f"and a field called 'gptText' that has the text you rewrote, the text is:\n{text}\n\n")
                text_file.write(prompt)  
            else:
                text_file.write(f"No text found for entry {index}.\n\n") 
```

The output is saved as `instructions_George_MacDonald.txt`.

### 6. Format the data for the model

Take the set of training prompts and format them properly for the specific model and training approach being used. In this case, convert the data into multiple formats (JSONL and JSON) to have flexibility in training.

The JSONL formatting is done with this function:

```python
def generate_train_data(gpt_file, modified_file, output_file):
    with open(gpt_file, 'r') as file:
        gpt_data = json.load(file)

    with open(modified_file, 'r') as file:
        modified_data = json.load(file)

    modified_dict = {entry['id']: entry for entry in modified_data}

    with open(output_file, 'w') as output:
        for gpt_entry in gpt_data:
            gpt_id = gpt_entry['id']
            gpt_text = gpt_entry['gptText']

            if gpt_id in modified_dict:
                modified_entry = modified_dict[gpt_id]
                author = modified_entry['author']
                tone = modified_entry['tone']
                text_type = modified_entry['type']
                modified_text = modified_entry['text']

                gpt_text = escape_special_chars(gpt_text)
                modified_text = escape_special_chars(modified_text)

                formatted_text = f"<human>:Rephrase the following text with in the style of {author}: {gpt_text}\\n<bot>: {modified_text}"

                metadata = {"source": "gutenberg"}

                jsonl_entry = {"text": formatted_text, "metadata": metadata}

                output.write(json.dumps(jsonl_entry) + '\n')
```

The output is saved as `train_data_George_MacDonald.jsonl`.

And finally, the conversion from JSONL to JSON is done with:

```python
def convert_jsonl_to_json(input_file, output_file):
    with open(input_file, 'r') as file:
        jsonl_data = file.readlines()

    json_data = []

    for line in jsonl_data:
        data = json.loads(line)
        text = data['text']

        parts = text.split('<bot>:')
        instruction = parts[0].strip().replace('<human>:', '').strip()
        output = parts[1].strip() if len(parts) > 1 else ''

        json_entry = {
            'instruction': instruction,
            'input': '',
            'output': output
        }

        json_data.append(json_entry)

    with open(output_file, 'w') as file:
        json.dump(json_data, file, indent=2)
```

The final output is `train_data_llama_factory_George_MacDonald.json`.

## Conclusion

By following this process and using these Python functions for the transformations, you can start with the full text of a novel and break it down into a set of several hundred training examples. Each example prompts the model to rephrase a passage of AI-generated text back into the distinctive style of the original author.

The goal is that after training on a sufficient number of these examples, the model will gain the ability to write new passages that capture MacDonald's voice and storytelling flair. Larger training datasets would likely lead to better results, but this method provides a way to generate a reasonably-sized set of custom training examples when adapting a model to emulate a specific author's style.