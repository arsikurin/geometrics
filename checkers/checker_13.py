from checkers.utils import check_point


def check(res):
    right = [[1.2461241921723998, 0.2659416462979376, 0.7781965628537533]]

    return check_point(res["elements"], right)