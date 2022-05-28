# encoding=utf8
# 处理笔记中的图片...
import re
from pathlib import Path
from datetime import datetime

from cv2 import goodFeaturesToTrack

path = Path(__file__).parent.absolute()
file_name = r"E:\coding\GolangCoding\notes\go 包的使用.md"

strr_t = datetime.now().strftime("%Y-%m-%d %H-%M-%S")

haeder1 = f"""
---
title: Golang 学习笔记——go 包的使用
date: {strr_t}
updated: {strr_t}
categories: Golang
tags: [Golang, Coding]
description: Golang 学习笔记——go 包的使用
index_img: https://cdn.jsdelivr.net/gh/WhaleFell/GolangCoding@master/icon_img.png
banner_img: https://cdn.jsdelivr.net/gh/WhaleFell/GolangCoding@master/icon_img.png
---
"""


def read_md() -> str:
    with open(f"{Path.joinpath(path,file_name)}", mode="r", encoding="utf8") as md:
        content = md.read()
    return content


def replare(md) -> str:
    return re.sub(r"\[\]\(img", "[](https://cdn.jsdelivr.net/gh/WhaleFell/GolangCoding@master/notes/img", md, count=0)


def main():
    global header1
    md = read_md()
    with open("test.txt", mode="w", encoding="utf8") as f:
        f.write(haeder1+"\n"+replare(md))


if __name__ == "__main__":
    main()
