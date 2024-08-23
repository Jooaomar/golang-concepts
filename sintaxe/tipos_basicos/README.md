# Tipos básicos
### Os tipos básicos de Go são

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // pseudônimo para uint8

rune // pseudônimo para int32
     // representa um ponto de código Unicode

float32 float64
complex64 complex128

```

O exemplo mostra vários tipos de variáveis e também que as declarações de variáveis podem ser "construídas" em blocos, como com as declarações de importação.

Os tipos int, uint e uintptr são geralmente de 32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits. Quando você precisar de um valor inteiro deverá usar int, a menos que tenha um motivo específico para usar um tipo de inteiro com tamanho especificado ou sem sinal.