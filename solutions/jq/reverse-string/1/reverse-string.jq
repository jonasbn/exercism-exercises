# echo "hello world" | jq -R '. |split("") |reverse|join("")'
.value |split("")|reverse|join("")
