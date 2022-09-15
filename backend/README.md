

docker build -t k8s-openbox-test:v0.0.1 .

docker run -itd --name openbox9383261832 -p 16027:1027 k8s-openbox-test:v1



docker rm -f openbox9383261832

