import os
from google.cloud import storage

# Define bucket name and blob (object) name
bucket_name = "public-models-and-datasets"
blob_name = "datasets/ray-v5.pickle"
# model folder
folder_path = "models/merged_llama_1b_v7"
destination_dir = "models/llama"
# define destination folders
destination_folder = "synth_data"
destination_path = os.path.join(destination_folder, "ray-v5.pickle")

# Ensure destination folders exist
os.makedirs(destination_folder, exist_ok=True)
os.makedirs(destination_dir, exist_ok=True)

# Download the dataset from GCS
try:
    # Initialize the anonymous Google Cloud Storage client
    client = storage.Client.create_anonymous_client()
    bucket = client.bucket(bucket_name)
    blob = bucket.blob(blob_name)

    # Download the dataset
    blob.download_to_filename(destination_path)
    print(f"Dataset downloaded successfully to {destination_path}.")
except Exception as e:
    print(f"Failed to download dataset: {e}")

# Download the model from GCS
try:
    # Initialize the anonymous Google Cloud Storage client
    client = storage.Client.create_anonymous_client()
    bucket = client.bucket(bucket_name)

    # List all blobs (files) in the folder
    blobs = bucket.list_blobs(prefix=folder_path)

    for blob in blobs:
        # Construct the relative path for the local file
        relative_path = os.path.relpath(blob.name, folder_path)
        local_path = os.path.join(destination_dir, relative_path)

        # If the blob name ends with '/', it's a folder (skip it)
        if blob.name.endswith('/'):
            continue

        # Ensure the local directory exists
        local_dir = os.path.dirname(local_path)
        os.makedirs(local_dir, exist_ok=True)

        # Download the file
        print(f"Downloading {blob.name} to {local_path}...")
        blob.download_to_filename(local_path)

    print("Model downloaded successfully.")
except Exception as e:
    print(f"Failed to download model: {e}")
