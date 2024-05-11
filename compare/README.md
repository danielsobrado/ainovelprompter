# AI Writing Model Comparison

## Overview

I conducted the following test:

1. Created a prompt (`output1_prompt.txt`) detailing the prose and exact steps for a 1500-word text.
2. Asked multiple AIs to write the text three times each (`output1_1`, `output1_2`, and `output1_3`) to check differences between evaluations.
3. Tested AIs on [https://chat.lmsys.org/](https://chat.lmsys.org/), ChatGPT4 website, Claude 3 Opus website, and some local models in LM Studio.
4. Wrote a prompt (`evaluate1_prompt.txt`) to get feedback on the writing of the first iteration (`output1_1`) for each model.
5. Asked ChatGPT 4 and Claude Opus 3 to provide feedback (all files starting with "evaluation").

## Technical Details

- When running models on my local machine, I sometimes used quantized models to fit in memory (indicated in the output filename).
- I always used 8K context and 0.8 temperature.
- I couldn't find a quantized version of Midnight-Miqu-70B-v1.5, so I did it myself with llama.cpp and uploaded to [https://huggingface.co/drusniel/Midnight-Miqu-70B-v1.5-GGUF](https://huggingface.co/drusniel/Midnight-Miqu-70B-v1.5-GGUF).

## Observations

- Some models have higher variability between executions, while others use the same wording quite often.
- Significant differences in the number of words generated and the amount of dialogue, descriptions, and paragraphs.
- More work is needed to fine-tune prompts for each model to circumvent their limitations.
- ChatGPT is known for its overly flowery prose.

## Evaluation Feedback

- ChatGPT suggests a more "refined" prose, while Claude suggests less purple prose.

## Takeaways for Next Prompt

- Go for precision and power over purple prose.
- Vary sentence lengths and structures. Occasional fragments or one-word sentences can enhance pacing and punch.
- Stay firmly in one character's perspective to create a stronger emotional connection.
- Use strong, active verbs.
- Avoid common fantasy motifs. Consider ways to give it a unique twist.
- Keep the tone consistent to build tension.
- Give the narrator a distinct voice.
- Avoid a pacing that feels rushed.
- Consider starting with a gripping hook that drops the reader right into the action and mystery.
- Ensure each scene has a clear purpose that advances the characters and plot.
- Avoid overly long and complex sentences.
- Avoid cliched or melodramatic word choices.
- Avoid cliches like "It was a dark and stormy night".
- Avoid overusing adjectives/adverbs.
- Show character emotions through body language, dialogue, and actions rather than just stating them.
- Avoid a distant omniscient narrative voice. Use a closer or more colorful POV voice.
- Use strong verbs and avoid overusing "was/were".
- Avoid overly ornate descriptions. Aim for a more natural voice.

Another technique is to ask for feedback and then ask to rewrite the text using that feedback.

Feel free to ping me if you want to collaborate on further fine-tuning prompts for each model.