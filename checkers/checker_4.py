from checkers.utils import check_line


def check(res):
    x, y, z = -0.5848532457, 1.0, -0.9503629013444963

    return check_line(res["elements"], [x, y, z])