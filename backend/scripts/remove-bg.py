import sys
import requests
from io import BytesIO
from PIL import Image
import numpy as np

WHITE_THRESHOLD = 245


def remove_white_bg(img: Image.Image) -> Image.Image:
    img = img.convert("RGBA")
    data = np.array(img)

    r, g, b, a = data.T
    white = (r > WHITE_THRESHOLD) & (g > WHITE_THRESHOLD) & (b > WHITE_THRESHOLD)

    data[..., 3][white.T] = 0
    return Image.fromarray(data)


def download_image(url: str) -> Image.Image:
    resp = requests.get(url, timeout=15)
    resp.raise_for_status()
    return Image.open(BytesIO(resp.content))


def main():
    if len(sys.argv) != 2:
        print("usage: python remove_bg.py <image_url>", file=sys.stderr)
        sys.exit(1)

    url = sys.argv[1]
    img = download_image(url)
    out = remove_white_bg(img)

    # Write PNG to stdout (binary)
    buffer = BytesIO()
    out.save(buffer, format="PNG")
    sys.stdout.buffer.write(buffer.getvalue())


if __name__ == "__main__":
    main()

