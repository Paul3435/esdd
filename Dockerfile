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
ENV SENDGRID_API_KEY=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
ENV MAILGUN_API_KEY=02057f3a67aec395a2efd2e70426f144-911539ec-96815074

#Server
EXPOSE ${APP_PORT}

#Run
CMD ["./DreamDataApp"]