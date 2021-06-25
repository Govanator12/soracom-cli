[![go report status](https://goreportcard.com/badge/github.com/soracom/soracom-cli)](https://goreportcard.com/report/github.com/soracom/soracom-cli)
![build-artifacts](https://github.com/soracom/soracom-cli/actions/workflows/build-artifacts.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/soracom/soracom-cli.svg)](https://pkg.go.dev/github.com/soracom/soracom-cli)

# soracom-cli

A command line tool `soracom` to invoke the SORACOM API.

# Features

The `soracom` command:

- Automatically supports new APIs functions as they are added to the platform. The binary file of the soracom command is automatically generated from the API definition file.

- Just works by copying the cross-compiled binary file into the target environment. There is no need to build an environment or solve dependencies.

- Constructs a request based on the specified argument and forwards it to the SORACOM API. The JSON response from the API is then converted and output directly as standard output, making it easier to process the output of the intial command and pass it to another command

- Supports bash completion. To enable thsi feature, please add the following line in .bashrc etc
  ```
  eval "$(soracom completion bash)"
  ```

   Mac OS users will need to either:
    1. Use `bash` version >= 4.0, or
    2. Use `brew install bash-completion` instead of using the Xcode version of bash-completion, and then add the following to either your `.bash_profile` or `.profile`:

  ```
  if [ -f $(brew --prefix)/etc/bash_completion ]; then
    . $(brew --prefix)/etc/bash_completion
  fi
  ```
  If this change is not made you may recieve the following error:
  ```
  -bash: __ltrim_colon_completions: command not found
  ```

- Supports zsh completion. To enable zsh completion, first generate the completion script by running `soracom completion zsh`, then rename the generated script as `_soracom` and place the file somewhere in your $fpath.

# Installation

## For Mac (macOS):

-  Intallation is performed using homebrew with the following commands:
```
$ brew tap soracom/soracom-cli
$ brew install soracom-cli
$ brew install bash-completion
```

## For Other Operating Systems:
Download a package file that match the environment of the target from [Releases page](https://github.com/soracom/soracom-cli/releases), unpack it, and place the executable file in the directory where included in PATH.


# Usage

## Initial Configuration

Before using the Soracom CLI tool to make API calls, you must configure it to associate those calls with your Soracom Operator account. To do so, create a profile by running the following command:

```
soracom configure
```

You will be asked to select whether your account was created as a Gloabal or Japanese coverage account.

```
Please select which coverage type to use.

1. Global
2. Japan

select (1-2) >
```

Please select the coverage type which you primarily use. For most users this will be Global, however if you live in Japan and use SIM cards in Japan, please select Japan.

Next you will be asked to select an authentication method.

```
Please select which authentication method to use.

1. Input AuthKeyId and AuthKey * Recommended *
2. Input Operator credentials (Operator Email and Password)
3. Input SAM credentials (OperatorId, User name and Password)

select (1-3) >
```

Please select 1 if an AuthKey (authentication key) has been issued to the SAM user or root account being configured.
(For details on how to issue an authentication key to SAM users, please see [Using SORACOM Access Management to Manage Operation Access](https://dev.soracom.io/en/start/sam/).

Future API calls will be made using the authentication information configured here.


## Using Multiple Profiles

If you have multiple SORACOM accounts or want to use multiple SAM users differently, specify the --profile option to configure and set the profile name.

```
soracom configure --profile user1
  :
  (Enter information for user1)

soracom configure --profile user2
  :
  (Enter information for user2)
```

This will create profiles named user1 and user2.
To use the profile, append the `--profile` flag to the normal command as shown below:

```
soracom subscribers list --profile user1
  :
  (SIM list for user1 will be displayed)

soracom groups list --profile user2
  :
  (Group list for user2 will be displayed)
```


## Creating a Profile for API Sandbox

soracom-cli can also be used for setting up a [SORACOM API Sandbox](https://dev.soracom.io/en/docs/api_sandbox/) environment.

In order to create a profile for sandbox, use the `configure-sandbox` subcommand.

```
soracom configure-sandbox
```

You will be prompted to configure a user for the API Sandbox. After configuring this user a profile named `sandbox` will be created. You can issue commands to the API sandbox with this profile by appending the `--profile sandbox` flag to your commands:

```
soracom subscribers list --profile sandbox
```

Appending the `--profile snadbox` flag will also enable the usage of sandbox specific API functions:

```
soracom sandbox subscribers create --profile sandbox
```

An alternate name can also be configured for the sandbox account with the following commands:

```
soracom configure-sandbox --profile test
soracom sandbox subscribers create --profile test
```

For easier to use with shell scripts or other prewritten commands, all parameters necessary for profile creation can be specified with argument:

```
soracom configure-sandbox --coverage-type jp --auth-key-id="$AUTHKEY_ID" --auth-key="$AUTHKEY" --email="$EMAIL" --password="$PASSWORD"
```


### Calling the API via a Proxy

To enable a proxy, add `http://your-proxy-name:port` to the HTTP_PROXY environment variable on your system, then execute soracom command.

e.g.) For Linux / Mac:
Assuming that the address of the proxy server is 10.0.1.2 and the port number is 8080
```
export HTTP_PROXY=http://10.0.1.2:8080
soracom subscribers list
```

Or

```
HTTP_PROXY=http://10.0.1.2:8080 soracom subscribers list
```


### Troubleshooting

If you get an error message like the following:

```
Error: Permissions for the file 'path/to/default.json' which contains your credentials are too open.
It is required that your credential files are NOT accessible by others.
```

Please try the following to fix it:

```
soracom unconfigure
soracom configure
```

i.e. perform `unconfigure` and then `configure` again in order to re-create a credentials file with the appropriate permissions.


# How to Build / Test

For developers that want to build from source or for those who wish to make a pull request such as a bug fix or function addition, please build and test in one of the following ways.

## Building in a Local Environment (Linux / Mac OS X)

In the environment where Go is installed, run the following build script:

```
./scripts/build.sh 1.2.3
```

In this case 1.2.3 is the version number, however the actual version number being used will differ by individual use case. Please be sure to specify an appropriate number.

If the build succeeds, then run the test:

```
export SORACOM_AUTHKEY_ID_FOR_TEST=...   # set AuthKey ID & AuthKey of a Soracom operator (account) to use the API sandbox.
export SORACOM_AUTHKEY_FOR_TEST=...
./test/test.sh
```


# How to Release

```
VERSION=1.2.3                         # => specify a version number to be released
./scripts/build.sh $VERSION           # => build a version to be released
./test/test.sh $VERSION               # => test the version
# commit & push all changes to github
./scripts/release.sh $VERSION         # => release the version to GitHub
# edit the release on github.com release page
./scripts/update-homebrew-formula.sh $VERSION $GITHUB_USERNAME $GITHUB_EMAIL
./scripts/build-snap.sh $VERSION
./scripts/release-snap.sh $VERSION
./scripts/build-lambda-layer.sh $VERSION
./scripts/release-lambda-layer.sh $VERSION $AWS_PROFILE   # => this command releases the layer to all regions (except ap-east-1)
```
