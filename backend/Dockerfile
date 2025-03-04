##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.23-alpine AS build

ARG PORT

ENV PORT $PORT

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY . .

# download Go modules and dependencies
RUN go mod tidy

# install playwright
RUN PWGO_VER=$(grep -oE "playwright-go v\S+" go.mod | sed 's/playwright-go //g') \
    && go install github.com/playwright-community/playwright-go/cmd/playwright@${PWGO_VER}

EXPOSE ${PORT}

# compile application
RUN go build -o /binary

##
## STEP 2 - DEPLOY
##
FROM ubuntu

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/playwright /playwright
COPY --from=build /binary /binary

COPY --from=build /app/templates /templates
COPY --from=build /app/docs /docs

RUN apt update && apt install -y ca-certificates tzdata
# Enable this if want to install at build time
RUN /playwright install --with-deps
RUN rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/binary"]