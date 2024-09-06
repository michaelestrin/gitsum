# GitSum

Augment and preserve your GitHub contribution history for git repositories hosted elsewhere.

Creates a separate git repository with a summary of your commit history but none of the original repository's content.

## Build

```go build cmd/gitsum/gitsum.go```

## Run

```./gitsum -author authorNameContent -in /path/to/original/repo -out /path/to/gitsum/repo/to/create```

Where 

- `authorNameContent` is text to match in the git commit author's name, 
- `/path/to/original/repo` is the path to the root of the repository to summarize, and 
- `/path/to/gitsum/repo/to/create` is the path to the root of the summary repository to create.