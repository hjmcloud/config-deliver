FROM alpine
# Install dependencies
ARG TARGETARCH
RUN echo "I'm building for $TARGETARCH"
ENV TZ                      Asia/Shanghai
RUN apk update && apk add tzdata ca-certificates bash
# Install hjm-certcheck
ENV WORKDIR                 /app
ADD resource                $WORKDIR/
COPY temp/config-deliver_linux_$TARGETARCH $WORKDIR/config-deliver
RUN chmod +x $WORKDIR/config-deliver

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./config-deliver