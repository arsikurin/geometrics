import base64
import os
import sys
import xml.etree.ElementTree as ET
from zipfile import ZipFile


def get_json():
    def convert_coords(coords):
        return {"x": float(coords["x"]), "y": float(coords["y"]), "z": float(coords["z"])}

    root = ET.parse("geogebra.xml").getroot()

    os.remove("geogebra.xml")

    response = {"elements": [], "commands": []}

    for tag in root.findall("construction/element"):
        element = tag.attrib
        coords = tag.find("coords")
        if coords is not None:
            element["coords"] = convert_coords(coords.attrib)

        response["elements"].append(element)

    for tag in root.findall("construction/command"):
        command = tag.attrib
        input_values = tag.find("input")
        if input_values is not None:
            command["input"] = input_values.attrib
        output_values = tag.find("output")
        if output_values is not None:
            command["output"] = output_values.attrib

        response["commands"].append(command)

    return response


def get_xml(ggbBase64):
    # декодирование base64
    decoded = base64.b64decode(ggbBase64)

    # сохраняем декодированную строку как zip-архив
    with open("temp.zip", "wb") as file:
        file.write(decoded)

    # получаем xml-файл и сохраняем его в текущей директории
    with ZipFile("temp.zip", "r") as zipped_file:
        zipped_file.extract("geogebra.xml")
    os.remove("temp.zip")


def retry(fn):
    def wrapped():
        for _ in range(3):
            try:
                return fn()
            except Exception:
                continue
        print(3, end="")

    return wrapped


@retry
def main():
    problem_id = sys.argv[1]
    ggb_base64 = sys.argv[2]

    check = getattr(__import__(f"checkers.checker_{problem_id}"), f"checker_{problem_id}").check

    get_xml(ggb_base64)
    response = get_json()

    print(check(response), end="")


if __name__ == "__main__":
    main()
