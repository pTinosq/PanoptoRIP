import requests
from datetime import datetime
import os


def download_fragmented_mp4(url):
    """Download a full fragmented.mp4 stream to a local file."""
    headers = {"Range": "bytes=0-", "Referer": "https://cardiff.cloud.panopto.eu/"}

    filename = f"lecture_{datetime.now().strftime('%Y%m%d_%H%M%S')}.mp4"
    print(f"Downloading to: {filename}")

    with requests.get(url, headers=headers, stream=True) as r:
        r.raise_for_status()
        with open(filename, "wb") as f:
            for chunk in r.iter_content(chunk_size=8192):
                if chunk:
                    f.write(chunk)

    print(f"✅ Download complete: {filename}\n")


def download_from_txt(file_path):
    with open(file_path, "r") as f:
        urls = [line.strip() for line in f if line.strip()]
    for url in urls:
        try:
            download_fragmented_mp4(url)
        except Exception as e:
            print(f"❌ Error downloading {url}: {e}\n")


if __name__ == "__main__":
    path = input("Path to TXT file containing Panopto URLs: ").strip()
    if os.path.isfile(path):
        download_from_txt(path)
    else:
        print("❌ File not found.")
