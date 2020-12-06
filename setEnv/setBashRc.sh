#!/bin/bash

echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> $HOME/.bashrc 
echo 'eval "$(rbenv init -)"' >> $HOME/.bashrc
source $HOME/.bashrc