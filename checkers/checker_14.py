from checkers.utils import check_point


def check(res):
    right = [[0.7688435340021545, 0.021212836292100246, 1.0]]

    return check_point(res["elements"], right)