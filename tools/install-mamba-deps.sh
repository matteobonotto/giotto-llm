mamba info
echo "------------------------------"
mamba install -c conda-forge \
    numpy==1.26.4 \
    scipy \
    scikit-learn \
    matplotlib \
    easydict \
    psutil \
    transformers \
    sentencepiece \
    polars \
    datasets \
    build \
    black==24.8.0 \
    pytest==8.3.3 \
    isort==5.13.2 \
    mypy==1.11.2 \
    protobuf \
    fire \
    google-cloud-storage

pip install --force-reinstall "unsloth[colab-new] @ git+https://github.com/unslothai/unsloth.git"
pip install --no-deps "trl<0.9.0" peft accelerate bitsandbytes huggingface-hub
pip install --no-deps --verbose --editable .