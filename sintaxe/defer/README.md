# Por que usar Defer?

https://go.dev/blog/defer-panic-and-recover

Go tem os mecanismos usuais para fluxo de controle: if, for, switch, goto. Ele também tem a declaração go para executar código em uma goroutine separada. Aqui eu gostaria de discutir alguns dos menos comuns: defer, panic e recover.

Uma instrução defer coloca uma chamada de função em uma lista. A lista de chamadas salvas é executada após a função circundante retornar. Defer é comumente usado para simplificar funções que realizam várias ações de limpeza.

Por exemplo, vamos analisar uma função que abre dois arquivos e copia o conteúdo de um arquivo para o outro:

```
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}
```

Isso funciona, mas há um bug. Se a chamada para os.Create falhar, a função retornará sem fechar o arquivo de origem. Isso pode ser facilmente remediado colocando uma chamada para src.Close antes da segunda declaração return, mas se a função fosse mais complexa, o problema poderia não ser tão facilmente percebido e resolvido. Ao introduzir declarações defer, podemos garantir que os arquivos sejam sempre fechados:

```
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

As instruções Defer nos permitem pensar em fechar cada arquivo logo após abri-lo, garantindo que, independentemente do número de instruções return na função, os arquivos serão fechados.

O comportamento das instruções defer é direto e previsível. Existem três regras simples:

1. Os argumentos de uma função adiada são avaliados quando a instrução defer é avaliada.

Neste exemplo, a expressão “i” é avaliada quando a chamada Println é adiada. A chamada adiada imprimirá “0” após o retorno da função.


```
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

2. Chamadas de função adiadas são executadas na ordem Último a Entrar, Primeiro a Sair, após o retorno da função circundante.

Esta função imprime “3210”:

```
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```
3. Funções adiadas podem ler e atribuir valores de retorno nomeados à função de retorno.

```
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

