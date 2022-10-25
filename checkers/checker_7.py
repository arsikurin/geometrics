def dist(coords1, coords2):
    x1, y1, x2, y2 = coords1["x"], coords1["y"], coords2["x"], coords2["y"]
    return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)


def eq(a, b):
    return abs(a - b) <= 1e-6


def check(res):
    lenghts = [3.56 ** 2, 6.32 ** 2, 4.88 ** 2]
    points = []
    for element in res["elements"]:
        if element["type"] != "point": continue
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
