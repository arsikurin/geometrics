from checkers.utils import check_point


def check(res):
    right = [[1.063273680834214, 11.18063510094251, 2.989713460868993],
             [1.1861271383876528, -21.375531150718725, 2.989713460868993]]

    return check_point(res["elements"], right)
