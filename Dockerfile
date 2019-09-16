FROM golang:latest 
ADD ./BookingApp.exe /
CMD ["/BookingApp.exe"]
EXPOSE 8080