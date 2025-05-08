# Build Stage
FROM python:3.8.10-slim AS builder

# Set environment variables for lighter image
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1 \
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    POETRY_VERSION=1.8.2

# Install Poetry
RUN pip install "poetry==${POETRY_VERSION}"

# Config Poetry to not create envs
RUN poetry config virtualenvs.create false

WORKDIR /app

# Copy package files only 
COPY pyproject.toml poetry.lock* ./

# Install production dependencies using Poetry
RUN poetry install --no-dev --no-interaction --no-ansi

# Production stage
FROM python:3.8.10-slim AS production

# Set environment variables for the runtime environment
ENV PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

WORKDIR /app

# Copy installed packages from builder
COPY --from=builder /usr/local/lib/python3.8/site-packages /usr/local/lib/python3.8/site-packages

# Copy installed executables from builder
COPY --from=builder /usr/local/bin /usr/local/bin

# Copy application code
COPY . .

# Expose the application port
EXPOSE 8000

# Running uvicorn on all interfaces
ENTRYPOINT ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]