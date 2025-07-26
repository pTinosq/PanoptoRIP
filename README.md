# PanoptoRIP

PanoptoRIP is a CLI tool to download videos from Panopto into MP4 files for offline viewing. Originally created for Cardiff University students who want to watch lectures on the go.

> ⚠️ Use responsibly. This is for **personal use only** and assumes you have rightful access to the content.

---

## Features

- Download individual Panopto videos or batches from a list
- Save videos as MP4 files for offline viewing
- Specify output folder for downloads

---

## How to Use (Non-technical Instructions)

### 1. Get the Panopto Video URL

- Open the video in your browser and hit play (Chrome recommended).
- Open **Developer Tools (F12)** → Click on the **Network tab**.
- Click on the **Fetch/XHR** button and look for any line which says `fragmented.mp4`.
- Right-click the line and select → _Copy > Copy URL_.

### 2. Download a Single Video

```sh
panoptorip rip --single "<PASTE_URL_HERE>"
```

Or with a custom output folder:

```sh
panoptorip rip --single "<PASTE_URL_HERE>" --output my_videos
```

### 3. Download Multiple Videos (Batch)

- Paste one URL per line into a text file (e.g. `urls.txt`).
- Run:

```sh
panoptorip rip --batch urls.txt
```

Or with a custom output folder:

```sh
panoptorip rip --batch urls.txt --output my_videos
```

### 4. Output

- Each video will be saved as `lecture_YYYYMMDD_HHMMSS.mp4` in the specified folder (default: `lectures`).

---

## Developer Instructions

### Prerequisites

- [Go 1.18+](https://golang.org/dl/)

### Build from Source

```sh
git clone https://github.com/<your-username>/PanoptoRIP.git
cd PanoptoRIP
go build -o panoptorip main.go
```

### Run

```sh
./panoptorip rip --single "<PASTE_URL_HERE>"
```

Or see all commands:

```sh
./panoptorip --help
```

---

## Notes

- The tool uses a `Range` header to request the full file, bypassing Panopto's streaming restrictions.
- The `Referer` is hardcoded to `cardiff.cloud.panopto.eu`. Change it in the code if needed for other institutions.
- Works on macOS, Linux, and Windows

---

## License

MIT - do what you want.
