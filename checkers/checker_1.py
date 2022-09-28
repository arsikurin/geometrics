def check(res):
    # тут должна быть проверка что не используются запрещенные штуки

    right_x = 10.28838179695913
    right_y = -6.858921197972754
    right_z = 30.144958665090243

    for element in res["elements"]:
        coords = element.get("coords")
        if element["type"] == "line":
            delta_x = abs(right_x - coords["x"])
            delta_y = abs(right_y - coords["y"])
            delta_z = abs(right_z - coords["z"])

            if delta_x + delta_y + delta_z < 1e-8:
                return 0

    return 1
