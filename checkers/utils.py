def check_point(elements, right_coords, precision=1e-6):
    for element in elements:
        coords = element.get("coords")
        if element["type"] == "point":
            delta_x = coords["x"] / coords["z"]
            delta_y = coords["y"] / coords["z"]

            deltas = [abs(delta_x - var[0] / var[2]) + abs(delta_y - var[1] / var[2]) for var in right_coords]

            if min(deltas) < precision:
                return 0
    return 1


def check_line(elements, right_coords, precision=1e-6):
    if not isinstance(right_coords[0], list):
        right_coords = [right_coords]

    for element in elements:
        coords = element.get("coords")
        if element["type"] in ["line", "ray"]:
            delta_x = -coords["x"] / coords["z"]
            delta_y = coords["y"] / coords["z"]

            deltas = [abs(delta_x + right[0] / right[2]) + abs(delta_y - right[1] / right[2]) for right in right_coords]

            if min(deltas) < precision:
                return 0
    return 1


def check_triangle(elements, lenghts, precision=1e-6, normalize_coords=False):
    def eq(a, b):
        return abs(a - b) <= precision

    def dist(coords1, coords2):
        x1, y1, x2, y2 = coords1["x"], coords1["y"], coords2["x"], coords2["y"]
        if normalize_coords:
            x1 = x1 / coords1["z"]
            y1 = y1 / coords1["z"]
            x2 = x2 / coords2["z"]
            y2 = y2 / coords2["z"]
        return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)

    points = []
    for element in elements:
        if element["type"] != "point":
            continue
        points.append(element)

    for element in points:
        coords = element.get("coords")
        B = []
        C = []
        for element2 in points:
            if element2 == element: continue
            coords2 = element2.get("coords")
            if eq(dist(coords, coords2), lenghts[0]):
                B.append(coords2)
            elif eq(dist(coords, coords2), lenghts[1]):
                C.append(coords2)

        for point1 in B:
            for point2 in C:
                if eq(dist(point1, point2), lenghts[2]):
                    return 0

    return 1
