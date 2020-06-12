.PHONY: all test clean build docker

build:
	npm run build
	docker build -t shippy-ui:latest .

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml