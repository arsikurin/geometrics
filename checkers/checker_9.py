from utils import check_point


def check(res):
    right = [[-172.3684288898976, -4.94476299083459, 24.823285023294982],
             [-62.02129357204571, -84.37797118548922, 5.465181408154309]]

    return check_point(res["elements"], right)
