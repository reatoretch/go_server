run:
	docker-compose up -d --build --scale app=2
test:
	docker-compose -f docker-compose-test.yml up -d --build --scale app_test=2
