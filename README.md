# git-streak

View commit GitHub commit streaks in the past year.

## Setup

You'll need Go installed. Clone the repo, and get dependencies:

```
go get
```

`git-streak` gets contributions from a specified username. You have the option of hardcoding this username in the code or passing in the username as an argument when running.

After the username is updated, build the binary:

```
go build
```

## Running

Get streaks for pre-defined username:

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