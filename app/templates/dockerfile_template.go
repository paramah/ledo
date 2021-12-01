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

# For Docker build cache
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
