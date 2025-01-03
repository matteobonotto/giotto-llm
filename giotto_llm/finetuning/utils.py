from typing import Literal, Type

from pydantic import BaseModel

from ..causal_lm.models import CausalLMWrapper
from ..multimodal.molmo import MolmoWrapper
from ..multimodal.qwen import QwenVLWrapper
from ..wrapper import ModelWrapper

MAP_WRAPPER: dict[str, Type[ModelWrapper]] = {
    "CausalLM": CausalLMWrapper,
    "Molmo": MolmoWrapper,
    "QwenVL": QwenVLWrapper,
}


class FinetuningConfig(BaseModel):
    """Fine-tuning configuation

    Args:
        model_id: The base model to fine-tune
        wrapper: the wrapper for the base-model
        output_dir: the directory to store outputs
        quantization: the quantization method
        dataset: the dataset to use for training
        transform_background_color: whether or not to transform the background color (0)
            during training.
        compress_colors: whether or not to do color compression during training.
        learning_rate: the initial learning rate
        per_device_batch_size: the batch size on each device
        gradient_accumulation_steps: the number of gradients to accumulate before updating
            model weights.
        num_train_epochs: the maximum number of training epochs
        neftune_noise_alpha: the noise to add to embeddings. If `None` add no noise.
        padding_side: Override padding side to either 'left' or 'right'. If `None` use wrapper defaults
        prompt_type: a type of text prompt.
        lora_target_modules: The target modules for Lora. If not specified, target all linear layers.
        training_set_size: the training set size. If unset, use the full dataset.
        evaluation_set_size: the evaluation set size. If unset, use a size corresponding to 20% of the training set size.
        logging_steps: Number of steps between each log entry. If not specified, log per epoch.
        eval_steps: Number of steps between each pass of the evaluation set. If not specified, evaluate once per epoch.
        low_memory: Enable options that save memory at the expense of performance.
        lora_dropout: the dropout to use in Lora
        lora_alpha: the alpha to use in Lora
        lora_r: the r to use in Lora
        early_stopping_patience: the early stopping patience
        save_total_limit: how many checkpoints to keep
        untie_word_embeddings: Untie word embeddings from the classifier head.
        eval_split_from_train: Use a random subset of the training set for evaluation instead of the kaggle eval set.
        gradient_checkpointing: Enable gradient checkpointing. Will also be enabled with the `low_memory` flag.
        disable_input_mask: Disable the masking of input tokens during training.
    """

    model_id: str
    wrapper: Literal["CausalLM", "Molmo", "QwenVL"]
    output_dir: str
    quantization: (
        Literal[
            "no", "4bit-nf4", "4bit-dq-nf4", "4bit-fp4", "4bit-dq-fp4", "8bit-6", "8bit-5", "8bit-4"
        ]
        | None
    ) = "4bit-dq-nf4"
    dataset: str = "combine-v155"
    transform_background_color: bool = False
    compress_colors: bool = False
    learning_rate: float = 3e-4
    per_device_batch_size: int = 2
    gradient_accumulation_steps: int = 4
    num_train_epochs: int = 5
    neftune_noise_alpha: float | None = None
    padding_side: Literal["left", "right"] | None = None
    prompt_type: str = "prompt_solve_short"
    lora_target_modules: list[str] | None = None
    training_set_size: int | None = None
    evaluation_set_size: int | None = None
    logging_steps: int | None = None
    eval_steps: int | None = None
    low_memory: bool = False
    lora_dropout: float = 0.2
    lora_alpha: int = 16
    lora_r: int = 8
    early_stopping_patience: int = 3
    save_total_limit: int | None = None
    untie_word_embeddings: bool = False
    eval_split_from_train: bool = False
    gradient_checkpointing: bool = False
    disable_input_mask: bool = False

    class Config:
        protected_namespaces = ()
        extra = "forbid"
