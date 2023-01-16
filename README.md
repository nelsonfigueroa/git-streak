# git-streak

View GitHub commit stats in the past year.

## Setup

You'll need Go installed. Then clone the repo. Within the repo install dependencies:

```
go get
```

Before building, know that `git-streak` gets contributions from a specified GitHub username. You have the option of hardcoding this username in the code or passing in the username as an argument when running. Regardless of the hardcoded username, you can still pass in a username as an argument.

If you'd like to hardcode the username, modify the `username` variable in the `main()` function.

Next, build the binary:

```
go build
```

## Running

Get streaks for pre-defined (hardcoded) username:

```
$ ./git-streak


	 ██████╗ ██╗████████╗    ███████╗████████╗██████╗ ███████╗ █████╗ ██╗  ██╗
	██╔════╝ ██║╚══██╔══╝    ██╔════╝╚══██╔══╝██╔══██╗██╔════╝██╔══██╗██║ ██╔╝
	██║  ███╗██║   ██║       ███████╗   ██║   ██████╔╝█████╗  ███████║█████╔╝
	██║   ██║██║   ██║       ╚════██║   ██║   ██╔══██╗██╔══╝  ██╔══██║██╔═██╗
	╚██████╔╝██║   ██║       ███████║   ██║   ██║  ██║███████╗██║  ██║██║  ██╗
	 ╚═════╝ ╚═╝   ╚═╝       ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝

Getting stats for nelsonfigueroa
Commits in the past year: 392
Current streak: 7 days, since 2023/01/10
Best day in the past year: 2022/11/06 with 10 commits.
```

Get streaks for a username specified as an argument:

```
$ ./git-streak torvalds

	 ██████╗ ██╗████████╗    ███████╗████████╗██████╗ ███████╗ █████╗ ██╗  ██╗
	██╔════╝ ██║╚══██╔══╝    ██╔════╝╚══██╔══╝██╔══██╗██╔════╝██╔══██╗██║ ██╔╝
	██║  ███╗██║   ██║       ███████╗   ██║   ██████╔╝█████╗  ███████║█████╔╝
	██║   ██║██║   ██║       ╚════██║   ██║   ██╔══██╗██╔══╝  ██╔══██║██╔═██╗
	╚██████╔╝██║   ██║       ███████║   ██║   ██║  ██║███████╗██║  ██║██║  ██╗
	 ╚═════╝ ╚═╝   ╚═╝       ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝

Getting stats for torvalds
Commits in the past year: 2,549
Current streak: 13 days, since 2023/01/04
Best day in the past year: 2022/12/12 with 81 commits.
```