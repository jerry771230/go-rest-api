package swagger

// The following statements MUST starts with '//go:'
//go:generate rm -rf server
//go:generate mkdir -p server
//go:generate swagger generate server --quiet --target server --name hello-api --spec swagger.yml --exclude-main
