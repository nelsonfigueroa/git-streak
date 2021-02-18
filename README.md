# git-streak

View GitHub commit stats in the past year.

## Setup

You'll need Go installed. Clone the repo, and get dependencies:

```
go get
```

`git-streak` gets contributions from a specified GitHub username. You have the option of hardcoding this username in the code or passing in the username as an argument when running. Regardless of the hardcoded username, you can still pass in a username as an argument.

If you'd like to hardcode the username, modify the `username` variable in the `main()` function.

Next, build the binary:

```
go build
```

## Running

Get streaks for pre-defined (hardcoded) username:

```
./git-streak

Commits in the past year: 547
Current streak: 7 day(s), since 2021/02/11
Best day: 2020/03/14, with 20 commit(s).
```

Get streaks for a username specified as an argument:

```
./git-streak torvalds

Commits in the past year: 2,550
Current streak: 0 days.
Best day: 2020/08/03, with 47 commit(s).
```