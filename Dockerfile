FROM nixos/nix as builder

RUN mkdir -p /etc/nix && echo "experimental-features = nix-command flakes" > /etc/nix/nix.conf

WORKDIR /app

COPY . /app

RUN nix build .#app -o ./result

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/result/bin/app /app/app

EXPOSE 8080

CMD ["./app"]
