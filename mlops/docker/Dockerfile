FROM nvidia/cuda:12.3.2-cudnn9-runtime-ubuntu22.04
RUN apt-get update && apt-get upgrade --yes
RUN apt-get install -y --no-install-recommends python3.10 python3.10-venv

RUN python3.10 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

COPY requirements.txt .
# Note: In the next line --mount=type=cache,target=/root/.cache is to avoid
# re-downloading packages with pip when just one package changes
# It is useful for local development
RUN --mount=type=cache,target=/root/.cache pip install -r requirements.txt

COPY llm_prompts-*.whl .
RUN pip install --verbose --no-deps llm_prompts-*.whl

CMD ["python3", "-m", "llm_prompts"]