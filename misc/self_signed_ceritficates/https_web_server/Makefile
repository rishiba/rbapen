postmp:
	curl -X POST http://localhost:9999/usersMPF -F name=name111111 -F email=rishi@gmail.com

posturl:

	curl -X POST \
		   http://localhost:9999/usersUEF \
		   -H 'Content-Type: application/x-www-form-urlencoded' \
		   -H 'Postman-Token: 9da5d67b-950a-46a3-8203-4099e758e20a' \
		   -H 'cache-control: no-cache' \
		   -d 'name=rishi&email=rishi%40gmail.com%0A&undefined='

postImage:
	curl -X POST  -F name=name111111 -F email=rishi@gmail.com  -F uploadfile=@image.png http://localhost:9999/usersMPFWI


postJson:
	curl -X POST -v  -H "Content-Type: application/json"  http://localhost:9999/users  -d '{"name":"rishi","phone1":"97","email":"email@email.com"}'
