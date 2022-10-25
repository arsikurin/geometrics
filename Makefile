.PHONY: prod

prod:
	ssh-keygen -t rsa -N "" -m PEM -f ./id.rsa && ssh-keygen -f id.rsa.pub -e -m pkcs8 > id.rsa.pub.pkcs8
	bin/geometrics
