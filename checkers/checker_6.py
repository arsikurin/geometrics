def check(res):
    x, y, z = -0.05330432415080688, 3.7617758559976604, 7.765464423969378
    x2, y2, z2 = 3.7563359269438426, 0.20913953641302374, 12.089112538084052

    variants = [[-x / y, z / y], [-x2 / y2, z2 / y2]]
    for element in res["elements"]:
        coords = element.get("coords")
        if element["type"] in ["line", "ray"]:
            delta_x = -coords["x"] / coords["y"]
            delta_z = coords["z"] / coords["y"]

            deltas = [abs(delta_x - right[0]) + abs(delta_z - right[1]) for right in variants]

            if min(deltas) < 1e-8:
                return 0

    return 1
