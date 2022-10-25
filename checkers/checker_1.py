from utils import check_line


def check(res):
    x, y, z = 10.28838179695913, -6.858921197972754, 30.144958665090243

    return check_line(res["elements"], [x, y, z])
