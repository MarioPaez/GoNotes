module taskManager

replace core => ./core

replace formatter.local => ./formatter

go 1.23.2

require core v0.0.0-00010101000000-000000000000

require (
	formatter.local v0.0.0-00010101000000-000000000000 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
