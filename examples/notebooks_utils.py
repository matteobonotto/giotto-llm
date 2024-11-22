from typing import Callable

from llm_prompts.prompts.grid_formatter import GridFormatter
from llm_prompts.prompts.text_prompts import TextPromptBase

from transformers import AutoTokenizer



def get_clean_prompts(
    tasks: dict,
    model_id: AutoTokenizer,
    max_seq_length: int,
    max_num_tasks: int,
    prompt_fn: TextPromptBase,
) -> dict[str, tuple[bool, str]]:
    tokenizer = AutoTokenizer.from_pretrained(model_id)

    all_messages: dict = {task_id: prompt_fn(task, idx_i=0) for task_id, task in tasks.items()}
    _dataset = {
        task_id: tokenizer.apply_chat_template(
            messages, # Note: only using first input
            tokenize=False,
            add_generation_prompt=False,
        )
        for idx_task, (task_id, messages) in enumerate(all_messages.items())
        if idx_task < max_num_tasks
    }
    print(f">>> {_dataset=}")
    return {task_id: (len(tokenizer.encode(conv)) < max_seq_length, conv) for task_id, conv in _dataset.items()}
