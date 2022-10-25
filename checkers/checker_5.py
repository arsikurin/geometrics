from checkers.utils import check_line


def check(res):
    x, y, z = -0.0735533113796428, 0.9972912866284801, 0.2705623989737358

    return check_line(res["elements"], [x, y, z])
