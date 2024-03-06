front:
	cd frontend && npm run dev
back:
	cd backend/cmd && go run main.go
files:
	cd static && go run main.go