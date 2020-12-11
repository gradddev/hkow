module github.com/AlexeySemigradsky/mi-hap

go 1.15

replace github.com/AlexeySemigradsky/mi-hap v0.0.0 => ./
replace github.com/AlexeySemigradsky/mi-hap/led-strip-light v0.0.0 => ./led-strip-light
replace github.com/AlexeySemigradsky/mi-hap/magic-home v0.0.0 => ./magic-home

require (
	github.com/AlexeySemigradsky/mi-hap/led-strip-light v0.0.0
	github.com/AlexeySemigradsky/mi-hap/magic-home v0.0.0 // indirect
	github.com/brutella/hc v1.2.3
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
)
