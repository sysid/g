# vim: fdm=marker ts=4 sts=4 sw=4 fdl=0
#!/bin/bash

if [ $(uname -n) == "twdev" ]; then
    TW_GBIN="/usr/bin/g"
else
    #TW_GBIN="$GOBIN/g"
    TW_GBIN="$HOME/dev/binx/g"
fi

################################################################################
# Jump to directory
################################################################################
#### Jumplist {{{
JumpList="${HOME}/dev/cfg/g/$(hostname).jump.csv"
g () {
    file=$($TW_GBIN -f $JumpList $1)
    if [ $? -eq 0 ]; then
        builtin cd $file
    fi
}

# Bash Automcompletion: http://askubuntu.com/questions/68175/how-to-create-script-with-auto-complete
_g()
{
  _script_commands=$($TW_GBIN -s -f $JumpList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _g g
####}}}}

################################################################################
# Edit files with standard editor
################################################################################
#### Edit tmux {{{
EditList="${HOME}/dev/cfg/g/$(hostname).edit.csv"
q () {
    file=$($TW_GBIN -f $EditList $1)
    if [ $? -eq 0 ]; then
        #tmux new-window -n $1\; send-keys "oVim $file" "Enter"
        tmux new-window -n $1 "docker run -ti --rm -v $HOME/dev/vim/oVim:/ext/ -v $HOME:/home/dvlpr/mnt sysid/ovimionated ${file#$HOME/}"
        #e1 $file # put your prefered editor here
    fi
}

# Bash Automcompletion
_q()
{
  _script_commands=$($TW_GBIN -s -f $EditList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _q q

####}}}}

#### Edit GUI {{{
EditList="${HOME}/dev/cfg/g/$(hostname).edit.csv"
e () {
    file=$($TW_GBIN -f $EditList $1)
    if [ $? -eq 0 ]; then
        e1 $file # put your prefered editor here
        #vim $file # put your prefered editor here
    fi
}
ee () {
    file=$($TW_GBIN -f $EditList $1)
    if [ $? -eq 0 ]; then
        e11 $file # put your prefered editor here
    fi
}

# Bash Automcompletion
_e()
{
  _script_commands=$($TW_GBIN -s -f $EditList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _e e
_ee()
{
  _script_commands=$($TW_GBIN -s -f $EditList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _ee ee
####}}}}


################################################################################
# Open files with OSX Standard Application
################################################################################
#### Open {{{
OpenList="${HOME}/dev/cfg/g/$(hostname).open.csv"
o () {
    file=$($TW_GBIN -f $OpenList $1)
    if [ $? -eq 0 ]; then
        open $file
    fi
}

# Bash Automcompletion
_o()
{
  _script_commands=$($TW_GBIN -s -f $OpenList $1)

  local cur prev
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  COMPREPLY=( $(compgen -W "${_script_commands}" -- ${cur}) )

  return 0
}
complete -o nospace -F _o o
####}}}}
