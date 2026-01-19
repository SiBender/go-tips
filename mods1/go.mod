module mods1

go 1.25.5

replace my-module => ./my-module

require (
	github.com/SiBender/moduleexample v1.0.1
	my-module v0.0.0-00010101000000-000000000000
)
