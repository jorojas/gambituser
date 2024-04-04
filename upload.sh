git add .
git commit -m "Ultimo Commit"
git push

go build main.go
rm main.zip
zip -r main.zip main