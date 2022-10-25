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
            except Exception as err:
                print(err, file=sys.stderr)
                continue
        print(3, end="")

    return wrapped


@retry
def main():
    problem_id = sys.argv[1]
    ggb_base64 = "UEsDBBQAAAAIAAFkWVXGcIogHQUAAGImAAAXAAAAZ2VvZ2VicmFfZGVmYXVsdHMyZC54bWztWl9z4jYQf+59Co2e2oeAbTCQTJyb3M10mplc7qbJdPoqjDBqjORacjD59F1JxjYB7hIDB/SOB+SV9W9/v9VqJfnyfT6N0RNNJRM8wG7LwYjyUIwYjwKcqfHZAL+/encZURHRYUrQWKRTogLs65JlPZBafc/UJkkSYCg+pSqdY5TEROk6AZ5hhHLJLri4I1MqExLS+3BCp+RWhESZZiZKJRft9mw2ay06bIk0akObsp3LUTuKVAtSjGDUXAa4eLiAdpdqzzqmnuc4bvvvT7e2nzPGpSI8pBiBRiM6JlmsJDzSmE4pV0jNExrgRDCuMIrJkMYB/qIl9Os4pfQ3jIpKAJSDr979ciknYobE8B8aQp5KM2i6qGeEti4Drz+KWKQoDXC/jxHgqpNhgD3fB7ziZEIC7NjCMZnTFD0RaKHIIZkSoalvcscklkXDpqdPYkTtm25RnjMgCOBEUlGgwmm5GMmE0hGMGhc6wgMQMzck11o0qt+z56JFv56r5nGRXQwsFCIdSZQH+I7cYTQv0mebQpHLdgHs6yAe0YTyERRawtlthHNvYHDWCeCsk33D3Bjkor2Dgtz7QUGGWbwHlD/zOrZeI2xdD1wDqGTSn65iCd8b/ieNYNR1lDung/JJYLxsw91G6EIwAPrA//Eha8CyGEr9H+BQTJOY5t8X+JjxCsRbI5Sge80ijDroOho7hMuAfteBrrW18KkJCx85lRDBgVmUlfTDH2wEy5NpTECIyBTg6fYHtgX6L18ijQFnDMpsTcQ446HWqgT3Y5Y+1dnodJ1D8FG12XgG7ImMzVhKGmmpxOV+IVem3Syo+6FMm+Z10xaZinWnN1zBlgsAg2HKFb0eKU0eoKPP/CElXOpt10szgg1QWvdhBSfFGHjhE1dfbTG3UgLbwM224J+CLfwPLWEH3pI/kbTkqc5ps2hsY7zQAiM5MLFvWDrqQGwfOJ2ycW9lRL1mjsFzuuvRa/WP2IieQD1RwfBXIVbhx0kEg9/TS66J4GFlo5IRvoP9UDyPajP6y0Iu+ehbPrZXYzOjNbSWtqx+x5DqQ3MvDNx17M/tnjuu24OTh6O1d43w0uZHQ2wzKoxtlLhPjI9i1rw+ztyMZyg4C6vNi5VKJLs/vcebNpKURZRblwwOxDFtzCGBlp+1pK85ctfIc0jg7bNOINtUB61SlqNrW+PaFrz2bNKxSdcmfoHeN5hNwLXVYugXi0O32XbqlDzJ/jnfWVx9TMbDsylNa47hbiGXtuNb1wA6ZMunWjJmIyB7ygDOM8B5SmAl1TH5UIo4U3B1BzdivLq6swY3YyM10UEYjG/Mck2sRQ9NRMqeBVclWEjb63VsLvmWDjrWEe19LcZ8ldfaZNObLbhmq9s5Z8KjuJqM11aqGLAXBKbQ6uHi14mBgRheei1v0HEHfsfpu/1zf9B7JU/uoOLJvtiOpk3zEehbnY8kDasjVghwNzAJvO2Uy2LNddx+1+94557vnp934QHGvuud4O9lRrWrOcZTRGMBK0X3dkAYizCT1bG3lUqEwCQbBcZHG62QLGcxI+l8tae9QaxoXgUMD0aofbpwhOHgZlUA9qga2o2Vat8HWGXGDFDk8F0JnCGYThj/QMLHKBUZL0y7NoLdqF4sPse4vxoKEVPYCS/U+rCQa/fSKyv/JoCKFfyQewX4fid8HIp8abH6xtWarGbArRFqt8VrZsDrtVxdkc4ObgpNzubeeM+5NkKpE9CufVLVXny/dfUfUEsDBBQAAAAIAAFkWVUVijgsdgMAAEgRAAAXAAAAZ2VvZ2VicmFfZGVmYXVsdHMzZC54bWztmN1y0zoQgK8PT6HRPfFP7KTu1GUynItzZoApww23qr1JBLZkJCWO+2q8A8/E6qetAy2QTigDQy6y+vHuWt+u15LPnu3ahmxBaS5FSZNJTAmIStZcrEq6McunJ/TZ+ZOzFcgVXCpGllK1zJQ0t1fe6GFvMk+dNuu6kuLlLRg1UNI1zFidkvaUkJ3mp0K+Yi3ojlXwplpDy17IihlnZm1MdxpFfd9Prh1OpFpFaFNHO11Hq5WZoKQE71rokobGKdrd0+6nTi+N4yR6+/KF9/OUC22YqIASXFENS7ZpjMYmNNCCMMQMHZS0koJXU/TRsEtoSvq/MLhMqOwtkmqjtqgflEs6TfKYnj/556ySUtWayF1JEYIcvLjyokeySMvPbf3c1s/1frD3g70bjKxBvZY9kZfv0HFJjdqg13BDruOuwennspGKqJKm6AFDlsQoL1EWKcai6dYMLU6S2P+SrIiTZJakXr9hAyiyZWg0eGUbIytn0o0uWaODL+f8pazBz2ThesExHSwZbQADj851B1C7lueJy8IkGFxCje1xAW/M0AAxa169F6AxnvlIyTb+43UNNi+9DnwQXkXb/5J2TGEqGcWrMM9XILZITCpNdrG7iQEFWruyPZufu8T1BxQ4e2UFDjt1XIniO7LwGgt/4SL1YupF5kUeiJ1FIXm+SiO243r6703QFqE7ypx46jLn0ECjeySJ/xhl+7yFGP+siGLy/KqYktAGXPWnj9/G7R7MiikDmjMxenyf24kvyc9+B/I/k/v9ING+gBG/C9ff44dl9UH8isIBTBOUiNDJmxKVHwvjktmXVzBxb+W7i1gAdV+mhhLvC7av16GGf7cYdLIZ1lArKW65joZu0U4D2oc8SYeGI8mnLh65f2OMMnqSBWp5MYuzWXa02Dw0xQ8iu1DVmrdQA9tHi7F/LLRp4t/G2dyhteLPYHsxYEXmWB3GXB8vZV3JwJsvPNf0j8nZC8V1u081eUSqM1+YPdUCe78hVQHmZp2vbHtcVfO/VfUQlh82rHY7sLDU19f9MVOfoMcsjbOssL/5LMlPkgxPNEcCdIx9KW+7hlfc/NBJ485zhh30h4nBiysUwduhRw+ymHkx9+LEi+K7OxG9UUs8ed+1Uw5T+0HOHhZk1LtzrzyZ/2jW3xp+lN3yWOmbu+Vo9OEguv5Kcf4ZUEsDBBQAAAAIAAFkWVXWN725GQAAABcAAAAWAAAAZ2VvZ2VicmFfamF2YXNjcmlwdC5qc0srzUsuyczPU0hPT/LP88zLLNHQVKiuBQBQSwMEFAAAAAgAAWRZVSoZ8azhBgAA+BUAAAwAAABnZW9nZWJyYS54bWztWG1v2zYQ/tz9CkKfNqCWJerNLuwOSTFsA9quaLph2Ddaom0tsqRJdGwX/fF77ijJdtKiSVNswDAnEsXj8d7veNLs+/2mEDe6afOqnDu+6zlCl2mV5eVq7mzNcjRxvn/+zWylq5VeNEosq2ajzNyJCHPYh5mbSN6t6nruAH2jTXNwRF0oQ3vmzs4ReTZ3PLVMIhmrUZoEySj0vOVoMQ3lKNHxUvtJsJDB0hFi3+bPyuq12ui2Vqm+Std6o15WqTLMcG1M/Ww83u12bi+aWzWrMbi3432bjVerhYvREdCvbOdO9/AMdM927wLeJz3PH//+6qXlM8rL1qgy1Y4g3bf582+ezHZ5mVU7scszs4al4iB2xFrnqzWskXiRI8aEVcMktU5NfqNb7D2ZsvZmUzuMpkpaf2KfRDEo5ogsv8kz3cBSrgzkZBJ5kyhKwgj2DRxRNbkuTYfsd0zHPbnZTa53li49McvQmyZwVd7mi0LPnaUqWuiVl8sGxh3mrTkUeqHA1jRbzI8S+U/5Dyj5e+DDYw6MQLbAmuc9pSvBFUVYIHFOeEe+dISpqoIpe+KD8EXk4RL+VDwVcQKIFH4kQkAmgCQiIFjkhyIQhOIHIgwxhgT2Y6zQMu5gJ3wfK0J6QkohfSEDTKNIREBLaK8Ebjxleh4uwoZEuAKCBQEuhgUhLklPIBRZMpAjCmJ+ivg+oT3gEoHfB8FLgIVTsCNAlPgigCSYJ54AXZCHxKxN6An690VITGQi5EQwVabvwUb3d08HuOWf3jvRx7wT42K33fJOeO4buMKDbhDQg5o8wIAEhcdo6pFhMLASnkduwRBZHChIUyjJg8Vh12FA4D5Ow16/4CH6TU70Ax5FDgYKCgyBILnxAPlpCLtpbKccbh7CxkLJ+RgQS4ioRyoDY3yBMjDBwNWm6EOY9ix9GYH9fXmeBuaDeR7VnIC9UYu5c/Hyxx8u317cFUBGn1D6kbb+qKXBi//5usMyeFA23imWX8AxPsvDr6NwOLk3e18iT/5hnglAHyk9dkSe8vh1HDH9jCNm4/7MnHUSiXZNuF3AG71B++CJmOtGf17hsKITBLc7hR0hW1dtPoix1gU6ok5g5piX9dZ0XDp4uqHmgDmaCuiq4J6m25BV6fXlIFi3RasWnceRLrqAY7Nhu4KzXuTJrFALXaCxuyKrCXGjCoo35rCsSiP6GhETbDbmvmemt2mRZ7kqf4OZ+hbj9Xaz0A3cg8eKtGQitF0MDRJXmmODFFictKqa7OrQwqxi/4dusDuU0p2e/BCOB7sShMH5CtWUVFFAhNOzlSl1Y4dPrE0sa31zpY2B/q1Qew2ndrZbNdyZsfHp+ef2sioA6ZbrKi/NC1WbbcNNMWRoSKmLclVoNiW7GY1jer2o9ldsQwlxiNa7Q03FwAqwWL2oiqoRCFcZoWEEMR5RGWlkHJJswEKIAgd3YHSOIqLDuj9FjgCDR+DQyFjwshWt0xTyWS07Kmqft7Ybh0FPg4pDhLrQbZmbl/3E5On1UVPaYAMAlNls5zQ7lMfTnI1vBd+sS4s+FDdVpm0Ys4Vn47P12bVuSl3YqCvh+G21bS26lYzF3rb6jTLrizJ7q1fI2TeKKoyBIBb1qGCm03yDjRbeWVpRFPwKxSw006tG9waxwlg/dFKKtm60ytq11siKzhs2J45oDJ6Ne/FnODsLzbVzk6OmwGcbtbe+MxrFwuK3aZPXFN5igTJ4rY8RnOUtURgAhA2LtFANxaEq4QtDfsDL29asKwQW9ihDECoAe4jc0nte70kc4ADOnW9H0vXCSRhEk6di5LlI5CCJkuA7eJfjnlOH2elCb/DecgY/kuPigQgQ1eJP1Kuhmtv1owewPMR+gpcahD4NXf4IVdRrNVi1UAcqUSdFjum9Oo8Z8iC8ymborOlSPNaa7EWNC0uMB7zJHrgAnARPirJAO5nepYWxclQFLLTLawu95ePO9a2AN4+2RCHD9GhOR7zvSiwSwtrxMxa9/I9Y9OJRFvXdCcoibCndT9gwrTYbVWai5J7irTqw3ezxrDwOdOUPjq22pl9YWkrd/jveaECpt/XyM7440f3UGeel/6GO6KHDMWDWKLf4RoDqCrdaKakHw8NPeZZp26pU+NSSG1jMT/oz86/SEuBOa+4g74scOINFCvLgzyXVS+o0UFXuVNhrrWs6B38p3zWqbOlr0O3SepYGyRD3kCRwJzFnhdf9rCtHaAxCGScTvNUHfMh3Tcv9EuTFv5Qg/fH7iAT5wmSQbmSTgQ38qHR4cScdVvdPh9X/6fCQdJCunCaxDCViDoKEbmSzgR0Yu/gI6QdS+vgEhd8dh8Ijx1OeW/ruW+nzvwFQSwECFAAUAAAACAABZFlVxnCKIB0FAABiJgAAFwAAAAAAAAAAAAAAAAAAAAAAZ2VvZ2VicmFfZGVmYXVsdHMyZC54bWxQSwECFAAUAAAACAABZFlVFYo4LHYDAABIEQAAFwAAAAAAAAAAAAAAAABSBQAAZ2VvZ2VicmFfZGVmYXVsdHMzZC54bWxQSwECFAAUAAAACAABZFlV1je9uRkAAAAXAAAAFgAAAAAAAAAAAAAAAAD9CAAAZ2VvZ2VicmFfamF2YXNjcmlwdC5qc1BLAQIUABQAAAAIAAFkWVUqGfGs4QYAAPgVAAAMAAAAAAAAAAAAAAAAAEoJAABnZW9nZWJyYS54bWxQSwUGAAAAAAQABAAIAQAAVRAAAAAA"

    check = __import__(f"checker_{problem_id}").check

    get_xml(ggb_base64)
    response = get_json()

    print(check(response), end="")


if __name__ == "__main__":
    main()
