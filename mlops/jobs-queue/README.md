# ML Ops

[[_TOC_]]

## Golang and `jobs_server`, `jobs_client`

```shell
wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
rm go1.23.2.linux-amd64.tar.gz
```

Do not forget to append `/usr/local/go/bin` in `$PATH`.

You can the following line in `~/.bashrc`

```shell
[[ ":$PATH:" != *":/usr/local/go/bin:"* ]] && export PATH="${PATH}:/usr/local/go/bin"
```

To build the server and client run

```shell
mkdir ${HOME}/bin
cd mlops/jobs-queue/
go mod tidy
go build -o ${HOME}/bin/jobs_server ./cmd/server
go build -o ${HOME}/bin/jobs_client ./cmd/client
```

Finally, add the following line to your `~/.bashrc`.

```shell
[[ ":$PATH:" != *":$HOME/bin:"* ]] && export PATH="${PATH}:$HOME/bin"
```