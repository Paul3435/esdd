FROM golang:1.22.5

#Copy destination
WORKDIR /app

#Download modules
COPY go.mod go.sum ./
RUN go mod download

#Copy src
COPY . .

#Build
RUN go build -o DreamDataApp ./cmd

#Export env variables
ENV APP_PORT=8080
ENV SENDGRID_API_KEY=
ENV MAILGUN_API_KEY=

#Server
EXPOSE ${APP_PORT}

#Run
CMD ["./DreamDataApp"]
