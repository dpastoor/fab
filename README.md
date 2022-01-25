# fab

`fab`ricate a new project from a template... in a `fab`ulous way :-) 

[![asciicast](https://asciinema.org/a/455511.svg)](https://asciinema.org/a/455511)

## setup

The first time on a machine, fab can be initialized via

```shell
# on unix os's this would create ~/.config/fab/config.yml
fab config init 
```

If you want to create it yourself, the config file location follows the XDG base directory spec. 


<details open>
    <summary><strong>Unix-like operating systems</strong></summary>
    <br/>

| <a href="#xdg-base-directory"><img width="400" height="0"></a> | <a href="#xdg-base-directory"><img width="500" height="0"></a><p>Unix</p> | <a href="#xdg-base-directory"><img width="600" height="0"></a><p>macOS</p>                                            | <a href="#xdg-base-directory"><img width="500" height="0"></a><p>Plan 9</p> |
| :------------------------------------------------------------: | :-----------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------: |
| <kbd><b>XDG_CONFIG_HOME</b></kbd>                              | <kbd>~/.config</kbd>                                                      | <kbd>~/Library/Application&nbsp;Support</kbd>                                                                         | <kbd>$home/lib</kbd>                                                        |
| <kbd><b>XDG_CONFIG_DIRS</b></kbd>                              | <kbd>/etc/xdg</kbd>                                                       | <kbd>~/Library/Preferences</kbd><br/><kbd>/Library/Application&nbsp;Support</kbd><br/><kbd>/Library/Preferences</kbd> | <kbd>/lib</kbd>                                                             |

</details>

<details open>
    <summary><strong>Microsoft Windows</strong></summary>
    <br/>

| <a href="#xdg-base-directory"><img width="400" height="0"></a> | <a href="#xdg-base-directory"><img width="700" height="0"></a><p>Known&nbsp;Folder(s)</p> | <a href="#xdg-base-directory"><img width="900" height="0"></a><p>Fallback(s)</p> |
| :------------------------------------------------------------: | :---------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------: |
| <kbd><b>XDG_CONFIG_HOME</b></kbd>                              | <kbd>LocalAppData</kbd>                                                                   | <kbd>%LOCALAPPDATA%</kbd>                                                        |
| <kbd><b>XDG_CONFIG_DIRS</b></kbd>                              | <kbd>ProgramData</kbd><br/><kbd>RoamingAppData</kbd>                                      | <kbd>%ProgramData%</kbd><br/><kbd>%APPDATA%</kbd>                                |

</details>

Hence within (one of) the config dirs, a `fab/config.yml` file should be created.

There are 3 key sections to the config, `settings`, `collections`, and `template_dirs`

* `settings` stores default parameter values to use in any fab prompts
* `collections` are a directory that within the directory have multiple template subdirectories
* `template_dirs` are single template directories

An example `config.yml` might look something like:

```yaml
settings:
  - name: "first_name"
    default: Devin
  - name: "last_name"
    default: Pastoor
  - name: "email"
    default: devin.pastoor@gmail.com 
  - name: "copyright_holder"
    default: "Devin Pastoor <devin.pastoor@gmail.com>"

# all folders within template that have a _setup.yml
collections:
- /Users/devin/templates

# specific folder that is a template
template_dirs:
- /Users/devin/single-repo

```

## how it works

To `fab`ricate a new project - run

```
fab generate
```


`fab` will then find all the template directories listed in template_dirs or collections,
by checking for the "magic" file `_setup.yml` at the root of any found.
It then generates a set of default prompt questions
_and_ questions based on what is provided in the `_setup.yml` plus some defaults.


```yaml
settings:
  - name: "package_name"
    type: string
    prompt: "package name:"
  - name: "use_renv"
    type: boolean
    prompt: "use renv with this project?"
  - name: "email"
    type: string
    prompt: "email:"
  - name: "copyright_holder"
    type: string
    prompt: "copyright holder:"
```

Currently `string` and `boolean` types are supported.

All the settings configured are set such that they can be accessed via
the [go-buffalo/plush](https://github.com/gobuffalo/plush) templating system.

When the template directory is copied to the result directory, the following happens:

* any file with the suffix `.tmpl` is passed through the plush rendering engine
* all file names are passed through go's `text/template` engine, which
uses a slightly different convention, but is simpler for file names.
 These variables can be referenced with a `.` in front of their name within double `{{}}`

for example, given `package_name` it would be referenced in the file path
as `{{.package_name}}` but if used inside a `.tmpl` file can be referred
to directly without the prefix'd `.`.

For booleans, plush's syntax can allow simple conditional logic such as:

```
<%= if (true) { %>
  some content here 
<% } else { %>
  some other content here
<% } %>
```

for example, given a `.Rprofile` with a prompt

```yaml
settings:
  - name: "use_renv"
    type: boolean
    prompt: "use renv with this project?"
```

could then could have within `.Rprofile.tmpl` set it up as:

```
<%= if (use_renv) { %>
source("renv/activate.R")
<% } %>
```

such that the line `source("renv/activate.R")` will only present
given use_renv was set to yes (true)

Plush can also do much more sophisticated templating/loops, writing inline functions, data transformations,
etc. Check out the docs for more info.

This was originally started as a PoC side-project at https://github.com/metrumresearchgroup/fab for general
project scaffolding to make it easier to generate R packages quickly. My hope is that such tooling
will enable people to more readily leverage the R package scaffolding rather than
floating scripts.
 