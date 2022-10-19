def check(res):
    right_x = -7.064979330151781/-2.449731068127919
    right_z = -31.513694449807144/-2.449731068127919
    
    for element in res["elements"]:
        coords = element.get("coords")
        if element["type"] == "line":
            delta_x = abs(right_x - (-coords["x"] / coords["y"]))
            delta_z = abs(right_z - coords["z"] / coords["y"])

            if delta_x + delta_z < 1e-7:
                return 0
            

    return 1