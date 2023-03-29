all: hassh

clean:
	rm -rf out || true

hassh: out/hassh

out/hassh: cmd/hassh internal/hassh x/crypto/ssh x/crypto/ssh/internal
	go build -o out/hassh ./cmd/hassh
