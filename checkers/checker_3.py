def check(res):
    right_x = -3.5387636
    right_z = -26.2756567
    
    for element in res["elements"]:
        coords = element.get("coords")
        if element["type"] == "line":
            delta_x = abs(right_x - (-coords["x"] / coords["y"]))
            delta_z = abs(right_z - coords["z"] / coords["y"])

            if delta_x + delta_z < 1e-7:
                return 0
            

    return 1