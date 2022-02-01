package templates

var DockerEntrypointTemplate_bash = `#!/bin/bash
set -eo pipefail
shopt -s nullglob

h1() {
    declare border padding text
    border='\e[1;34m'"$(printf '=%.0s' $(seq 1 "$term_width"))"'\e[0m'
    padding="$(printf ' %.0s' $(seq 1 $(((term_width - $(wc -m <<<"$*")) / 2))))"
    text="\\e[1m$*\\e[0m"
    echo -e "$border"
    echo -e "${padding}${text}${padding}"
    echo -e "$border"
}

h2() {
    printf '\e[1;33m==>\e[37;1m %s\e[0m\n' "$*"
}

h1 "Project init"

###
### Your code here
###

h1 "End of init"

exec "$@"
`

var TestEntrypointTemplate_bash = `#!/bin/bash
set -eo pipefail
shopt -s nullglob

h1() {
    declare border padding text
    border='\e[1;34m'"$(printf '=%.0s' $(seq 1 "$term_width"))"'\e[0m'
    padding="$(printf ' %.0s' $(seq 1 $(((term_width - $(wc -m <<<"$*")) / 2))))"
    text="\\e[1m$*\\e[0m"
    echo -e "$border"
    echo -e "${padding}${text}${padding}"
    echo -e "$border"
}

h2() {
    printf '\e[1;33m==>\e[37;1m %s\e[0m\n' "$*"
}

h1 "Project test init"

###
### Your code here
###

h1 "End of test init"

exec "$@"
`