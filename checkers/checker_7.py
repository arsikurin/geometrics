from utils import check_triangle


def check(res):
    lenghts = [3.56 ** 2, 6.32 ** 2, 4.88 ** 2]

    return check_triangle(res["elements"], lenghts)
