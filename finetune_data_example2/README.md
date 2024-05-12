# Generating Custom Training Data to Fine-Tune a Language Model with Local LLM and LM Studio

This is a follow up to automate the more manual process used on [the previous example](https://github.com/danielsobrado/ainovelprompter/tree/main/finetune_data_example) provides an overview of the process for generating a fine-tuning JSONL file by leveraging a local language model (LLM) running on LM Studio. The goal is to paraphrase text from a given source file and create a dataset suitable for fine-tuning a language model.

## Process Overview

1. **Text Preprocessing:**
   - Read the source text file and split it into paragraphs.
   - Filter out paragraphs that contain less than 10 words.

2. **Paragraph Paraphrasing:**
   - Iterate over each paragraph and send it to the local LLM running on LM Studio for paraphrasing.
   - Use the LM Studio API to communicate with the local LLM.
   - Provide a system prompt to guide the LLM in paraphrasing the text while maintaining the original meaning and form.
   - Retry paraphrasing a paragraph up to 5 times if the generated response does not meet the specified criteria.

3. **Response Validation:**
   - Check the generated paraphrase against various criteria to ensure its quality and adherence to the desired format.
   - Validate that the response does not contain any unwanted phrases or patterns, such as "###", "Paraphrased Text", "Alternate Refinement", "Extended Alternate Refinement", or "[Note:".
   - Ensure that the response consists of a single paragraph and is not significantly shorter or longer than the original input.

4. **JSONL Generation:**
   - For each successfully paraphrased paragraph, create a JSON object containing the original text and the paraphrased version.
   - Format the JSON object in a specific structure, including the "text" field with the original and paraphrased text, and the "metadata" field with the source information.
   - Append each JSON object to a list, which will be used to generate the final JSONL file.

5. **Output and Summary:**
   - Write the generated JSONL data to a file named "output.jsonl".
   - Provide a summary of the paraphrasing process, including the total number of paragraphs, the number of successfully paraphrased paragraphs, and the number of skipped paragraphs along with their respective reasons.

## Rules and Considerations

- The local LLM should be running on LM Studio and accessible via the provided API.
- The system prompt plays a crucial role in guiding the LLM to perform the paraphrasing task effectively. It should instruct the LLM to maintain the original meaning and form while refining the prose.
- The response validation criteria are designed to ensure the quality and consistency of the paraphrased text. Adjust these criteria based on your specific requirements and the behavior of the LLM.
- Skipped paragraphs are those that fail to meet the specified criteria after the maximum number of retries. Keep track of the reasons for skipping paragraphs to identify potential issues or areas for improvement.
- The JSONL format is used to structure the fine-tuning data, with each line representing a JSON object containing the original and paraphrased text along with metadata.
- The summary provides insights into the paraphrasing process, helping you understand the effectiveness of the LLM and identify any bottlenecks or challenges.

## Lesson learnt during this process

When you fine-tune models, you significantly lobotomize them, causing them to lose their ability to perform well in many areas. They conducted an experiment using the Dolphin version of LLaMA 3, which is supposed to have better metrics than the original LLaMA 3.

The task was simple: given a text, determine whether it is a dialogue or not, and respond with only "yes" or "no". When asked, the original LLaMA 3 model performed the task perfectly. However, when the same question was posed to the Dolphin version, it made many mistakes and did not follow the prompt correctly.

Llama 3 version:

Dolphin finetuned version:

Generic fine-tunes should be approached with caution. A fine-tuned model, such as a LoRA (Low-Rank Adaptation) model, should be used for specific tasks, and the resulting model may no longer be suitable for other tasks.