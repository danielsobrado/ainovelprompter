# Generating Custom Training Data to Fine-Tune a Language Model with Local LLM and LM Studio using ORPO

## Steps

1. Prepare your input data in JSONL format. Each line of the JSONL file should contain a JSON object with the following structure:
   ```json
   {
     "text": "<human>: Rephrase the following text in the style of George MacDonald: ...\n<bot>: ...",
     "metadata": {"source": "gutenberg"}
   }
   ```
   The `text` field should contain the prompt and the chosen response separated by `\n<bot>: `.

2. Set up the OpenAI client in the script:
   ```python
   client = OpenAI(base_url="http://localhost:1234/v1", api_key="lm-studio")
   ```
   Replace `"http://localhost:1234/v1"` with the appropriate base URL for your OpenAI API and `"lm-studio"` with your OpenAI API key.

3. Specify the input and output file paths in the script:
   ```python
   input_file_path = 'G:\\texts\\train_data_book1_Drusniel_Meta-Llama-3-8B-Instruct-Q8_0.jsonl'
   output_file_path = 'G:\\texts\\orpo_data_book1_Drusniel_Meta-Llama-3-8B-Instruct-Q8_0.csv'
   ```
   Modify the `input_file_path` to point to your JSONL input file and the `output_file_path` to specify the desired location for the generated CSV output file.

4. Run the script:
   ```
   python script_name.py
   ```
   Replace `script_name.py` with the actual name of your Python script file.

5. The script will process the JSONL file and generate a CSV file with the following columns:
   - `prompt`: The prompt text extracted from the JSONL file.
   - `chosen`: The chosen response extracted from the JSONL file.
   - `rejected`: A generated rejected response using the same prompt and model parameters.

   The script will save the CSV file every 100 new lines and print a small summary during each save, providing regular updates on the progress of the processing.

   If the script is interrupted or stopped, it will automatically resume processing from where it left off based on the existing output file.

6. Once the script finishes processing all the lines in the JSONL file, it will provide a final summary with the following information:
   - Total lines: The total number of lines in the JSONL file.
   - Written lines: The number of lines successfully written to the CSV file.
   - Skipped lines: The number of lines skipped due to retries exceeding the maximum limit.
   - Total retries: The total number of retries performed during the processing.
   - Skipped line details: The line numbers and corresponding failure reasons for the skipped lines.

## Notes

- The script uses the OpenAI API to generate rejected responses. Make sure you have a valid OpenAI API key and the necessary permissions to access the API.
- The script is set to retry generating a rejected response up to 5 times if the generated response does not meet certain criteria. You can adjust the retry limit and criteria as needed.
- The script saves the CSV file every 100 new lines to provide regular progress updates. You can modify the save interval by changing the condition `if written_lines % 100 == 0`.

That's it! You should now have a CSV file with the training data for ORPO, including the prompt, chosen response, and rejected response for each line in the JSONL file.