#!/bin/bash

# Fungsi untuk validasi nama proyek
validate_project_name() {
	if [[ ! $project_name =~ ^[a-zA-Z0-9_-]+$ ]]; then
		echo "invalid project name!"
		exit 1
	fi
}

# Fungsi untuk membuat proyek
generate_project() {
	echo "creating" $project_name "project.."

	if [ $project_type == "simple" ]; then
		mkdir $project_name

		cd $project_name

		go mod init $project_name

		touch main.go
	elif [ $project_type == "api" ]; then
		mkdir $project_name

		cd $project_name

		go mod init $project_name

		mkdir routes models repositories services configs controllers

		touch main.go
	else
		echo "invalid project type"
		exit 1
	fi

	echo "project" $project_name "created successfully"
}

# Membaca input nama proyek
read -p "enter project name: " project_name

# Validasi nama proyek
validate_project_name $project_name

# Membaca input jenis proyek
read -p "choose project type (simple, api): " project_type

# Membuat proyek
generate_project $project_type
