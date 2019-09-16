FROM golang:latest 
ADD ./GoLangCICD.exe /
CMD ["/GoLangCICD.exe"]
EXPOSE 8080