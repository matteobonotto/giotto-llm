import os
from google.cloud import storage

# Initialize the Google Cloud Storage client
client = storage.Client()

# Define bucket name and blob (object) name
bucket_name = "public-models-and-datasets"
blob_name = "datasets/ray-v5.pickle"
# model folder
folder_path = "models/merged_llama_1b_v7"
destination_dir = "models/llama"
# define destination folders
destination_folder = "synth_data"
destination_path = os.path.join(destination_folder, "ray-v5.pickle")

# Download the dataset from GCS
try:
    bucket = client.get_bucket(bucket_name)
    blob = bucket.blob(blob_name)
    blob.download_to_filename(destination_path)
    print("Dataset downloaded successfully.")
except Exception as e:
    print(f"Failed to download dataset: {e}")

# Download the model from GCS
try:
    bucket = client.get_bucket(bucket_name)
    # List all blobs (files) in the folder
    blobs = client.list_blobs(bucket_name, prefix=folder_path)

    for blob in blobs:
        # Construct the relative path for the local file
        relative_path = os.path.relpath(blob.name, "llama_3_2_1B_instruct")
        local_path = os.path.join(destination_dir, relative_path)

        # If the blob name ends with '/', it's a folder (skip it)
        if blob.name.endswith('/'):
            continue

        # Ensure the local directory exists
        local_dir = os.path.dirname(local_path)
        if not os.path.exists(local_dir):
            os.makedirs(local_dir)

        # Download the file
        print(f"Downloading {blob.name} to {local_path}...")
        blob.download_to_filename(local_path)

    print("Model downloaded successfully.")
except Exception as e:
    print(f"Failed to download model: {e}")
