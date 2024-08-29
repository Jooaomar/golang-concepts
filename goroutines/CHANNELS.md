# Canais

Canais são um canudo tipado través do qual você pode enviar e receber valores com o opérador de canal, <-.

```
ch <- v // v envia para o canal ch.
v := <- ch // Recebe do ch, e atribui o valor de v
```

(Os dados fluem na direção da seta)

Como maps e slices, os canais devem ser criados antes de se usar:

```
ch := make(chan int)
```

Por padrão, enviam e recebem bloco até o outro lado estar pronto. Isso permite que goroutines sincronizem sem bloqueios explícitos
ou variáveis de condição

O código do exemplo soma os números em uma slice, distribuindo o trabalho entre duas goroutines. Uma vez que ambas as goroutines
complementaram seu processamento, calcula o resultado final.