from utils import check_triangle


def check(res):
    right = [28.090399999999995, 31.397642480441633, 45.455457024427545]
    return check_triangle(res["elements"], right, normalize_coords=True)
