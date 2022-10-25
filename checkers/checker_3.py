from checkers.utils import check_line


def check(res):
    x, y, z = -1.0, -0.282584572574, 7.425095226980985

    return check_line(res["elements"], [x, y, z])
