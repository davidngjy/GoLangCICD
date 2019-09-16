FROM golang:latest 
ADD ./BookingApp /
CMD ["/BookingApp"]
EXPOSE 8080