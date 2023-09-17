package templates

var DockerFileTemplate_default = `
FROM {{.DockerBaseImage}}:{{.DockerBaseTag}}

ENV DIR /usr/local
WORKDIR ${DIR}

# Copy entrypoint
COPY docker/docker-entrypoint.sh /bin/docker-entrypoint.sh

# Copy project content
COPY ./app $DIR

ENTRYPOINT ["docker-entrypoint.sh"]
CMD [""]
`
var DockerFileTemplate_php = `
FROM {{.DockerBaseImage}}:{{.DockerBaseTag}}
ARG ENVIRONMENT=production

RUN ngxconfig sf.conf

ENV DIR /var/www
WORKDIR ${DIR}

# Copy entrypoint
COPY docker/docker-entrypoint.sh /bin/docker-entrypoint.sh
RUN chmod +x /bin/docker-entrypoint.sh

# Develop packages
RUN xdebug_enable

RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown www-data:www-data ${DIR} && /bin/composer self-update --2
USER www-data

# For Container build cache
COPY ./composer.* $DIR/
RUN /bin/composer install --no-scripts --no-interaction --no-autoloader && composer clear-cache

# Copy application
COPY --chown=www-data:www-data ./ $DIR


ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 80
# done

USER root
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]
`

var DockerFileTemplate_golang = `
#
# C O M P I L E R
#
######################################

FROM {{.DockerBaseImage}}:{{.DockerBaseTag}} as compiler
ARG PACKAGE_NAME={{.}}

# Copy app sources
COPY ./app ${GOPATH}/src/${PACKAGE_NAME}

# # Build app
RUN cd ${GOPATH}/src/${PACKAGE_NAME} \
    && go build ./... \
    && go build main.go

# # Copy built app to workdir and set permissions
RUN mv ${GOPATH}/src/${PACKAGE_NAME}/main /tmp/app \
    && chmod 777 /tmp/app

#
# F I N A L   I M A G E 
#
#######################################

FROM busybox:stable-glibc

COPY --from=compiler /tmp/app /usr/bin/app
COPY docker/*-entrypoint.sh /usr/local/bin/

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["/usr/bin/app"]

EXPOSE 80

`
