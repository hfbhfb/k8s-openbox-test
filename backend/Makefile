# Self documented Makefile
# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
# SERVER1=aws1
SERVER1=gz1
# SERVER1=ubuntu111
DockerImage=qushuiyin
DockerImageVersion=v1
#DockerPortMap="-p 9801:9801" #不用这个改成了docker-compose
RemotePWD=`ssh $(SERVER1) "pwd"`


FixedDockerImageInfo=$(DockerImage):$(DockerImageVersion)
FixedServerDir=~/tmp/$(FixedDockerImageInfo)
FixedDockerlizeDir=/Users/hfb/test/dockerlizedir
FixedDir=program
FixedAppName=runableapp
#FixedDir #FixedAppName  在dockerfile里面写死了的
FixedDockerlizeProgram=$(FixedDockerlizeDir)/$(FixedDir)
FixedProgramTar=$(FixedDir).tar
FixedDockerlizeTar=$(FixedDockerlizeDir)/$(FixedProgramTar)

all:dockerlize
	@echo " all"

dockerlize:dirperpare
	ssh $(SERVER1) "rm -Rf $(FixedServerDir)"
	ssh $(SERVER1) "mkdir -p $(FixedServerDir)"
	scp $(FixedDockerlizeTar) $(SERVER1):$(FixedServerDir)/$(FixedProgramTar)
	ssh $(SERVER1) "cd $(FixedServerDir) && tar xf $(FixedProgramTar)"
	ssh $(SERVER1) "cd $(FixedServerDir)/$(FixedDir) && docker build -t $(FixedDockerImageInfo) ."
	-ssh $(SERVER1) "cd $(FixedServerDir)/$(FixedDir) && docker-compose rm -f -s "
	-ssh $(SERVER1) "cd $(FixedServerDir)/$(FixedDir) && docker-compose up -d "
	@echo 'dockerlized'

dirperpare:
	rm -Rf $(FixedDockerlizeDir)
	mkdir -p $(FixedDockerlizeProgram)
	env GOOS=linux GOARCH=386 go build -o $(FixedAppName)
	cp config-pro.json $(FixedDockerlizeProgram) 
	cp $(FixedAppName) $(FixedDockerlizeProgram) 
	cp Dockerfile $(FixedDockerlizeProgram)
	cp docker-compose.yml $(FixedDockerlizeProgram)
	cd $(FixedDockerlizeDir) && tar cf $(FixedProgramTar) $(FixedDir)



