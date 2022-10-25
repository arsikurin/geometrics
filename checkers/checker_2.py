from utils import check_line


def check(res):
    x, y, z = -1.0, 0.34674285, 4.460550127204706

    return check_line(res["elements"], [x, y, z])
