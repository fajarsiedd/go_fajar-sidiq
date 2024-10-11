#!/bin/sh

#Membuat folder
mkdir foods

mkdir drinks

mkdir snacks

#Mendeklarasikan variabel untuk filepath
foodsMenu="foods/menu.txt"

drinksMenu="drinks/menu.txt"

snacksMenu="snacks/menu.txt"

#Membuat file
touch $foodsMenu

touch $drinksMenu

touch $snacksMenu

#Mengisi file menu foods

echo "nasi goreng" >> $foodsMenu

echo "ayam goreng" >> $foodsMenu

echo "bubur ayam" >> $foodsMenu

#Mengisi file menu drinks
echo "kopi susu" >> $drinksMenu

echo "susu oat" >> $drinksMenu

echo "es teh" >> $drinksMenu

#Mengisi file menu snacks
curl https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json | grep "name" | awk -F':' '{gsub(/[",]/, ""); print $2}' > snacks.txt

mv snacks.txt $snacksMenu
