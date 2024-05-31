# setgituser

`setgituser` is a simple command-line application that sets the git name and email based on environment variables.

## Usage

The application takes an optional command line argument `suffix`. If provided, the suffix is converted to uppercase and used to form the environment variable names for the git name and email. For example, if the suffix is "john", the application will look for the environment variables "GIT_NAME_JOHN" and "GIT_EMAIL_JOHN".

If the suffix is not provided, the application will use `GIT_NAME` and `GIT_EMAIL` as the environment variable names.

The values of these environment variables are then used to set the git name and email. If the name or email is not set in the environment variables, the application will print an error message and exit.

The `suffix` can be provided in two ways:

1. As a command line argument with the `-suffix` flag. For example:

```bash
./setgituser -suffix john
```

2. As the first non-flag argument. For example:

```bash
./setgituser john
```

In both cases, the application will look for the environment variables "GIT_NAME_JOHN" and "GIT_EMAIL_JOHN".

If no suffix is passed to the application, it will default to using the environment variables "GIT_NAME" and "GIT_EMAIL" to set the git name and email respectively. 

Here's an example:

```bash
export GIT_NAME="Default User"
export GIT_EMAIL="default.user@example.com"

./setgituser
```

In this case, the application will set the git name to "Default User" and the git email to "default.user@example.com". If either "GIT_NAME" or "GIT_EMAIL" is not set in the environment variables, the application will print an error message and exit.