#!/bin/sh
# 
# tmux script to initialize go dev environment 
#
# deprecated
#   use tmuxinator instead. see: 'make create_env'
#
if [ "$#" -lt 2 ]; then
  echo "Usage: $0 session_name window" >&2
  echo "  e.g. $0 work go-endp" >&2
  exit 1
fi

SESSION=$1
WINDOW=$2

# if session doesn't exist, create
if [ ! `tmux has-session -t $SESSION` ]; then
  tmux new-session -s $SESSION -n shell -d 'cd ~; bash -i'
fi

# create a new window and fire up panes for development

# editor
tmux new-window -t $SESSION -n $WINDOW 'pwd ; bash -i'
tmux send-keys -t $SESSION:$WINDOW 'emacs makefile' 'C-m'

# server
tmux split-window -v -t $SESSION:$WINDOW -d 'pwd ; bash -i'

# client/console
tmux split-window -v -t $SESSION:$WINDOW.1 -d 'pwd ; bash -i'

# start the server and issue a command in the client console
tmux send-keys -t $SESSION:$WINDOW.2 'make run_server_gin' 'C-m'
tmux send-keys -t $SESSION:$WINDOW.1 'ls' 'C-m'

# attach to the session
tmux attach -t $SESSION 

