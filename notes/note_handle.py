# encoding=utf8
# 处理笔记中的图片...
import re
from pathlib import Path

path = Path(__file__).parent.absolute()
file_name = r"E:\coding\GolangCoding\notes\go type关键字.md"


def read_md() -> str:
    with open(f"{Path.joinpath(path,file_name)}", mode="r", encoding="utf8") as md:
        content = md.read()
    return content

def replare(md)->str:
    return re.sub(r"\[\]\(img","[](https://cdn.jsdelivr.net/gh/WhaleFell/GolangCoding@master/notes/img",md,count=0)

def main():
    md = read_md()
    with open("test.txt",mode="w",encoding="utf8") as f:
        f.write(replare(md))

if __name__ == "__main__":
    main()
