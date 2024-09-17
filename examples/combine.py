"""Combine the nested Go files in this directory into snippets for use in JavaScript."""

import json
import os
import re

TITLE_PATTERN = re.compile(r"(\d+)_(.*)")


def to_title(filename: str) -> str:
    """Create a title from the Go filename.

    Example: "01_hello_world.go" -> "Hello World"
    """
    if match := TITLE_PATTERN.match(filename):
        filename = match.groups()[1]
    title = " ".join((part[0].upper() + part[1:] for part in filename.split("_")))
    return title


if __name__ == "__main__":
    folders = [filename for filename in os.listdir() if os.path.isdir(filename)]

    go: list[dict] = []
    for folder in folders:
        filename = folder + "/main.go"
        with open(filename) as fin:
            code = fin.read()
        # replace tabs with spaces to make it easier to edit in the browser
        code = code.replace("\t", "    ")
        title = to_title(folder)
        go.append({"title": title, "code": code})

    with open("../examples.js", "w") as fout:
        fout.write("const EXAMPLES = ")
        fout.write(json.dumps(go, indent=4))
