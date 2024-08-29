# Parâmetros de tipo

As funções Go podem ser escritas para funcionar em vários tipos usando parâmetros de tipo. Os parâmetros de tipo de uma função aparecem entre colchetes, antes dos argumentos da função.

```
func Index[T comparable](s []T, x T) int
```

Esta declaração significa que s é um slice de qualquer tipo T que cumpre o restrição interna comparable. x também é um valor do mesmo tipo.

`comparable` é uma restrição útil que torna possível usar os operadores `==` e `!=` em valores do tipo. Neste exemplo, usamos para comparar um valor para todos os elementos da slice até que uma correspondência seja encontrada. Esta função Index funciona para qualquer tipo que suporte comparação.