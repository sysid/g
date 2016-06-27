# g
Jumper Utility 'G' uses CSV file to provide shortcuts for your most used paths or files to edit.

### Usage
First source the g.sh file in your bashrc:
    . path/g.sh

Now th shortcuts g: goto and e: edit are available. You need to provide CSV configuration
file for both shortcuts.

    g dev
will jump to directory 
    ~/dev.

    e profile
will open the profile file in your editor.

#### Supports full bash completion.

#### Example CSV configuration file for jumping:
    # MY JUMPLISTFILE
    dev, $HOME/dev
    bla, /foo/bar

#### Example CSV configuration file for editing:
    # MY JUMPLISTFILE
    profile, $HOME/dev/$COMPUTERNAME.profile
