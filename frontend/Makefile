SERVER1=gz1
all:prepare
	scp -r dist/* $(SERVER1):/home/usera/nginxdir/var/www
	@echo "aaa"


prepare:
	yarn run build
	@echo "prepare"