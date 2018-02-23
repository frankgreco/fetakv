package integration

type mockIO struct {
	stdin, stdout, stderr []byte
}

var simple = &mockIO{
	stdin: []byte(
		`WRITE foo bar
READ foo`,
	),
	stdout: []byte(
		`bar
`,
	),
}

var complex = &mockIO{
	stdin: []byte(
		`ABORT
reAD foo
WRITE foo bar
READ foo
DELETE foo
read foo
COMMIT
STart
ABORT
start
COMMIT
START
START
ABORT
COMMIT
ABORT
COMMIT
START
WRITE one one
READ one
ABORT
READ one
FOO
START
WRITE one two
READ one
commit
READ one
ABORT
start
START
WRITE one THREE
COMMIT
read one
ABORT
READ one
START
write one four
READ one
START
WRITE one five
COMMIT
READ one
COMMIT
read one
COMMIT`,
	),
	stdout: []byte(
		`bar
one
two
two
THREE
two
four
five
five
`,
	),
	stderr: []byte(
		`There are no current transactions to abort.
Key not found: foo
Key not found: foo
There are no current transactions to commit.
There are no current transactions to abort.
There are no current transactions to commit.
Key not found: one
Command not found: FOO
There are no current transactions to abort.
There are no current transactions to commit.
`,
	),
}
