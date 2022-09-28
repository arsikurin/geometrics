import sys

from get_json import get_json
from get_xml import get_xml

problem_id = sys.argv[1]
ggbBase64 = sys.argv[2]

check = getattr(__import__(f"checkers.checker_{problem_id}"), f"checker_{problem_id}").check

get_xml(ggbBase64)
response = get_json()

print(check(response), end="")
