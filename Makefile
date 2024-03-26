
build: rotten-blood-boss-timer ebb-and-flow-timer mastermind-timer

mastermind-timer:
	GOARCH=amd64 GOOS=linux go build -o mastermind-timer-linux -ldflags="-X 'main.WindowTitle=Mastermind Timer' -X main.FirstAlarm=1s -X main.SecondAlarm=0s -X main.Interval=10m" ./cmd/tibia-timer
	GOARCH=amd64 GOOS=windows go build -o mastermind-timer.exe -ldflags="-X 'main.WindowTitle=Mastermind Timer' -X main.FirstAlarm=1s -X main.SecondAlarm=0s -X main.Interval=10m" ./cmd/tibia-timer

rotten-blood-boss-timer:
	GOARCH=amd64 GOOS=linux go build -o rotten-blood-boss-timer-linux -ldflags="-X 'main.WindowTitle=Rotten Blood Boss Timer' -X main.FirstAlarm=5s -X main.SecondAlarm=0s -X main.Interval=90s" ./cmd/tibia-timer
	GOARCH=amd64 GOOS=windows go build -o rotten-blood-boss-timer.exe -ldflags="-X 'main.WindowTitle=Rotten Blood Boss Timer' -X main.FirstAlarm=5s -X main.SecondAlarm=0s -X main.Interval=90s" ./cmd/tibia-timer

ebb-and-flow-timer:
	GOARCH=amd64 GOOS=linux go build -o ebb-and-flow-timer-linux -ldflags="-X 'main.WindowTitle=Ebb and Flow Timer' -X main.FirstAlarm=15s -X main.SecondAlarm=0 -X main.Interval=120s" ./cmd/tibia-timer
	GOARCH=amd64 GOOS=windows go build -o ebb-and-flow-timer.exe -ldflags="-X 'main.WindowTitle=Ebb and Flow Timer' -X main.FirstAlarm=15s -X main.SecondAlarm=0 -X main.Interval=120s" ./cmd/tibia-timer
