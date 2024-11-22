export DOCKER_DEV_DIR=mlops/docker

poetry export \
    --with dev \
    --format requirements.txt \
    -o mlops/docker/requirements.txt

rm -r dist/
poetry build
cp dist/llm_prompts-*.whl ${DOCKER_DEV_DIR}

echo "The contents of ${DOCKER_DEV_DIR} before running docker build are"
ls -alh ${DOCKER_DEV_DIR}