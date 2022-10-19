def check(res):
    right_x = 0.5969876648683643/8.094409146651714
    right_z = 2.1959910673609344/8.094409146651714
    
    for element in res["elements"]:
        coords = element.get("coords")
        if element["type"] == "line":
            delta_x = abs(right_x - (-coords["x"] / coords["y"]))
            delta_z = abs(right_z - coords["z"] / coords["y"])

            if delta_x + delta_z < 1e-8:
                return 0
            

    return 1