def check(res):
    x, y, z = 10.28838179695913, -6.858921197972754, 30.144958665090243
    right_x = -x / y
    right_z = z / y

    for element in res["elements"]:
        coords = element.get("coords")
        if element["type"] == "line":
            delta_x = abs(right_x - (- coords['x'] / coords['y']))
            delta_z = abs(right_z - coords['z'] / coords['y'])

            if delta_x + delta_z < 1e-8:
                return 0

    return 1