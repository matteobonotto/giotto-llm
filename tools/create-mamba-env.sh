mamba create --name unsloth_env -c pytorch -c nvidia -c xformers \
    python=3.10 \
    pytorch-cuda=12.4 \
    pytorch==2.4.1 \
    cudatoolkit \
    xformers

echo ">>> Remember to then run 'mamba activate unsloth_env'"