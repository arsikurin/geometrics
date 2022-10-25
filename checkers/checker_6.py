from utils import check_line


def check(res):
    x, y, z = -0.05330432415080688, 3.7617758559976604, 7.765464423969378
    x2, y2, z2 = 3.7563359269438426, 0.20913953641302374, 12.089112538084052

    return check_line(res["elements"], [[x, y, z], [x2, y2, z2]])
