build-image:
	docker build -t gabrielcamurcab/finance_auth -f Dockerfile.app .

run-app:
	docker-compose -f .microsservice/app.yml up -d