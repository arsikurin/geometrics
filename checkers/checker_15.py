from checkers.utils import check_point


def check(res):
    right = [[-2.236210907161168, 4.985878791110917, 1.0],
             [4.366884719355621, 3.285350558935912, 1.0]]

    return check_point(res["elements"], right)
