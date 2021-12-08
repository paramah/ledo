#!/bin/sh
target=/usr/local/bin
user_arg=$1
stream_cmd="curl -sL install.leaddocker.tech"
readme="https://leaddocker.tech"

determine_os() {
  case "$(uname -s)" in
  Darwin)
    echo "Darwin"
    ;;
  MINGW64*)
    echo "Windows"
    ;;
  *)
    echo "Linux"
    ;;
  esac
}

determine_user_install() {
  case "$(determine_os)" in
  windows)
    echo "--user"
    ;;
  *)
    echo "$user_arg"
    ;;
  esac
}

determine_local_target() {
  case "$(determine_os)" in
  windows)
    # shellcheck disable=SC1003
    echo "$USERPROFILE/bin" | tr '\\' '/'
    ;;
  linux)
    systemd-path user-binaries
    ;;
  esac
}

determine_ledo_binary() {
 case "$(determine_os)" in
  windows)
    echo "ledo.exe"
    ;;
  *)
    echo "ledo"
    ;;
  esac
}

determine_ending() {
 case "$(determine_os)" in
  windows)
    echo "tar.gz"
    ;;
  *)
    echo "tar.gz"
    ;;
  esac
}

handle_user_installation() {
  user_install=$(determine_user_install)
  if [ "$user_install" = "--user" ]; then
    local_target=$(determine_local_target)
    if [ "$local_target" != "" ] && [ ! -d "$local_target" ]; then
      mkdir -p "$local_target"
    fi

    if [ -d "$local_target" ]; then
      target=$local_target
    else
      echo "unfortunately, there is no user-binaries path on your system. aborting installation."
      exit 1
    fi
  fi
}

check_access_rights() {
  if [ ! -w "$target" ]; then
    echo "you do not have access rights to $target."
    echo
    local_target=$(determine_local_target)
    if [ "$local_target" != "" ]; then
      echo "we recommend that you use the --user flag"
      echo "to install the app into your user binary path $local_target"
      echo
      echo "  $stream_cmd | sh -s - --user"
      echo
    fi
    if [ "$(command -v sudo)" != "" ]; then
      echo "calling the installation with sudo might help."
      echo
      echo "  $stream_cmd | sudo sh"
      echo
    fi
    exit 1
  fi
}

install_remote_binary() {
  echo "installing latest 'ledo' release from GitHub to $target..."
  url=$(curl -s https://api.github.com/repos/paramah/ledo/releases/latest |
    grep "browser_download_url.*ledo_.*$(determine_os)_x86_64\.$(determine_ending)" |
    cut -d ":" -f 2,3 |
    tr -d ' \"')
  curl -sSL "$url" | tar xz -C "$target" "$(determine_ledo_binary)" && chmod +x "$target"/ledo
}

add_to_path() {
  case "$(determine_os)" in
  windows)
    powershell -command "[System.Environment]::SetEnvironmentVariable('Path', [System.Environment]::GetEnvironmentVariable('Path', [System.EnvironmentVariableTarget]::User)+';$target', [System.EnvironmentVariableTarget]::User)"
    ;;
  esac
}

display_success() {
  location="$(command -v ledo)"
  echo "Ledo binary location: $location"
}

check_installation_path() {
  location="$(command -v ledo)"
  if [ "$location" != "$target/ledo" ]; then
    echo "(!) The installation location doesn't match the location of the ledo binary."
    echo "    This means that the binary that's used is not the binary that has just been installed"
    echo "    You probably want to delete the binary at $location"
  fi
}

main() {
  handle_user_installation
  check_access_rights
  install_remote_binary
  add_to_path
  display_success
  check_installation_path
}

main
