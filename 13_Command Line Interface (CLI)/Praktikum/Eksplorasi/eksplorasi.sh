#!/bin/bash

# Mengambil input endpoint
read -p "insert endpoint: " endpoint

# Mengambil input HTTP method
read -p "insert HTTP method: " method

# Mengambil input request body
if [ $method != "GET" ]; then
	read -p "insert request body: " request_body
fi

# Mengambil input tipe request body
if [[ ! -z $request_body ]]; then
	read -p "insert request body type: " req_body_type
fi

# Mendeklarasikan variabel untuk menampung curl
curl_cmd="curl -X $method $endpoint"

# Mengecek jika request body tidak kosong maka tambahkan ke curl
if [[ ! -z $request_body ]]; then
	curl_cmd+=' -H "Content-Type: $req_body_type" -d "$request_body"'
fi

# Mengeksekusi curl
eval $curl_cmd
