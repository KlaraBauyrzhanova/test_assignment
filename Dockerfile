FROM golang
LABEL maintainer="klara.chess.school@gmail.com"
RUN mkdir bandlab
WORKDIR /bandlab
COPY . .
RUN go get -v && go build -o /bandlab
EXPOSE 8000
CMD ["./bandlab"]