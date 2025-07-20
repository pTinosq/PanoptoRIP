# PanoptoRIP

At Cardiff University, we had to use Panopto for lecture recordings. This sucks because it won't let us download the videos for offline watching. I used to travel a lot during my studies and wanted to keep up with my lectures while on the move. That's what this script is for.  

> ⚠️ Use responsibly. This is for **personal use only** and assumes you have rightful access to the content.

---

## Setup

1. Install the required packages:

```bash
pip install -r requirements.txt
```

2. Run the script:

```bash
python main.py
```

---

## How to use

1. Open the video in your browser and hit play (Chrome recommended).
2. Open **Developer Tools** → **Network tab**.
3. Filter for **XHR** requests and look for a request with `fragmented.mp4`.
4. Right-click → *Copy URL*.
5. Paste one URL per line into a text file (e.g. `urls.txt`).
6. Run the script and provide the path to the `.txt` file when prompted.
7. Each video will be saved as `lecture_YYYYMMDD_HHMMSS.mp4`.

---

## Example

Path to TXT file containing Panopto URLs: urls.txt
✅ Download complete: lecture_20250720_101530.mp4
✅ Download complete: lecture_20250720_101601.mp4

---

## Notes

- This script uses a `Range` header to request the full file, bypassing Panopto's streaming restrictions.
- The `Referer` is hardcoded to `cardiff.cloud.panopto.eu`. Change it if needed for other institutions.
- Works on macOS, Linux, and Windows (Python 3.7+).

---

## License

MIT - do what you want.
