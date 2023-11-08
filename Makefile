
build: rotten-blood-boss-timer ebb-and-flow-timer

rotten-blood-boss-timer:
	GOARCH=amd64 GOOS=linux go build -o rotten-blood-boss-timer-linux -ldflags="-X 'main.WindowTitle=Rotten Blood Boss Timer' -X main.FirstAlarm=5s -X main.SecondAlarm=78s -X main.Interval=90s" ./cmd/tibia-timer
	GOARCH=amd64 GOOS=windows go build -o rotten-blood-boss-timer.exe -ldflags="-X 'main.WindowTitle=Rotten Blood Boss Timer' -X main.FirstAlarm=5s -X main.SecondAlarm=78s -X main.Interval=90s" ./cmd/tibia-timer

ebb-and-flow-timer:
	GOARCH=amd64 GOOS=linux go build -o ebb-and-flow-timer-linux -ldflags="-X 'main.WindowTitle=Ebb and Flow Timer' -X main.FirstAlarm=15s -X main.SecondAlarm=0 -X main.Interval=120s" ./cmd/tibia-timer
	GOARCH=amd64 GOOS=windows go build -o ebb-and-flow-timer.exe -ldflags="-X 'main.WindowTitle=Ebb and Flow Timer' -X main.FirstAlarm=15s -X main.SecondAlarm=0 -X main.Interval=120s" ./cmd/tibia-timer
