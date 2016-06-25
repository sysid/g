#!/bin/bash

################################################################################
# Jump to directory
################################################################################
JumpList="${HOME}/dev/cfg/g/$(hostname).jump.csv"
g () {
    file=$($GOBIN/g -f $JumpList $1)
    if [ $? -eq 0 ]; then
        builtin cd $file
    fi
}

# Bash Automcompletion: http://askubuntu.com/questions/68175/how-to-create-script-with-auto-complete
_g()
{
  _script_commands=$($GOBIN/g -s -f $JumpList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _g g

################################################################################
# Edit files with standard editor
################################################################################
EditList="${HOME}/dev/cfg/g/$(hostname).edit.csv"
e () {
    file=$($GOBIN/g -f $EditList $1)
    if [ $? -eq 0 ]; then
        e1 $file
    fi
}

# Bash Automcompletion
_e()
{
  _script_commands=$($GOBIN/g -s -f $EditList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _e e
