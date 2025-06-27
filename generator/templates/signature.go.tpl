{{- if eq $.N 1}}
// region: 0

type FnSig0R[R any] struct {
	RV0 R
}

type FnSig0R2[R0, R1 any] struct {
	RV0 R0
	RV1 R1
}

func Signature0R[R any](fn func() R) (s FnSig0R[R]) {
	return
}

func Signature0R2[R0, R1 any](fn func() (R0, R1)) (s FnSig0R2[R0, R1]) {
	return
}
{{end}}
// region: {{.N}}

type FnSig{{.N}}[{{template "Gen" .}}] struct {
	{{- if $.Last}}
	FnSig{{.Last}}[{{range $i, $a := .Args}}{{if ne $i $.Last}}{{if $i}}, {{end}}T{{$i}}{{end}}{{end}}]
	{{- end}}
	Arg{{.Last}} T{{.Last}}
}

type FnSig{{.N}}R[{{template "GenR" .}}] struct {
	FnSig{{.N}}[{{template "Types" .}}]
	FnSig0R[R]
}

type FnSig{{.N}}R2[{{template "GenR2" .}}] struct {
	FnSig{{.N}}[{{template "Types" .}}]
	FnSig0R2[R0, R1]
}

func Signature{{.N}}R[{{template "GenR" .}}](fn func({{template "Types" .}}) R) (s FnSig{{.N}}R[{{range $i, $a := .Args }}T{{$i}}, {{end}}R]) {
	return
}

func Signature{{.N}}[{{template "Gen" .}}](fn func({{template "Types" .}})) (s FnSig{{.N}}[{{template "Types" .}}]) {
	return
}

func Signature{{.N}}R2[{{template "GenR2" .}}](fn func({{template "Types" .}}) (R0, R1)) (s FnSig{{.N}}R2[{{range $i, $a := .Args }}T{{$i}}, {{end}}R0, R1]) {
	return
}
